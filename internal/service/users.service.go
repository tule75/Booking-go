package service

import (
	"ecommerce_go/internal/repo"
	"ecommerce_go/internal/utils/crypto"
	"ecommerce_go/internal/utils/random"
	"ecommerce_go/internal/utils/sendemail"
	"ecommerce_go/pkg/response"
	"fmt"
	"strconv"
)

type IUserService interface {
	Register(email string, password string) int
	GetUserInfo() string
}

type UserService struct {
	userRepo     repo.IUserRepo
	userAuthRepo repo.IUserAuthRepo
}

// Register implements IUserService.
func (us *UserService) Register(email string, password string) int {
	// 0. Hash email
	hashEmail, _ := crypto.GetHash(email)
	fmt.Printf("Email after hashing: %s\n", hashEmail)

	// 1. Hash password

	// 2. check email valid
	us.userRepo.GetUserByEmail(hashEmail)
	// 3. generate otp
	var otp = random.GenerateOTP()

	fmt.Printf("Otp is :::%d\n", otp)
	// 4. save otp in redis
	us.userAuthRepo.AddOTP(email, strconv.Itoa(otp), 300)
	// 5. send otp
	sendemail.SendTemplateEmailOtp([]string{email}, "ERP.nhacungcap1@gmail.com", "otp-auth.html", map[string]interface{}{
		"otp": strconv.Itoa(otp),
	})

	return response.SuccessResponseCode
}

func NewUserService(ur repo.IUserRepo, uar repo.IUserAuthRepo) IUserService {
	return &UserService{userRepo: ur, userAuthRepo: uar}
}

func (us *UserService) GetUserInfo() string {
	return us.userRepo.GetUserInfo()
}
