package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type Log struct {
	Log *logrus.Logger
	Service string
}

func NewLogger(service string) *Log {
	var level = logrus.DebugLevel

	return &Log{
		Log: &logrus.Logger{
			Level:     level,
			Out:       os.Stdout,
			Formatter: &logrus.JSONFormatter{},
		},
		Service: service,
	}
}

func (logger *Log) Debug(msg string, tags ...string) {
	if logger.Log.Level < logrus.DebugLevel {
		return
	}
	logger.Log.WithFields(logger.parseFields(tags...)).Debug(msg)
}

func (logger *Log) Info(msg string, tags ...string) {
	if logger.Log.Level < logrus.InfoLevel {
		return
	}
	logger.Log.WithFields(logger.parseFields(tags...)).Info(msg)
}

func (logger *Log) Error(msg string, err error, tags ...string) {
	if logger.Log.Level < logrus.ErrorLevel {
		return
	}
	msg = fmt.Sprintf("%s - ERROR - %v", msg, err)
	logger.Log.WithFields(logger.parseFields(tags...)).Error(msg)
}

func (logger *Log) parseFields(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))
	for _, tag := range tags {
		els := strings.Split(tag, ":")
		result[strings.TrimSpace(els[0])] = strings.TrimSpace(els[1])
	}
	result["service"] = logger.Service
	result["facility"] = logger.Service
	return result
}
