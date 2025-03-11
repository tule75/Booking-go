package initialize

import (
	"ecommerce_go/global"
	Router "ecommerce_go/internal/router"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	manageRouter := Router.RouterAppGroup.AdminGroupRouter
	userRouter := Router.RouterAppGroup.UserGroupRouter
	propertiesRouter := Router.RouterAppGroup.PropertiesGroupRouter

	MainGroup := r.Group("/v1/2025")
	{
		userRouter.InitUserRouter(MainGroup)
	}
	{
		manageRouter.InitUserRouter(MainGroup)
		manageRouter.InitAdminRouter(MainGroup)
	}
	{
		propertiesRouter.InitPropertiesRouter(MainGroup)
	}

	return r
}
