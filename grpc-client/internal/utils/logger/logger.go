package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

// InitLogger initializes the logger
func InitLogger() (*zap.SugaredLogger, error) {
	// Configure Zap logger
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Initialize logger
	logger, err := config.Build()
	if err != nil {
		fmt.Println("Unable to initialize logging. Using default logging.")
		sugarLogger = zap.NewExample().Sugar()
		return sugarLogger, nil
	}

	// Replace logger with a SugaredLogger for structured, human-readable logging
	sugarLogger = logger.Sugar()

	return sugarLogger, nil
}

// GetLogger retrieves the logger instance
func GetLogger() *zap.SugaredLogger {
	return sugarLogger
}

// Info logs an info message
func Info(args ...interface{}) {
	sugarLogger.Info(args...)
}

// Warn logs a warning message
func Warn(args ...interface{}) {
	sugarLogger.Warn(args...)
}

// Debug logs a debug message
func Debug(args ...interface{}) {
	sugarLogger.Debug(args...)
}

// Error logs an error message
func Error(args ...interface{}) {
	sugarLogger.Error(args...)
}

// Infof logs a formatted info message
func Infof(format string, args ...interface{}) {
	sugarLogger.Infof(format, args...)
}

// Warnf logs a formatted warning message
func Warnf(format string, args ...interface{}) {
	sugarLogger.Warnf(format, args...)
}

// Debugf logs a formatted debug message
func Debugf(format string, args ...interface{}) {
	sugarLogger.Debugf(format, args...)
}

// Errorf logs a formatted error message
func Errorf(format string, args ...interface{}) {
	sugarLogger.Errorf(format, args...)
}
