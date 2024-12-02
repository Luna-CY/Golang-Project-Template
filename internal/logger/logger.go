package logger

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/context/contextutil"
	"go.uber.org/zap"
)

func SugarLogger(ctx context.Context) *zap.SugaredLogger {
	return logger.Sugar().With("request_id", contextutil.GetRequestId(ctx))
}
