package jiraScheduler

import (
	"team-meter/config"
	jiraService "team-meter/internal/service/jira"
	"time"

	"go.uber.org/zap"
)

func Run(c *config.Config, l *zap.Logger) {

	service := jiraService.New(c, l)

	for {
		service.SyncIssues(c.Jira.Projects)
		time.Sleep(c.Jira.SyncInterval)
	}
}
