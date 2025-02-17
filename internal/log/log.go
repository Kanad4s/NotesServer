package log

import (
	"log/slog"
	"os"
)

var file *os.File

func Setup() {
	logLevel := &slog.LevelVar{}
	opts := &slog.HandlerOptions{
		AddSource: false,
		Level:     logLevel,
	}
	var err error
	file, err = os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	handler := getHandler(file, opts)
	Logger := slog.New(handler)
	slog.SetDefault(Logger)
	setLogLevel(logLevel)

	Logger.Debug("Log init done")
}

func Finish() {
	slog.SetDefault(slog.Default())
	file.Close()
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
	switch os.Getenv("NOTES_LOG_LVL") {
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

func TestAll() {
	slog.Debug("Test debug")
	slog.Error("Test error")
	slog.Warn("Test warn")
}

func TestDebug() {
	slog.Debug("Test debug")
	slog.SetDefault(slog.Default())
}

func TestError() {
	slog.Error("Test error")
}

func TestWarn() {
	slog.Warn("Test warn")
}
