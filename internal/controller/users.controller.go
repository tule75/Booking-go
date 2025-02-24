package controller

import (
	"ecommerce_go/internal/service"
	"ecommerce_go/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserInfo(c *gin.Context) {
	uInfo := uc.userService.GetUserInfo()
	c.JSON(200, gin.H{
		"message": uInfo,
	})
}

func (uc *UserController) CheckEmail(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")
	if email == "" {
		response.BadResponse(c, response.FalseEmailResponseCode)
	} else if password == "" {
		response.BadResponse(c, response.FalsePasswordResponseCode)
	} else {
		response.SuccessResponse(c, response.SuccessResponseCode, map[string]string{"email": email, "password": password})
	}
}
