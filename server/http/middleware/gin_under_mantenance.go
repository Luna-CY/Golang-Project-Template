package middleware

import "github.com/gin-gonic/gin"

func UnderMaintenance() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(501, gin.H{
			"code": 501,
			"msg":  "UnderMaintenance",
		})
	}
}
