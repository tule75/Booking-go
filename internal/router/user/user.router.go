package user

import (
	"ecommerce_go/internal/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userPublicRouter := Router.Group("/users")
	{
		userPublicRouter.GET("/register", handle)
		userPublicRouter.GET("/login", handle)
		userPublicRouter.GET("/sendOTP", handle)
	}

	userPrivateRouter := Router.Group("/users")
	userPrivateRouter.Use(middleware.AuthenticationMiddleware())
	{
		userPrivateRouter.GET("/:id", handle)

	}
}

func handle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}
