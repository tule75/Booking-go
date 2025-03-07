package user

import (
	"context"
	"ecommerce_go/global"
	"ecommerce_go/internal/controller"
	"ecommerce_go/internal/middleware"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type UserRouter struct {
}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//uc, _ := wire.InitUserRouterHanlder()
	uc := controller.NewUserController()
	userPublicRouter := Router.Group("/users")
	{
		userPublicRouter.POST("/register", uc.Register)
		userPublicRouter.POST("/login", uc.Login)
		userPublicRouter.POST("/verify-otp", uc.VerifyOTP)
		userPublicRouter.POST("/create-password", uc.PasswordRegister)
	}

	userPrivateRouter := Router.Group("/users")
	userPrivateRouter.Use(middleware.AuthenticationMiddleware())
	{
		userPrivateRouter.GET("/:id", handle)

	}
}

func handle(c *gin.Context) {
	var s = "test"
	body, _ := json.Marshal(s)

	km := kafka.Message{Key: []byte("test"), Value: body}

	err := global.KafkaProducer.WriteMessages(context.Background(), km)

	if err != nil {
		global.Logger.Error("error writing messages in kafka", zap.Error(err))
	}

	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}
