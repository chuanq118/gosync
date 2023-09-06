package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: true,
		Encoding:          "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			MessageKey:    "msg",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime:    zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	log = zap.Must(cfg.Build())
}

// Logger returns the global logger instance
func Logger() *zap.Logger {
	return log
}

func Info(args ...interface{}) {
	log.Sugar().Info(args)
}

func Infoln(args ...interface{}) {
	log.Sugar().Infoln(args)
}

func Infof(template string, args ...interface{}) {
	log.Sugar().Infof(template, args)
}

func Infow(template string, kvs ...interface{}) {
	log.Sugar().Infow(template, kvs)
}

func Debug(args ...interface{}) {
	log.Sugar().Debug(args...)
}

func Debugln(args ...interface{}) {
	log.Sugar().Debugln(args...)
}

func Debugf(template string, args ...interface{}) {
	log.Sugar().Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	log.Sugar().Debugw(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	log.Sugar().Warn(args...)
}

func Warnln(args ...interface{}) {
	log.Sugar().Warnln(args...)
}

func Warnf(template string, args ...interface{}) {
	log.Sugar().Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	log.Sugar().Warnw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	log.Sugar().Error(args...)
}

func Errorln(args ...interface{}) {
	log.Sugar().Errorln(args...)
}

func Errorf(template string, args ...interface{}) {
	log.Sugar().Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	log.Sugar().Errorw(msg, keysAndValues...)
}

func Panic(args ...interface{}) {
	log.Sugar().Panic(args...)
}

func Panicln(args ...interface{}) {
	log.Sugar().Panicln(args...)
}

func Panicf(template string, args ...interface{}) {
	log.Sugar().Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	log.Sugar().Panicw(msg, keysAndValues...)
}

func Fatal(args ...interface{}) {
	log.Sugar().Fatal(args)
}

func Fatalln(args ...interface{}) {
	log.Sugar().Fatalln(args)
}

func Fatalf(template string, args ...interface{}) {
	log.Sugar().Fatalf(template, args)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	log.Sugar().Fatalw(msg, keysAndValues)
}
