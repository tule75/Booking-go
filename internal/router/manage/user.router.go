package manage

import (
	"ecommerce_go/internal/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// userPublicRouter := Router.Group("/admin/user")
	// {
	// 	userPublicRouter.GET("/register")
	// 	userPublicRouter.GET("/login")
	// 	userPublicRouter.GET("/sendOTP")
	// }

	userPrivateRouter := Router.Group("/admin/user")
	userPrivateRouter.Use(middleware.AuthenticationMiddleware())
	{
		userPrivateRouter.POST("/active", handle)
		userPrivateRouter.POST("/info", handle)
	}
}

func handle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}
