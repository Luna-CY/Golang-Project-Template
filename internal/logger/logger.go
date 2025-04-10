package logger

import (
	"runtime/debug"

	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/runtime"
	"go.uber.org/zap"
)

type Option func(*zap.SugaredLogger) *zap.SugaredLogger

func WithStack() Option {
	return func(logger *zap.SugaredLogger) *zap.SugaredLogger {
		return logger.With("stack", string(debug.Stack()))
	}
}

func SugarLogger(ctx context.Context, options ...Option) *zap.SugaredLogger {
	var logger = logger.Sugar().With("request_id", GetRequestId(ctx), "release", runtime.Release, "env", runtime.GetEnvironment())

	for _, option := range options {
		logger = option(logger)
	}

	return logger
}

func GetRequestId(ctx context.Context) string {
	var value = ctx.Value("x-request-id")
	if nil != value {
		return value.(string)
	}

	return ""
}
