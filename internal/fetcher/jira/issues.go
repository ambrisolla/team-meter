package jiraFetcher

import (
	"fmt"
	"net/url"

	"resty.dev/v3"
)

func (f *JiraFetcher) GetIssues(project string, page int) (Issues, error) {

	// Define valores para montar o resultado
	maxResults := 100
	var startAt int
	if page == 0 {
		startAt = 0
	} else {
		startAt = page * maxResults
	}

	// monta a query
	jqlRaw := fmt.Sprintf("project = %s AND updated >= %s", project, f.JiraConfig.SyncStartDate)
	jqlEscaped := url.QueryEscape(jqlRaw)

	query := fmt.Sprintf(
		"%s/rest/api/3/search?jql=%s&maxResults=%d&startAt=%d&expand=changelog",
		f.JiraConfig.Url,
		jqlEscaped,
		maxResults,
		startAt,
	)

	// Inicia a requisicao
	var issues Issues
	client := resty.New()
	defer client.Close()

	res, err := client.R().
		EnableTrace().
		SetResult(&issues).
		SetBasicAuth(f.JiraConfig.User, f.JiraConfig.Pass).
		Get(query)

	if err != nil {
		fmt.Println("Erro na requisicao", err)
	}

	if res.IsError() {
		return issues, fmt.Errorf("http error: %d %s", res.StatusCode(), res.Status())
	}

	return issues, nil

}
