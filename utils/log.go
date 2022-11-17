package utils

import (
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func At() *zerolog.Logger {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	return &zlog.Logger
}
