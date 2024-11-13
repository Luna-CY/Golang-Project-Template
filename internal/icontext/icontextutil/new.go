package icontextutil

import (
	"context"
	"github.com/Luna-CY/Golang-Project-Template/internal/icontext"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid"
	"time"
)

func NewContext() icontext.Context {
	var ctx = context.WithValue(context.Background(), "x-request-id", gonanoid.MustID(64))

	return &icontext.IContext{Context: ctx}
}

func NewContextWithParent(parent context.Context) icontext.Context {
	return &icontext.IContext{Context: parent}
}

func NewContextWithValue(parent icontext.Context, key string, value any) icontext.Context {
	return &icontext.IContext{Context: context.WithValue(parent, key, value)}
}

func NewContextWithTimeout(parent context.Context, timeout time.Duration) (icontext.Context, context.CancelFunc) {
	var ctx, cancel = context.WithTimeout(parent, timeout)

	return &icontext.IContext{Context: ctx}, cancel
}

func NewContextWithGin(c *gin.Context) icontext.Context {
	var ctx = context.WithValue(context.Background(), "x-request-id", c.Value("X-Request-ID"))

	return &icontext.IContext{Context: ctx}
}
