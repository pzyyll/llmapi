package log

import (
	"log/slog"
	"sync"

	"llmapi/src/pkg/logger"
)

const (
	Type        = "msg_type"
	SysType     = "sys"
	GormType    = "gorm"
	RequestType = "request"
)

var typeLoggerCache sync.Map

func WithType(typ string) *slog.Logger {
	if loggerVal, ok := typeLoggerCache.Load(typ); ok {
		return loggerVal.(*slog.Logger)
	}
	logger := logger.Get().With(slog.String(Type, typ))
	typeLoggerCache.Store(typ, logger)
	return logger
}

func Sys() *slog.Logger {
	return WithType(SysType)
}
