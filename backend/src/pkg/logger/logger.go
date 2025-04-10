package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	log *logrus.Logger
}

var Log *Logger

func NewLogger() *Logger {
	return &Logger{
		log: logrus.New(),
	}
}

func (l *Logger) Info(message string, fields logrus.Fields) {
	l.log.WithFields(fields).Info(message)
}

func (l *Logger) SysInfo(message string, fields logrus.Fields) {
	l.log.WithFields(fields).WithFields(logrus.Fields{
		"type": "sys",
	}).Info(message)
}

func (l *Logger) SysError(message string, fields logrus.Fields) {
	l.log.WithFields(fields).WithFields(logrus.Fields{
		"type": "sys",
	}).Error(message)
}

func (l *Logger) SysDebug(message string, fields logrus.Fields) {
	l.log.WithFields(fields).WithFields(logrus.Fields{
		"type": "sys",
	}).Debug(message)
}

func (l *Logger) SysWarn(message string, fields logrus.Fields) {
	l.log.WithFields(fields).WithFields(logrus.Fields{
		"type": "sys",
	}).Warn(message)
}

func (l *Logger) SysFatal(message string, fields logrus.Fields) {
	l.log.WithFields(fields).WithFields(logrus.Fields{
		"type": "sys",
	}).Fatal(message)
}

func init() {
	// Initialize any middleware or configurations here if needed
	Log = NewLogger()
	Log.log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	Log.log.SetLevel(logrus.InfoLevel)
	Log.log.SetOutput(gin.DefaultWriter) // Set output to gin's default writer
}
