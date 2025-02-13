package log

import (
	"fmt"
	"log/slog"
	"os"
)

var Logger *slog.Logger

func init() {
	logLevel := &slog.LevelVar{}
	opts := &slog.HandlerOptions{
        Level: logLevel,
    }
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	handler := getHandler(file, opts) 
	Logger = slog.New(handler)

	setLogLevel(logLevel)

	slog.Info("logger settings done")
	Logger.Info("logger settings done")
}

func getHandler(file *os.File, opts *slog.HandlerOptions) slog.Handler {
	switch os.Getenv("NOTES_LOG_TYPE") {
	case "json":
		return slog.NewJSONHandler(file, opts)
	default:
		return slog.NewTextHandler(file, opts)
	}
}

func setLogLevel(logLevel *slog.LevelVar) {
	switch os.Getenv("NOTES_LOG_LEVEL") {
	case "debug":
		logLevel.Set(slog.LevelDebug)
	case "info":
		logLevel.Set(slog.LevelInfo)
	case "warn":
		logLevel.Set(slog.LevelWarn)
	case "error":
		logLevel.Set(slog.LevelError)
	}
}

func Lab() {
	fmt.Println("b")
}