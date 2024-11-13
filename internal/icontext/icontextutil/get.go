package icontextutil

import "github.com/Luna-CY/Golang-Project-Template/internal/icontext"

func GetRequestId(ctx icontext.Context) string {
	var value = ctx.Value("x-request-id")
	if nil != value {
		return value.(string)
	}

	return ""
}
