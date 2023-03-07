package sys_init

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"
)

// optionFunc wraps a func so it satisfies the Option interface.
// An Option configures a Logger.
type Config interface {
	apply(*loggerOptions)
}

type loggerOptions struct {
	zapOptions []zap.Option
}

// optionFunc wraps a func so it satisfies the Option interface.
type configFunc func(logger *loggerOptions)

func (f configFunc) apply(log *loggerOptions) {
	f(log)
}

func WrapOptions(opts []zap.Option) Config {
	return configFunc(func(l *loggerOptions) {
		l.zapOptions = opts
	})
}

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func InitAllZap(opts ...configFunc) {
	config := &loggerOptions{}
	for _, o := range opts {
		o.apply(config)
	}
	options := []zap.Option{}
	options = append(options, config.zapOptions...)
	tee := zapcore.NewTee(NewErrorCore(), NewInfoCore())
	logger := zap.New(tee, options...)
	zap.ReplaceGlobals(logger)
	zap.S().Infof("Init Zap Log success")
}

func NewInfoCore() zapcore.Core {
	return zapcore.NewCore(
		GetEncoder(),
		NewWarnWrite(),
		zap.InfoLevel,
	)
}

func NewErrorCore() zapcore.Core {
	return zapcore.NewCore(
		GetEncoder(),
		NewErrorWrite(),
		zap.ErrorLevel,
	)
}

func GetEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}

func NewErrorWrite() zapcore.WriteSyncer {
	l := &lumberjack.Logger{
		Filename:   "/go/src/spider-golang-web/logs/error.log",
		MaxSize:    512,
		MaxAge:     30,
		MaxBackups: 2,
		Compress:   false,
	}
	return zapcore.AddSync(l)
}

func NewWarnWrite() zapcore.WriteSyncer {
	l := &lumberjack.Logger{
		Filename:   "/go/src/spider-golang-web/logs/info.log",
		MaxSize:    512,
		MaxAge:     30,
		MaxBackups: 2,
		Compress:   false,
	}
	return zapcore.AddSync(l)
}
