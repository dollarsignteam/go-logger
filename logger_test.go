package logger_test

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/dollarsignteam/go-logger"
)

func TestLogger(t *testing.T) {
	log := logger.NewLogger(logger.LoggerOptions{})
	log.Info("I am a info log")
}

func TestLoggerWithOptions(t *testing.T) {
	opts := logger.LoggerOptions{
		Level:      "debug",
		Name:       "DEMO",
		HideCaller: false,
	}
	log := logger.NewLogger(opts)
	log.Debug("I am a debug log")
	log.Info("I am a info log")
	log.Warn("I am a warn log")
	log.Error("I am a error log")
}

func TestToLevel(t *testing.T) {
	testcases := []struct {
		level    string
		expected zapcore.Level
	}{
		{level: "debug", expected: zap.DebugLevel},
		{level: "info", expected: zap.InfoLevel},
		{level: "warn", expected: zap.WarnLevel},
		{level: "error", expected: zap.ErrorLevel},
		{level: "dpanic", expected: zap.DPanicLevel},
		{level: "panic", expected: zap.PanicLevel},
		{level: "fatal", expected: zap.FatalLevel},
		{level: "", expected: zap.InfoLevel},
	}
	for _, tc := range testcases {
		if r := logger.ToLevel(tc.level); r != tc.expected {
			t.Fatalf("expect %v but got %v", tc.expected, r)
		}
	}
}
