package log

import (
	"log/slog"
	"sync"

	"llmapi/src/internal/constants"
	"llmapi/src/pkg/logger"

	"github.com/gin-gonic/gin"
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

func GetContextLogger(c *gin.Context) *slog.Logger {
	if logger, exists := c.Get(constants.ContextLoggerKey); exists {
		if reqLogger, ok := logger.(*slog.Logger); ok {
			return reqLogger
		}
	}
	return Sys()
}
