package jiraRepository

import "time"

type Issue struct {
	ID            string    `json:"id"`
	Self          string    `json:"self"`
	Key           string    `json:"key"`
	Summary       string    `json:"summary"`
	AssigneeName  string    `json:"assigneeName"`
	AssigneeEmail string    `json:"assigneeEmailAddress"`
	AccountID     string    `json:"assigneeAccountId"`
	IssueType     string    `json:"issueType"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime:false"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"autoUpdateTime:false"`
	Status        string    `json:"status"`
	ProjectKey    string    `json:"projectKey"`
	Product       string    `json:"product"`
}

type IssueChangelog struct {
	IssueID           string    `json:"issueId"`
	ID                string    `json:"id"`
	AssigneeAccountID string    `json:"assigneeAccountId"`
	CreatedAt         time.Time `json:"createdAt" gorm:"autoCreateTime:false"`
	ToString          string    `json:"to"`
	FromString        string    `json:"from"`
}

func (Issue) TableName() string {
	return "issues"
}
func (IssueChangelog) TableName() string {
	return "issue_changelogs"
}
