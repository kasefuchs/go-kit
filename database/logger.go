package database

import (
	"context"
	"time"

	"codeberg.org/kasefuchs/go-kit/log"
	"github.com/rs/zerolog"
	"gorm.io/gorm/logger"
)

var _ logger.Interface = (*gormLogger)(nil)

type gormLogger struct {
	log zerolog.Logger
}

func newGormLogger() gormLogger {
	logger := log.Logger().With().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + 1).Logger()

	return gormLogger{logger}
}

func (l gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l gormLogger) Info(ctx context.Context, msg string, data ...any) {
	l.log.Info().Msgf(msg, data...)
}

func (l gormLogger) Warn(ctx context.Context, msg string, data ...any) {
	l.log.Warn().Msgf(msg, data...)
}

func (l gormLogger) Error(ctx context.Context, msg string, data ...any) {
	l.log.Error().Msgf(msg, data...)
}

func (l gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	elapsed := time.Since(begin)

	if err != nil {
		l.log.Trace().Err(err).Dur("elapsed", elapsed).Int64("rows", rows).Msg(sql)
		return
	}

	l.log.Trace().Dur("elapsed", elapsed).Int64("rows", rows).Msg(sql)
}
