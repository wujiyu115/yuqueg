package yuqueg

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// L logger
var (
	L    *zap.SugaredLogger
	atom zap.AtomicLevel
)

func init() {
	L = NewLogger()
}

//NewLogger of zap
func NewLogger() *zap.SugaredLogger {
	if L != nil {
		return L
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// log level
	atom = zap.NewAtomicLevelAt(zap.InfoLevel)

	config := zap.Config{
		Level:            atom,
		Development:      true,
		Encoding:         "console", // console or json
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout", "./yuque.log"}, // log file
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("log init fail: %v", err))
	}
	return logger.Sugar()
}

//SetDebugLevel set debug level
func SetDebugLevel() {
	atom.SetLevel(zap.DebugLevel)
}
