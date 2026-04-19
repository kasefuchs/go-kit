package log

import (
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var l = log.With().Logger()

// Init sets up the global logger
func Init(cfg Config) error {
	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(level)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	var output io.Writer = os.Stdout
	if cfg.Pretty {
		output = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: zerolog.TimeFieldFormat,
		}
	}

	l = zerolog.New(output).With().Timestamp().Caller().Logger()

	return nil
}

// Logger returns the global logger
func Logger() *zerolog.Logger {
	return &l
}

// Trace logs at trace level
func Trace() *zerolog.Event {
	return l.Trace()
}

// Debug logs at debug level
func Debug() *zerolog.Event {
	return l.Debug()
}

// Info logs at info level
func Info() *zerolog.Event {
	return l.Info()
}

// Warn logs at warn level
func Warn() *zerolog.Event {
	return l.Warn()
}

// Error logs at debug level
func Error() *zerolog.Event {
	return l.Error()
}

// Fatal logs at error level
func Fatal() *zerolog.Event {
	return l.Fatal()
}

// Panic logs at panic level
func Panic() *zerolog.Event {
	return l.Panic()
}
