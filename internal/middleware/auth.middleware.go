package middleware

import (
	"ecommerce_go/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.BadResponse(c, response.InvalidTokenResponseCode)
			c.Abort()
			return
		}
		c.Next()
	}
}
