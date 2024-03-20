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

// Logger is ...
type Logger struct {
	*zerolog.Logger
}

// New is ...
func New() *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &Logger{
		&log,
	}
}

// ErrorStack is ...
func (l *Logger) ErrorStack(err error) {
	l.Error().Caller(1).Stack().Err(err).Send()
}

// Log is ...
//func Log() Logger {
//	return log
//}
