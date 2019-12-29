package log

import (
	"github.com/zhulinwei/go-dc/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var logger *zap.Logger

func Debug(msg string, fields ...zapcore.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields zapcore.Field) {
	logger.Error(msg, fields)
}

func String(key, value string) zapcore.Field {
	return zap.String(key, value)
}

func Reflect(key string, value interface{}) zapcore.Field {
	return zap.Reflect(key, value)
}

func init() {
	logConfig := zap.NewProductionConfig()
	// Error级别以下是否显示堆栈信息
	logConfig.DisableCaller = config.ServerConfig().Log.DisableCaller
	// 设置日志打印级别
	logConfig.Level.SetLevel(zapcore.Level(config.ServerConfig().Log.Level))

	var err error
	if logger, err = logConfig.Build(); err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
}
