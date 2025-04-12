package logger

import (
	"testing"

	"log/slog"
)

func TestLog(t *testing.T) {
	logger := InitDefaultLogger()
	logger.Info("This is an info log")
	logger.Debug("This is a debug log")
	SetLevelString("DEBUG")
	logger.Debug("This is a debug log after setting level to DEBUG")
	SetLevelString("info")
	logger.Debug("This debug log should not be shown")

	logger.Info("Mess", "key", "value")
	logger.Info("With Key", slog.String("field0", "value0"), slog.Int("field1", 1))

	sys := logger.With(slog.String("type", "sys"))
	sys.Info("Sys message")
	sys.Error("Sys error")
}