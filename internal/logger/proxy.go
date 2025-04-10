package logger

import "go.uber.org/zap"

type LoggerProxy struct {
	logger *zap.SugaredLogger
}

func (cls *LoggerProxy) Error(args ...any) {
	cls.logger.Error(args...)
}

func (cls *LoggerProxy) Errorf(template string, args ...any) {
	cls.logger.Errorf(template, args...)
}

func (cls *LoggerProxy) Errorw(msg string, keysAndValues ...any) {
	cls.logger.Errorw(msg, keysAndValues...)
}

func (cls *LoggerProxy) Info(args ...any) {
	cls.logger.Info(args...)
}

func (cls *LoggerProxy) Infof(template string, args ...any) {
	cls.logger.Infof(template, args...)
}

func (cls *LoggerProxy) Infow(msg string, keysAndValues ...any) {
	cls.logger.Infow(msg, keysAndValues...)
}

func (cls *LoggerProxy) Debug(args ...any) {
	cls.logger.Debug(args...)
}

func (cls *LoggerProxy) Debugf(template string, args ...any) {
	cls.logger.Debugf(template, args...)
}

func (cls *LoggerProxy) Debugw(msg string, keysAndValues ...any) {
	cls.logger.Debugw(msg, keysAndValues...)
}

func (cls *LoggerProxy) Warn(args ...any) {
	cls.logger.Warn(args...)
}

func (cls *LoggerProxy) Warnf(template string, args ...any) {
	cls.logger.Warnf(template, args...)
}

func (cls *LoggerProxy) Warnw(msg string, keysAndValues ...any) {
	cls.logger.Warnw(msg, keysAndValues...)
}

func (cls *LoggerProxy) Fatal(args ...any) {
	cls.logger.Fatal(args...)
}

func (cls *LoggerProxy) Fatalf(template string, args ...any) {
	cls.logger.Fatalf(template, args...)
}

func (cls *LoggerProxy) Fatalw(msg string, keysAndValues ...any) {
	cls.logger.Fatalw(msg, keysAndValues...)
}

func (cls *LoggerProxy) Panic(args ...any) {
	cls.logger.Panic(args...)
}

func (cls *LoggerProxy) Panicf(template string, args ...any) {
	cls.logger.Panicf(template, args...)
}

func (cls *LoggerProxy) Panicw(msg string, keysAndValues ...any) {
	cls.logger.Panicw(msg, keysAndValues...)
}
