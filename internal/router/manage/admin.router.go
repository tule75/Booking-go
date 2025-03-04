package manage

import (
	"ecommerce_go/internal/middleware"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (r *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	adminPublicRouter := Router.Group("/admin")
	{
		adminPublicRouter.GET("/login", handle)
	}

	adminPrivateRouter := Router.Group("/admin")
	adminPrivateRouter.Use(middleware.AuthenticationMiddleware())
	{
		adminPrivateRouter.GET("/:id", handle)
	}

}
