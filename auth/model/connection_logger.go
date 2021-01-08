package model

import (
	"context"
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// DBLogger struct
type DBLogger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
}

// NewDBLogger func
func NewDBLogger() *DBLogger {
	return &DBLogger{
		SkipErrRecordNotFound: true,
	}
}

func (l *DBLogger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *DBLogger) Info(ctx context.Context, s string, args ...interface{}) {
	log.WithContext(ctx).Infof(s, args)
}

func (l *DBLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	log.WithContext(ctx).Warnf(s, args)
}

func (l *DBLogger) Error(ctx context.Context, s string, args ...interface{}) {
	log.WithContext(ctx).Errorf(s, args)
}

func (l *DBLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := log.Fields{}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[log.ErrorKey] = err
		log.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		log.WithContext(ctx).WithFields(fields).Warnf("%s [%s]", sql, elapsed)
		return
	}

	log.WithContext(ctx).WithFields(fields).Debugf("%s [%s]", sql, elapsed)
}
