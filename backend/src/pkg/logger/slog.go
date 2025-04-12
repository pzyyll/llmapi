package logger

import (
	"log/slog"
	"os"
	"sync"
)

const (
	Type     = "msg_type"
	SysType  = "sys"
	GromType = "gorm"
	RequestType = "request"
)

var (
	logLevel        = new(slog.LevelVar)
	typeLoggerCache sync.Map
)

func InitDefaultLogger() *slog.Logger {
	logLevel.Set(slog.LevelInfo)
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	})

	slog.SetDefault(slog.New(h))

	return slog.Default()
}

func ParseLevelString(lv string) (slog.Level, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(lv)); err != nil {
		return 0, err
	}
	return level, nil
}

func SetLevelString(lv string) error {
	level, err := ParseLevelString(lv)
	if err != nil {
		return err
	}
	SetLevel(level)
	return nil
}

func SetLevel(level slog.Level) {
	logLevel.Set(level)
}

func WithType(typ string) *slog.Logger {
	if loggerVal, ok := typeLoggerCache.Load(typ); ok {
		return loggerVal.(*slog.Logger)
	}
	logger := slog.With(slog.String(Type, typ))
	typeLoggerCache.Store(typ, logger)
	return logger
}

func Sys() *slog.Logger {
	return WithType(SysType)
}

func Get() *slog.Logger {
	return slog.Default()
}
