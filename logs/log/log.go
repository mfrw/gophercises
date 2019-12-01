package log

import "go.uber.org/zap"

type Logger interface {
	Log(kv ...interface{})
}

type zapSugarLogger func(msg string, kvs ...interface{})

var sloger zapSugarLogger

func init() {
	l, _ := zap.NewProduction()
	sugar := l.WithOptions(zap.AddCallerSkip(2)).Sugar()
	sloger = sugar.Infof
}

func (z zapSugarLogger) Log(kv ...interface{}) {
	z("", kv...)
}

func NewImpl() Logger {
	l, err := zap.NewProduction()
	if err != nil {
		return nil
	}
	sugar := l.WithOptions(zap.AddCallerSkip(1)).Sugar()
	var sloger zapSugarLogger
	sloger = sugar.Infof
	return sloger
}

func Log(kv ...interface{}) {
	sloger.Log(kv...)
}
