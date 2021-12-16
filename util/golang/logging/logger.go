package logging

import (
	"go.uber.org/zap"
	"sync"
)

func init() {
	_defaultLogger = New()
	logs[DefaultLoggerName] = _defaultLogger
}

// name for default loggers
const (
	DefaultLoggerName = "_default"
)

type Logger struct {
	*zap.SugaredLogger
}

var (
	_defaultLogger *Logger
)

var logs = map[string]*Logger{}
var logsMtx sync.RWMutex

func Log(name string) *Logger {
	logsMtx.RLock()
	defer logsMtx.RUnlock()
	return logs[name]
}

func New() *Logger {
	//logger, _ := zap.NewProduction()
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()

	return &Logger{
		SugaredLogger: sugar,
	}
}

func Debug(v ...interface{}) {
	_defaultLogger.Debug(v...)
}

func Info(v ...interface{}) {
	_defaultLogger.Info(v...)
}

func Warn(v ...interface{}) {
	_defaultLogger.Warn(v...)
}

func Error(v ...interface{}) {
	_defaultLogger.Error(v...)
}

func Debugf(format string, v ...interface{}) {
	_defaultLogger.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	_defaultLogger.Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	_defaultLogger.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	_defaultLogger.Errorf(format, v...)
}

// Sync all log data
func Sync() {
	logsMtx.RLock()
	defer logsMtx.RUnlock()
	for _, l := range logs {
		_ = l.Sync()
	}
}
