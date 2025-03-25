package rooms

import (
	"ecommerce_go/internal/controller"
	"ecommerce_go/internal/middleware"
	iservice "ecommerce_go/internal/service/interface"
	constant "ecommerce_go/pkg"

	"github.com/gin-gonic/gin"
)

type RoomsRouter struct {
}

func (r *RoomsRouter) InitRoomsRouter(Router *gin.RouterGroup) {
	rc := controller.NewRoomController(iservice.GetRoom())
	roomPublicRouter := Router.Group("/rooms")
	{
		roomPublicRouter.GET("/property/:id", rc.ListRoomsByProperty)
		roomPublicRouter.GET("/:id", rc.GetRoomById)
	}

	roomPrivateRouter := Router.Group("/rooms")
	roomPrivateRouter.Use(middleware.AuthenticationMiddleware())
	roomPrivateRouter.Use(middleware.Authorization([]string{constant.RoleAdmin, constant.RoleHost}))
	{
		roomPrivateRouter.POST("/", rc.CreateRoom)
		roomPrivateRouter.PUT("/:id", rc.UpdateRoom)
		roomPrivateRouter.DELETE("/", rc.DeleteRoom)
	}
}
