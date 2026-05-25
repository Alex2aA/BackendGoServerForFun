package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func Init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	var err error
	Log, err = config.Build(zap.AddCaller())
	if err != nil {
		panic(err)
	}
	Log.Info("User-Service logger initialized")
}

func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}
