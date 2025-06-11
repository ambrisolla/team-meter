package jiraService

import (
	"fmt"
	"regexp"
	"strings"
	"team-meter/config"
	jiraFetcher "team-meter/internal/fetcher/jira"
	jiraRepository "team-meter/internal/repository/jira"
	"time"
)

const customTimeLayout = "2006-01-02T15:04:05.000-0700"

func (s *JiraService) SyncIssues(jiraConfig config.JiraConfig) {
	projects := jiraConfig.Projects
	for _, project := range projects {
		s.getIssues(project)
	}
}

func (s *JiraService) getIssues(project string) {
	page := 0
	totalIssues := 0

	logMessage := fmt.Sprintf("starting processing %s issues", project)
	s.Logger.Info(logMessage)

	for {
		issues, err := s.Fetcher.GetIssues(project, page)
		if err != nil {
			panic(err)
		}

		if len(issues.Issues) == 0 {
			break
		}

		totalIssues = totalIssues + len(issues.Issues)

		logMessage := fmt.Sprintf("Processing %d issues for project %s (page %d)", len(issues.Issues), project, page+1)
		s.Logger.Info(logMessage)
		s.saveIssues(issues, project)

		page++
	}
	msg := fmt.Sprintf("Finished processing %d issues for project %s", totalIssues, project)
	s.Logger.Info(msg)
}

func (s *JiraService) saveIssues(issues jiraFetcher.Issues, project string) {
	for _, issue := range issues.Issues {
		createdTime, err := time.Parse(customTimeLayout, issue.Fields.Created)
		if err != nil {
			panic(err)
		}

		updatedTime, err := time.Parse(customTimeLayout, issue.Fields.Updated)
		if err != nil {
			panic(err)
		}

		// Get product
		product, err := s.getIssueProduct(issue)
		if err != nil {
			panic(err)
		}

		err = s.Repository.SaveIssue(jiraRepository.Issue{
			ID:            issue.ID,
			Self:          issue.Self,
			Key:           issue.Key,
			Summary:       issue.Fields.Summary,
			AssigneeName:  issue.Fields.Assignee.DisplayName,
			AssigneeEmail: issue.Fields.Assignee.EmailAddress,
			AccountID:     issue.Fields.Assignee.AccountID,
			IssueType:     issue.Fields.IssueType.Name,
			CreatedAt:     createdTime,
			UpdatedAt:     updatedTime,
			Status:        issue.Fields.Status.StatusCategory.Name,
			ProjectKey:    project,
			Product:       product,
		})
		if err != nil {
			msg := fmt.Sprintf("Fail to save issue %s", issue.Key)
			s.Logger.Error(msg)
		}

		s.saveIssuesChangelogs(issue)

	}
}

func (s *JiraService) saveIssuesChangelogs(issue jiraFetcher.Issue) error {
	for _, history := range issue.Changelog.Histories {
		for _, item := range history.Items {
			if item.FieldType == "jira" && item.FieldID == "status" {
				createdTime, err := time.Parse(customTimeLayout, history.Created)
				if err != nil {
					panic(err)
				}

				issueChangelogItem := &jiraRepository.IssueChangelog{
					IssueID:           issue.ID,
					ID:                history.ID,
					AssigneeAccountID: history.Author.AccountID,
					CreatedAt:         createdTime,
					ToString:          item.To,
					FromString:        item.From,
				}

				err = s.Repository.SaveIssueChangelog(issueChangelogItem)
				if err != nil {
					msg := fmt.Sprintf("Fail to save histyory changelog %s", history.ID)
					s.Logger.Error(msg)
				}
			}
		}
	}
	return nil
}

func (s *JiraService) getIssueProduct(issue jiraFetcher.Issue) (string, error) {
	isProductRow := false // controle para determinar quando que o produto for encontrado
	reMatch := regexp.MustCompile("^Produto")
	product := "uncategorized"

	// first filter
	for _, d := range issue.Fields.Description.Content {
		for _, c := range d.Content {
			if isProductRow {
				product = strings.ReplaceAll(c.Text, ": ", "")
				break
			}
			if reMatch.MatchString(c.Text) {
				isProductRow = true
			}
		}
		if isProductRow {
			break
		}
	}

	// second filter
	for _, p := range s.Config.Products {
		for _, stringToMatch := range p.MatchesWith {
			if strings.Contains(strings.ToLower(issue.Fields.Summary), strings.ToLower(stringToMatch)) || strings.Contains(strings.ToLower(product), strings.ToLower(stringToMatch)) {
				product = p.Name
			}
		}
	}
	return product, nil
}
