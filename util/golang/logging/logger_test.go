package logging

import (
	"go.uber.org/zap"
	"testing"
)

func TestDev(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()
	_defaultLogger = &Logger{
		SugaredLogger: sugar,
	}

	Debug("Debug log")
	Info("Info log")
	Warn("Warn log")
	Error("Error log")
}

func TestPrd(t *testing.T) {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	_defaultLogger = &Logger{
		SugaredLogger: sugar,
	}

	Debug("Debug log") // 生产默认不打印 debug
	Info("Info log")
	Warn("Warn log")
	Error("Error log")
}
