package user

import (
	"ecommerce_go/internal/middleware"
	"ecommerce_go/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	uc, _ := wire.InitUserRouterHanlder()

	userPublicRouter := Router.Group("/users")
	{
		userPublicRouter.POST("/register", uc.Register)
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
