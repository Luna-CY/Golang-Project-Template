package contextutil

import "github.com/Luna-CY/Golang-Project-Template/internal/context"

func GetRequestId(ctx context.Context) string {
	var value = ctx.Value("x-request-id")
	if nil != value {
		return value.(string)
	}

	return ""
}
