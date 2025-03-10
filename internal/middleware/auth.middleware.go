package middleware

import (
	"ecommerce_go/internal/utils/auth"
	"ecommerce_go/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.UnauthorizedResponse(c, response.UnauthorizedResponseCode)
			return
		}

		// Lấy token từ header
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.UnauthorizedResponse(c, response.UnauthorizedResponseCode)
			return
		}

		token := parts[1]

		// Giải mã & xác thực token (cần thêm thư viện JWT)
		userID, err := auth.VerifyToken(token)
		if err != nil {
			response.UnauthorizedResponse(c, response.UnauthorizedResponseCode)
			return
		}

		// Lưu userID vào context
		c.Set("currentUser", userID)

		c.Next()
	}
}

func Authorization(allowedRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user, err = auth.GetCurrentUser(ctx)

		if err != nil {
			response.UnauthorizedResponse(ctx, response.UnauthorizedResponseCode)
			return
		}

		hasRole := false
		ur := string(user.Role.UsersRole)
		for _, ar := range allowedRoles {
			if strings.EqualFold(ur, ar) {
				hasRole = true
				break
			}
		}

		if !hasRole {
			response.UnauthorizedResponse(ctx, response.UnauthorizedResponseCode)
			return
		}

		ctx.Next()
	}
}
