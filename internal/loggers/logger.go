package loggers

import (
	"os"

	"github.com/Reshnyak/AcornsTask/internal/configs"
	"github.com/rs/zerolog"
)

var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

func Log() *zerolog.Logger {
	return &logger
}
func InitLogger(cfg *configs.Config) *zerolog.Logger {
	logger.Level(cfg.ZeroLogLevel())
	return &logger
}
