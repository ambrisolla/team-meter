package jiraService

import (
	"fmt"
	jiraFetcher "team-meter/internal/fetcher/jira"
	jiraRepository "team-meter/internal/repository/jira"
	"time"
)

const customTimeLayout = "2006-01-02T15:04:05.000-0700"

func (s *JiraService) SyncIssues(projects []string) {
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
