package logger

import (
	"fmt"
	"log/slog"
)

type GormLogger struct {
	logger *slog.Logger
}

func NewGormLogger(logger *slog.Logger) *GormLogger {
	return &GormLogger{
		logger: logger,
	}
}

func (g *GormLogger) Printf(format string, args ...interface{}) {
	g.logger.Info(fmt.Sprintf(format, args...))
}
