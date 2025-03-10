package iservice

import (
	"context"
	requestDTO "ecommerce_go/internal/models/request"
	responseDTO "ecommerce_go/internal/models/response"
)

type IUserLogin interface {
	Register(ctx context.Context, request *requestDTO.RegisterRequestModel) (int, error)
	VerifyOTP(ctx context.Context, in *requestDTO.VerifyRequest) (out responseDTO.VerifyOTPResponse, err error)
	Login(ctx context.Context, request *requestDTO.LoginRequestModel) (int, responseDTO.LoginResponse, error)
	UpdatePasswordRegister(ctx context.Context, in *requestDTO.UserCreateRequestModel) (userId string, code int, err error)
}

type IUserCustomer interface{}

type IUserHost interface{}

type IUserAdmin interface{}

var (
	localUserLogin    IUserLogin
	localUserCustomer IUserCustomer
	localUserHost     IUserHost
	localUserAdmin    IUserAdmin
)

func InitUserLogin(ul IUserLogin) {
	localUserLogin = ul
}

func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("implement UserLogin not found for interface IUserLogin")
	}
	return localUserLogin
}

func InitUserCustomer(ul IUserCustomer) {
	localUserCustomer = ul
}

func UserCustomer() IUserCustomer {
	if localUserCustomer == nil {
		panic("implement UserCustomer not found for interface IUserCustomer")
	}
	return localUserCustomer
}

func InitlocalUserHost(ul IUserHost) {
	localUserHost = ul
}

func UserHost() IUserHost {
	if localUserHost == nil {
		panic("implement localUserHost not found for interface IlocalUserHost")
	}
	return localUserHost
}

func InitUserAdmin(ul IUserAdmin) {
	localUserAdmin = ul
}

func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("implement UserLogin not found for interface IUserAdmin")
	}
	return localUserAdmin
}
