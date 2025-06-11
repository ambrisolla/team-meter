package jiraService

import (
	"team-meter/config"
	jiraFetcher "team-meter/internal/fetcher/jira"
	jiraRepository "team-meter/internal/repository/jira"

	"go.uber.org/zap"
)

type JiraService struct {
	Repository jiraRepository.JiraRepository
	Fetcher    jiraFetcher.JiraFetcher
	Logger     *zap.Logger
	Config     config.JiraConfig
}

func New(c *config.Config, logger *zap.Logger) *JiraService {
	repository := jiraRepository.New(c, logger)
	fetcher := jiraFetcher.New(c, logger)
	return &JiraService{
		Repository: *repository,
		Fetcher:    *fetcher,
		Logger:     logger,
		Config:     c.Jira,
	}
}
