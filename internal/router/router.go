package Router

import (
	"ecommerce_go/internal/controller"
	"ecommerce_go/internal/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.AuthenticationMiddleware())

	r.GET("/ping", Pong)

	v1 := r.Group("/v1/2025")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello world",
			})
		})

		v1.GET("/:name", GetName)
	}

	u := r.Group("/users")
	{
		u.GET("/info", controller.NewUserController().GetUserInfo)
		u.GET("/login", controller.NewUserController().CheckEmail)
	}

	return r
}

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetName(c *gin.Context) {
	name := c.Param("name")

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Tên: %s", name),
	})
}
