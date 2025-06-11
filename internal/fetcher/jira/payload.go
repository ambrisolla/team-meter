package jiraFetcher

type Issue struct {
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Type   string `json:"type"`
	Fields struct {
		Summary     string `json:"summary"`
		Created     string `json:"created"`
		Updated     string `json:"updated"`
		Description struct {
			Content []DescriptionContent
		} `json:"description"`
		Status struct {
			StatusCategory struct {
				Name string `json:"name"`
			} `json:"statusCategory"`
		} `json:"status"`
		Assignee struct {
			DisplayName  string `json:"displayName"`
			AccountID    string `json:"accountId"`
			EmailAddress string `json:"emailAddress"`
			Self         string `json:"self"`
		} `json:"assignee"`
		IssueType struct {
			Name string `json:"name"`
		} `json:"issueType"`
	} `json:"fields"`
	Changelog struct {
		Histories []History `json:"histories"`
	} `json:"changelog"`
}

type Issues struct {
	Issues []Issue `json:"issues"`
}

type History struct {
	IssueID int64  `json:"-"`
	ID      string `json:"id"`
	Author  struct {
		DisplayName  string `json:"displayName"`
		AccountID    string `json:"accountId"`
		EmailAddress string `json:"emailAddress"`
		Self         string `json:"self"`
		Active       bool   `json:"active"`
	} `json:"author"`
	Created string `json:"created"`
	Items   []struct {
		FieldType string `json:"fieldtype"`
		Field     string `json:"field"`
		FieldID   string `json:"fieldId"`
		From      string `json:"fromString"`
		To        string `json:"toString"`
	} `json:"items"`
}

type DescriptionContent struct {
	Type    string `json:"type"`
	Content []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"content"`
}
