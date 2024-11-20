package logger

import (
	"fmt"
	"log/slog"
	"os"
)

var logger_ *slog.Logger

func SetupLogging(logFile string) error {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	handler := slog.NewJSONHandler(file, nil)
	logger_ = slog.New(handler)

	logger_.Info("Logging setup complete")

	return nil
}

func Info(msg string, args ...any) {
	logger_.Info(msg, args...)
}

func Error(msg string, args ...any) {
	logger_.Error(msg, args...)
}

func Warn(msg string, args ...any) {
	logger_.Warn(msg, args...)
}

func Debug(msg string, args ...any) {
	logger_.Debug(msg, args...)
}

func Infof(msg string, args ...any) {
	logger_.Info(fmt.Sprintf(msg, args...))
}