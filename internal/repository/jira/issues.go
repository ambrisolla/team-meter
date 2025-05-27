package jiraRepository

func (r *JiraRepository) SaveIssue(issue Issue) error {
	return r.Db.Save(issue).Error
}

func (r *JiraRepository) SaveIssueChangelog(issueChangelog *IssueChangelog) error {
	return r.Db.Save(issueChangelog).Error
}
