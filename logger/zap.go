package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func NewZap() (*zap.Logger, func()) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := config.Build(zap.AddCaller())
	if err != nil {
		log.Fatal(err)
	}
	undo := zap.ReplaceGlobals(logger)

	return logger, func() {
		undo()
		logger.Sync()
	}
}
