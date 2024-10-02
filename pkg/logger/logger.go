package logger

import (
	"context"
	"log/slog"
	"os"
)

type Loggers struct {
	Info  *slog.Logger
	Debug *slog.Logger
}

func NewLoggers() (*Loggers, error) {
	infoHandler, err := newFileHandler("./logs/info.json")
	if err != nil {
		return nil, err
	}

	debugHandler, err := newFileHandler("./logs/debug.json")
	if err != nil {
		return nil, err
	}

	infoLogger := slog.New(infoHandler)
	debugLogger := slog.New(debugHandler)

	return &Loggers{
		Info:  infoLogger,
		Debug: debugLogger,
	}, nil
}

func newFileHandler(filePath string) (*slog.JSONHandler, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return slog.NewJSONHandler(file, nil), nil
}

func (l *Loggers) InfoLog(msg string, attrs ...slog.Attr) {
	l.Info.LogAttrs(context.Background(), slog.LevelInfo, msg, attrs...)
}

func (l *Loggers) DebugLog(msg string, attrs ...slog.Attr) {
	l.Debug.LogAttrs(context.Background(), slog.LevelDebug, msg, attrs...)
}
