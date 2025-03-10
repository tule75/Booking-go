package properties

import (
	"ecommerce_go/internal/controller"
	"ecommerce_go/internal/middleware"
	iservice "ecommerce_go/internal/service/interface"

	"github.com/gin-gonic/gin"
)

type PropertyRouter struct {
}

func (r *PropertyRouter) InitPropertiesRouter(Router *gin.RouterGroup) {
	uc := controller.NewPropertiesController(iservice.GetProperty())
	userPublicRouter := Router.Group("/properties")
	{
		userPublicRouter.GET("/owner/:id", uc.GetPropertiesByOwnerID)
		userPublicRouter.GET("/:id", uc.GetPropertyByID)
		userPublicRouter.POST("/filter", uc.SearchProperties)
	}

	userPrivateRouter := Router.Group("/users")
	userPrivateRouter.Use(middleware.AuthenticationMiddleware())
	userPrivateRouter.Use(middleware.Authorization([]string{"ADMIN", "HOST"}))
	{
		userPrivateRouter.GET("/owner/:id", uc.GetPropertiesByOwnerID)
		userPrivateRouter.GET("/:id", uc.GetPropertyByID)
		userPrivateRouter.POST("/filter", uc.SearchProperties)
		userPrivateRouter.POST("/", uc.CreateProperty)
		userPrivateRouter.PUT("/", uc.UpdateProperty)
	}
}
