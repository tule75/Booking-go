package main

import (
	_ "ecommerce_go/cmd/swag/docs"
	"ecommerce_go/global"
	"ecommerce_go/internal/initialize"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           Booking Go
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://github.com/tule75/Booking-Go

// @contact.name   Tú Lê
// @contact.url    http://github.com/tule75/Booking-Go
// @contact.email  tulehd03@gmail.com

// @license.name  Booking Go v1.0
// @license.url   http://github.com/tule75/Booking-Go

// @host      localhost:8081
// @BasePath  /v1/2025

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := initialize.Run()

	r.GET("/", func(c *gin.Context) { c.Redirect(http.StatusFound, "/swagger/index") })
	r.GET("/swagger", func(c *gin.Context) { c.Redirect(http.StatusFound, "/swagger/index") })
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(fmt.Sprintf(":%v", global.Config.Server.Port))

}
