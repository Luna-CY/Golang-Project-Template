package logger

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/icontext"
	"github.com/Luna-CY/Golang-Project-Template/internal/icontext/icontextutil"
	"go.uber.org/zap"
)

func SugarLogger(ctx icontext.Context) *zap.SugaredLogger {
	return logger.Sugar().With("request_id", icontextutil.GetRequestId(ctx))
}
