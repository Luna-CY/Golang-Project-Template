package logger

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/runtime"
	"go.uber.org/zap"
)

func SugarLogger(ctx context.Context) *zap.SugaredLogger {
	return logger.Sugar().With("request_id", GetRequestId(ctx), "release", runtime.Release)
}

func GetRequestId(ctx context.Context) string {
	var value = ctx.Value("x-request-id")
	if nil != value {
		return value.(string)
	}

	return ""
}
