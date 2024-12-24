package middleware

import (
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("X-Request-ID", gonanoid.MustID(21))
	}
}
