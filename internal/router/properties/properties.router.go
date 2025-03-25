package properties

import (
	"ecommerce_go/internal/controller"
	"ecommerce_go/internal/middleware"
	iservice "ecommerce_go/internal/service/interface"
	constant "ecommerce_go/pkg"

	"github.com/gin-gonic/gin"
)

type PropertyRouter struct {
}

func (r *PropertyRouter) InitPropertiesRouter(Router *gin.RouterGroup) {
	uc := controller.NewPropertiesController(iservice.GetProperty())
	propertyPublicRouter := Router.Group("/properties")
	{
		propertyPublicRouter.GET("/owner/:id", uc.GetPropertiesByOwnerID)
		propertyPublicRouter.GET("/:id", uc.GetPropertyByID)
		propertyPublicRouter.POST("/filter", uc.SearchProperties)
	}

	propertyPrivateRouter := Router.Group("/properties")
	propertyPrivateRouter.Use(middleware.AuthenticationMiddleware())
	propertyPrivateRouter.Use(middleware.Authorization([]string{constant.RoleAdmin, constant.RoleHost}))
	{
		propertyPrivateRouter.GET("/owner/current_property", uc.GetPropertiesByOwnerID)
		propertyPrivateRouter.POST("/", uc.CreateProperty)
		propertyPrivateRouter.PUT("/:id", uc.UpdateProperty)
		propertyPrivateRouter.DELETE("/", uc.DeleteProperty)
	}
}
