package logger

import (
	"context"
	"runtime/debug"

	"go.uber.org/zap"
)

type Option func(*zap.SugaredLogger) *zap.SugaredLogger

func WithRequestId(ctx context.Context) Option {
	return func(logger *zap.SugaredLogger) *zap.SugaredLogger {
		var value = ctx.Value("x-request-id")
		if nil != value {
			return logger.With("request_id", value.(string))
		}

		return logger
	}
}

func WithStack() Option {
	return func(logger *zap.SugaredLogger) *zap.SugaredLogger {
		return logger.With("stack", string(debug.Stack()))
	}
}
