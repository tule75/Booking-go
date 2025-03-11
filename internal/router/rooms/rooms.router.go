package rooms

import (
	"ecommerce_go/internal/controller"
	"ecommerce_go/internal/middleware"
	iservice "ecommerce_go/internal/service/interface"

	"github.com/gin-gonic/gin"
)

type RoomsRouter struct {
}

func (r *RoomsRouter) InitPropertiesRouter(Router *gin.RouterGroup) {
	rc := controller.NewRoomController(iservice.GetRoom())
	userPublicRouter := Router.Group("/rooms")
	{
		userPublicRouter.GET("/property/:id", rc.ListRoomsByProperty)
		userPublicRouter.GET("/:id", rc.GetRoomById)
	}

	userPrivateRouter := Router.Group("/users")
	userPrivateRouter.Use(middleware.AuthenticationMiddleware())
	userPrivateRouter.Use(middleware.Authorization([]string{"ADMIN", "HOST"}))
	{
		userPrivateRouter.POST("/", rc.CreateRoom)
		userPrivateRouter.PUT("/", rc.UpdateRoom)
		userPrivateRouter.DELETE("/", rc.DeleteRoom)
	}
}
