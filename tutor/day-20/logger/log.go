package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// NewLogger ...
func NewLogger() *zap.Logger {
	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	log, err := config.Build()

	if err != nil {
		panic(err)
	}

	return log
}

// Info ...
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Warn ...
func Warn(message string, fields ...zap.Field) {
	log.Warn(message, fields...)
}

// Error ...
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
