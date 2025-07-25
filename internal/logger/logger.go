package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger(isDev bool) {
	// Use pretty console logging for dev, JSON for prod
	if isDev {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
