package logger

import (
	"fmt"
	"log/slog"
)

type GormLogger struct {
	logger *slog.Logger
}

func NewGormLogger() *GormLogger {
	return &GormLogger{
		logger: slog.With(slog.String(Type, GromType)),
	}
}

func (g *GormLogger) Printf(format string, args ...interface{}) {
	g.logger.Info(fmt.Sprintf(format, args...))
}
