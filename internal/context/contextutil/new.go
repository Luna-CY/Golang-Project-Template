package contextutil

import (
	context2 "context"
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid"
	"time"
)

func NewContext() context.Context {
	var ctx = context2.WithValue(context2.Background(), "x-request-id", gonanoid.MustID(64))

	return &context.IContext{Context: ctx}
}

func NewContextWithParent(parent context.Context) context.Context {
	return &context.IContext{Context: parent}
}

func NewContextWithValue(parent context.Context, key string, value any) context.Context {
	return &context.IContext{Context: context2.WithValue(parent, key, value)}
}

func NewContextWithTimeout(parent context.Context, timeout time.Duration) (context.Context, context2.CancelFunc) {
	var ctx, cancel = context2.WithTimeout(parent, timeout)

	return &context.IContext{Context: ctx}, cancel
}

func NewContextWithGin(c *gin.Context) context.Context {
	var ctx = context2.WithValue(context2.Background(), "x-request-id", c.Value("X-Request-ID"))

	return &context.IContext{Context: ctx}
}
