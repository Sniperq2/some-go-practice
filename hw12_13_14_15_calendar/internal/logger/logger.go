package logger

import (
	"io"

	"github.com/rs/zerolog"
)

type LoggerIntface interface {
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
}

type Logger struct {
	logger *zerolog.Logger
}

func New(logWriter io.Writer, level string) *Logger {
	logger := zerolog.New(logWriter).With().Logger()
	return &Logger{logger: &logger}
}

func (l *Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Error().Msg(msg)
}

func (l *Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func (l *Logger) Warning(msg string) {
	l.logger.Warn().Msg(msg)
}
