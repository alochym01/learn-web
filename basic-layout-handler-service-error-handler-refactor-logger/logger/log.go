package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func NewLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	enConfig := zap.NewProductionEncoderConfig()
	enConfig.TimeKey = "timestamp"
	// enConfig.StacktraceKey = "" // dont include stackstrace
	enConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = enConfig
	config.DisableStacktrace = true // dont include stackstrace
	// log, _ := config.Build(zap.AddCallerSkip(1))
	log, _ := config.Build()
	return log
}

func Info(msg string, fields ...zap.Field) {
	// func Info(msg string) {
	log.Info(msg)
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg)
}
