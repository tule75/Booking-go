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
	UserLogin iservice.IUserLogin
}

// Register implements IUserController.
// Register godoc
// @Summary      Register to get OTP
// @Description  Register to get OTP
// @Tags         Account management
// @Accept       json
// @Produce      json
// @Param        payload body requestDTO.RegisterRequestModel true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /users/register [POST]
func (uc *UserController) Register(c *gin.Context) {
	var param requestDTO.RegisterRequestModel
	if err := c.ShouldBindJSON(&param); err != nil {
		response.ErrorResponse(c, response.FalseEmailResponseCode, err.Error())
		return
	}
	var code, _ = uc.UserLogin.Register(c, &param)

	response.SuccessResponse(c, code, nil)
}

// Login
// @Summary      Login
// @Description  Login
// @Tags         Account management
// @Accept       json
// @Produce      json
// @Param        payload body requestDTO.LoginRequestModel true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ResponseData
// @Router       /users/login [post]
func (uc *UserController) Login(c *gin.Context) {
	var param requestDTO.LoginRequestModel
	if err := c.ShouldBindJSON(&param); err != nil {
		response.ErrorResponse(c, response.FalseEmailResponseCode, err.Error())
		return
	}
	var code, out, _ = uc.UserLogin.Login(c, &param)

	response.SuccessResponse(c, code, out)
}

// @Summary      Verify OTP
// @Description  Verify OTP
// @Tags         Account management
// @Accept       json
// @Produce      json
// @Param        payload body requestDTO.VerifyRequest true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ResponseData
// @Router       /users/verify-otp [post]
func (uc *UserController) VerifyOTP(c *gin.Context) {
	var param requestDTO.VerifyRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		response.ErrorResponse(c, response.FalseEmailResponseCode, err.Error())
		return
	}
	var out, err = uc.UserLogin.VerifyOTP(c, &param)

	if err != nil {
		response.ErrorResponse(c, response.ErrorGenAuthCode, err.Error())
		return
	}
	response.SuccessResponse(c, response.SuccessResponseCode, out)
}

// @Summary      Verify OTP
// @Description  Verify OTP
// @Tags         Account management
// @Accept       json
// @Produce      json
// @Param        payload body requestDTO.UserCreateRequestModel true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ResponseData
// @Router       /users/create-password [post]
func (uc *UserController) PasswordRegister(c *gin.Context) {
	var param requestDTO.UserCreateRequestModel
	if err := c.ShouldBindJSON(&param); err != nil {
		response.ErrorResponse(c, response.FalseEmailResponseCode, err.Error())
		return
	}
	out, code, err := uc.UserLogin.UpdatePasswordRegister(c, &param)
	if err != nil {
		response.ErrorResponse(c, code, err.Error())
		return
	}
	response.SuccessResponse(c, code, out)
}

func NewUserController(userLogin iservice.IUserLogin) IUserController {
	return &UserController{UserLogin: userLogin}
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
