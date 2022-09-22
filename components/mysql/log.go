package mysql

import (
	"fmt"
	"github.com/rs/zerolog"

	"gorm.io/gorm/utils"
	"time"
)

// from https://github.com/wei840222/gorm-zerolog/blob/main/logger.go
import (
	"context"
	"errors"
	gormlogger "gorm.io/gorm/logger"
)

type Logger struct {
	gormlogger.Config
	Logger                              zerolog.Logger
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func New(zl zerolog.Logger, config gormlogger.Config) *Logger {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	if config.Colorful {
		infoStr = gormlogger.Green + "%s\n" + gormlogger.Reset + gormlogger.Green + "[info] " + gormlogger.Reset
		warnStr = gormlogger.BlueBold + "%s\n" + gormlogger.Reset + gormlogger.Magenta + "[warn] " + gormlogger.Reset
		errStr = gormlogger.Magenta + "%s\n" + gormlogger.Reset + gormlogger.Red + "[error] " + gormlogger.Reset
		traceStr = gormlogger.Green + "%s\n" + gormlogger.Reset + gormlogger.Yellow + "[%.3fms] " + gormlogger.BlueBold + "[rows:%v]" + gormlogger.Reset + " %s"
		traceWarnStr = gormlogger.Green + "%s " + gormlogger.Yellow + "%s\n" + gormlogger.Reset + gormlogger.RedBold + "[%.3fms] " + gormlogger.Yellow + "[rows:%v]" + gormlogger.Magenta + " %s" + gormlogger.Reset
		traceErrStr = gormlogger.RedBold + "%s " + gormlogger.MagentaBold + "%s\n" + gormlogger.Reset + gormlogger.Yellow + "[%.3fms] " + gormlogger.BlueBold + "[rows:%v]" + gormlogger.Reset + " %s"
	}

	return &Logger{
		Logger:       zl,
		Config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

//func NewWithLogger(l zerolog.Logger, level gormlogger.LogLevel) *Logger {
//	return &Logger{
//		Logger:                    l,
//		LogLevel:                  level,
//		IgnoreRecordNotFoundError: true,
//	}
//}

func (l *Logger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *Logger) Info(ctx context.Context, s string, args ...interface{}) {
	l.Logger.Info().Msgf(s, args...)
}

func (l *Logger) Warn(ctx context.Context, s string, args ...interface{}) {
	l.Logger.Info().Msgf(s, args...)
}

func (l *Logger) Error(ctx context.Context, s string, args ...interface{}) {
	l.Logger.Info().Msgf(s, args...)
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	if l.LogLevel <= gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error && (!errors.Is(err, gormlogger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.Logger.Info().Msgf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)

		} else {
			l.Logger.Info().Msgf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.Logger.Info().Msgf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Logger.Info().Msgf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.LogLevel == gormlogger.Info:
		sql, rows := fc()
		if rows == -1 {
			l.Logger.Info().Msgf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Logger.Info().Msgf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
