package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var l *zap.SugaredLogger

func init() {
	l = NewLogger()
}

func NewLogger() *zap.SugaredLogger {
    if l != nil {
        return l
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
    atom := zap.NewAtomicLevelAt(zap.DebugLevel)

    config := zap.Config{
        Level:            atom,
        Development:      true,
        Encoding:         "console",                                              // console or json
        EncoderConfig:    encoderConfig,
        OutputPaths:      []string{"stdout", "./yuque.log"},         // log file
        ErrorOutputPaths: []string{"stderr"},
    }

    logger, err := config.Build()
    if err != nil {
        panic(fmt.Sprintf("log init fail: %v", err))
    }
    return logger.Sugar()
}
