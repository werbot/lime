package logging

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

/*
var log zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	log = zerolog.New(os.Stderr).With().Timestamp().Logger()
}

func Log() zerolog.Logger {
	return log
}
*/

type Log struct {
	*zerolog.Logger
}

func New() *Log {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &Log{
		&log,
	}
}

func (l *Log) ErrorStack(err error) {
	l.Error().Caller(1).Stack().Err(err).Send()
}
