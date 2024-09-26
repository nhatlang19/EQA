package config

import (
	"context"
	"fmt"
	"time"

	"log"
	"os"

	"github.com/getsentry/sentry-go"
	"gorm.io/gorm/logger"
)

type CustomLogger struct {
	logger.Interface
}

const (
	reset   = "\033[0m"
	red     = "\033[31m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
)

func (c CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil:
		sql, rows := fc()
		sentry.CaptureMessage(fmt.Sprintf("%s[ERROR]%s %s | %s | %s | %d rows", red, reset, err, elapsed, sql, rows))
		log.Printf("%s[ERROR]%s %s | %s | %s | %d rows", red, reset, err, elapsed, sql, rows)

	case elapsed > 500*time.Millisecond:
		sql, rows := fc()
		sentry.CaptureMessage(fmt.Sprintf("%s[SLOW SQL >= 500ms]%s | %s | %s | %d rows", yellow, reset, elapsed, sql, rows))
		log.Printf("%s[SLOW SQL >= 500ms]%s | %s | %s | %d rows", yellow, reset, elapsed, sql, rows)
	default:
		c.Interface.Trace(ctx, begin, fc, err)
	}
}

func NewCustomLogger() CustomLogger {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 500 * time.Millisecond, // Slow SQL threshold
			LogLevel:      logger.Silent,          // Log level
			Colorful:      true,                   // Disable color
		},
	)

	return CustomLogger{newLogger}
}
