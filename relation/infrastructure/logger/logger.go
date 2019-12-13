package logger

import (
	"github.com/tinyhole/im/relation/domain/logger"
	"github.com/tinyhole/im/relation/infrastructure/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(conf *config.Config) logger.Logger {
	level := zap.InfoLevel
	level.Set(conf.LogLevel)
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   conf.LogFilePath,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		level,
	)
	logger := zap.New(core)
	return logger.Sugar()
}
