package controller

import (
	requestDTO "ecommerce_go/internal/models/request"
	iservice "ecommerce_go/internal/service/interface"
	"ecommerce_go/pkg/response"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	VerifyOTP(c *gin.Context)
	PasswordRegister(c *gin.Context)
}

type UserController struct {
}

// Register implements IUserController.
func (uc *UserController) Register(c *gin.Context) {
	var userLogin = iservice.UserLogin()
	var param requestDTO.RegisterRequestModel
	if err := c.ShouldBindJSON(&param); err != nil {
		response.ErrorResponse(c, response.FalseEmailResponseCode, err.Error())
		return
	}
	var code, _ = userLogin.Register(c, &param)

	response.SuccessResponse(c, code, nil)
}

// Login
func (uc *UserController) Login(c *gin.Context) {
	var userLogin = iservice.UserLogin()
	var param requestDTO.LoginRequestModel
	if err := c.ShouldBindJSON(&param); err != nil {
		response.ErrorResponse(c, response.FalseEmailResponseCode, err.Error())
		return
	}
	var code, out, _ = userLogin.Login(c, &param)

	response.SuccessResponse(c, code, out)
}

func (uc *UserController) VerifyOTP(c *gin.Context) {
	var userLogin = iservice.UserLogin()
	var param requestDTO.VerifyRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		response.ErrorResponse(c, response.FalseEmailResponseCode, err.Error())
		return
	}
	var out, err = userLogin.VerifyOTP(c, &param)

	if err != nil {
		response.ErrorResponse(c, response.ErrorGenAuthCode, err.Error())
		return
	}
	response.SuccessResponse(c, response.SuccessResponseCode, out)
}

func (uc *UserController) PasswordRegister(c *gin.Context) {
	var userLogin = iservice.UserLogin()
	var param requestDTO.UserCreateRequestModel
	if err := c.ShouldBindJSON(&param); err != nil {
		response.ErrorResponse(c, response.FalseEmailResponseCode, err.Error())
		return
	}
	out, err := userLogin.UpdatePasswordRegister(c, &param)
	if err != nil {
		response.ErrorResponse(c, response.ErrorGenAuthCode, err.Error())
		return
	}
	response.SuccessResponse(c, response.SuccessResponseCode, out)
}

func NewUserController() IUserController {
	return &UserController{}
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
