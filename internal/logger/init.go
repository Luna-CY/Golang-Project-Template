package logger

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/configuration"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var logger *zap.Logger

func init() {
	var level = zap.PanicLevel

	switch configuration.Configuration.Logger.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	}

	var config = zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	var cores = make([]zapcore.Core, 0)
	for _, output := range configuration.Configuration.Logger.Outputs {
		switch output {
		case "stdout":
			cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.Lock(os.Stdout), level))
		case "stderr":
			cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.Lock(os.Stderr), level))
		default:
			lumberjackWriter := &lumberjack.Logger{
				Filename:   output,
				MaxSize:    configuration.Configuration.Logger.MaxSize,
				MaxAge:     configuration.Configuration.Logger.MaxAge,
				MaxBackups: configuration.Configuration.Logger.MaxBackups,
				LocalTime:  true,
				Compress:   true,
			}

			cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(lumberjackWriter), level))
		}
	}

	logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller())
}
