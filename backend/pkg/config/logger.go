package config

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
)

// SetupLogger sets up the logger.
func SetupLogger() {
	standardHandler := handler.MustRotateFile("./logs/system.log", 60*60*12)
	errorHandler := handler.MustRotateFile("./logs/system_err.log", 60*60*12, handler.WithLogLevels(slog.DangerLevels))

	slog.Info("Setting up the Logger...")

	slog.PushHandler(standardHandler)
	slog.PushHandler(errorHandler)

	slog.Configure(func(logger *slog.SugaredLogger) {
		f := logger.Formatter.(*slog.TextFormatter)
		f.EnableColor = true
	})

	slog.Info("Logger setup successfully.")
}

// Info logs an info message.
func Info(message string) {
	slog.Info(message)
	slog.Flush()
}

// Fatal logs a fatal message.
func Fatal(message string) {
	slog.Fatal(message)
	slog.Flush()
}

// Error logs an error message.
func Error(message string) {
	slog.Error(message)
	slog.Flush()
}

// Warn logs a warning message.
func Warn(message string) {
	slog.Warn(message)
	slog.Flush()
}
