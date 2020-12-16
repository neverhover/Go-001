package logger

import (
	"github.com/mainflux/mainflux/logger"
	"github.com/neverhover/Go-001/tree/main/Week04/internal/pkg/config"
	"io"
)

func GetLoggerLevel(config *config.Config) string{
	return config.LogLevel
}

func NewLoggerIns(out io.Writer, levelText string) (logger.Logger, error) {
	return logger.New(out, levelText)
}