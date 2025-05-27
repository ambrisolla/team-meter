package jiraFetcher

import (
	"team-meter/config"

	"go.uber.org/zap"
)

type JiraFetcher struct {
	JiraConfig config.JiraConfig
	Logget     *zap.Logger
}

func New(c *config.Config, l *zap.Logger) *JiraFetcher {
	return &JiraFetcher{
		JiraConfig: c.Jira,
	}
}
