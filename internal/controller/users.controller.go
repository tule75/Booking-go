package controller

import (
	dto "ecommerce_go/internal/models/DTO"
	"ecommerce_go/internal/service"
	"ecommerce_go/pkg/response"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	Register(c *gin.Context)
}

type UserController struct {
	userService service.IUserService
}

// Register implements IUserController.
func (uc *UserController) Register(c *gin.Context) {
	var param dto.RegisterRequestModel
	if err := c.ShouldBindJSON(&param); err != nil {
		response.ErrorResponse(c, response.FalseEmailResponseCode, err.Error())
		return
	}
	var code = uc.userService.Register(param.Email, param.Password)

	response.SuccessResponse(c, code, nil)
}

func NewUserController(us service.IUserService) IUserController {
	return &UserController{
		userService: us,
	}
}

// func (uc *UserController) GetUserInfo(c *gin.Context) {
// 	uInfo := uc.userService.GetUserInfo()
// 	c.JSON(200, gin.H{
// 		"message": uInfo,
// 	})
// }

// func (uc *UserController) CheckEmail(c *gin.Context) {
// 	email := c.Query("email")
// 	password := c.Query("password")
// 	if email == "" {
// 		response.BadResponse(c, response.FalseEmailResponseCode)
// 	} else if password == "" {
// 		response.BadResponse(c, response.FalsePasswordResponseCode)
// 	} else {
// 		response.SuccessResponse(c, response.SuccessResponseCode, map[string]string{"email": email, "password": password})
// 	}
// }
