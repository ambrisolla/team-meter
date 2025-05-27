package main

import (
	"team-meter/config"
	"team-meter/internal/migrate"
	jiraScheduler "team-meter/internal/scheduler/jira"
	"team-meter/pkg/logger"

	"go.uber.org/zap"
)

func main() {

	cfg := config.Get()

	// Config logger
	logger, err := logger.Logger(cfg)
	if err != nil {
		panic("error starting logger: " + err.Error())
	}
	logger.Info("starting application")

	// Migration
	if err := migrate.Run(cfg.Database); err != nil {
		logger.Error("error executing migrations", zap.Error(err))
		panic("error executing migrations")
	}

	// Run schedulers
	go jiraScheduler.Run(cfg, logger)
	select {}

}
