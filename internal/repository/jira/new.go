package jiraRepository

import (
	"team-meter/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type JiraRepository struct {
	Db     *gorm.DB
	Logger *zap.Logger
}

func New(c *config.Config, logger *zap.Logger) *JiraRepository {
	return &JiraRepository{
		Db:     c.Database,
		Logger: logger,
	}
}
