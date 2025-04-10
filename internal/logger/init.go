package logger

import (
	"os"

	"github.com/Luna-CY/Golang-Project-Template/internal/configuration"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func init() {
	var config = zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	var cores = make([]zapcore.Core, 0)
	for _, output := range configuration.Configuration.Logger.Outputs {
		switch output {
		case "stdout":
			cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.Lock(os.Stdout), coverZapLevel(configuration.Configuration.Logger.Level)))
		case "stderr":
			cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.Lock(os.Stderr), coverZapLevel(configuration.Configuration.Logger.Level)))
		case "feishu":
			cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(&FeishuWriter{}), coverZapLevel(configuration.Configuration.Logger.CustomizeWriter.Feishu.Level)))
		default:
			lumberjackWriter := &lumberjack.Logger{
				Filename:   output,
				MaxSize:    configuration.Configuration.Logger.MaxSize,
				MaxAge:     configuration.Configuration.Logger.MaxAge,
				MaxBackups: configuration.Configuration.Logger.MaxBackups,
				LocalTime:  true,
				Compress:   true,
			}

			cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(lumberjackWriter), coverZapLevel(configuration.Configuration.Logger.Level)))
		}
	}

	logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller())
}

func coverZapLevel(level string) zapcore.Level {
	var zapLevel = zap.PanicLevel

	switch level {
	case "debug":
		zapLevel = zap.DebugLevel
	case "info":
		zapLevel = zap.InfoLevel
	case "warn":
		zapLevel = zap.WarnLevel
	case "error":
		zapLevel = zap.ErrorLevel
	}

	return zapLevel
}
