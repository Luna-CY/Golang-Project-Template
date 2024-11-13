package middleware

import "github.com/gin-gonic/gin"

func CustomGinRecovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		c.AbortWithStatusJSON(500, gin.H{
			"code": 500,
			"msg":  "InternalServerError",
		})
	})
}
