package logger

import (
	"os"
	"sync"

	"github.com/rs/zerolog"

	"shirikaru/config"
)

type Logger struct {
	*zerolog.Logger
}

var (
	logger *Logger
	once   sync.Once
)

func GetLogger(cfg *config.Config) *Logger {
	once.Do(func() {
		loggingLevel := cfg.Server.LoggingLevel
		zeroLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()

		switch loggingLevel {
		case "debug":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "error":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "fatal":
			zerolog.SetGlobalLevel(zerolog.FatalLevel)
		case "panic":
			zerolog.SetGlobalLevel(zerolog.PanicLevel)
		default:
			zerolog.SetGlobalLevel(zerolog.NoLevel)
		}
		logger = &Logger{&zeroLogger}
	})

	return logger
}
