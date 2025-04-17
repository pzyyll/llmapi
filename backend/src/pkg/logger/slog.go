package logger

import (
	"log/slog"
	"os"
)

var logLevel = new(slog.LevelVar)

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

func Get() *slog.Logger {
	return slog.Default()
}
