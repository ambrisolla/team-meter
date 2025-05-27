package logger

import (
	"team-meter/config"

	"go.uber.org/zap"
)

func Logger(config *config.Config) (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err.Error())
	}
	logger = logger.With(
		zap.String("service", config.App.Name),
		zap.String("version", config.App.Version),
	)
	return logger, nil
}
