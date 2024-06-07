package logger

import "go.uber.org/zap"

func New(env string) (logger *zap.Logger) {
	switch env {
	case "dev":
		logger, _ = zap.NewDevelopment()
	case "prod":
		logger, _ = zap.NewProduction()
	default:
		logger = zap.NewNop()
		logger.Warn("logger is set unknown environment")
	}
	return logger
}
