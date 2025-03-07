package service

import (
	"context"
	"ecommerce_go/global"
	"ecommerce_go/internal/database"
	requestDTO "ecommerce_go/internal/models/request"
	responseDTO "ecommerce_go/internal/models/response"
	iservice "ecommerce_go/internal/service/interface"
	"ecommerce_go/internal/utils/auth"
	"ecommerce_go/internal/utils/crypto"
	"ecommerce_go/internal/utils/random"
	"ecommerce_go/internal/utils/sendemail"
	"ecommerce_go/pkg/response"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type UserLogin struct {
	sqlc *database.Queries
}

// Register implements IUserLogin.
func (ul *UserLogin) Register(ctx context.Context, request *requestDTO.RegisterRequestModel) (int, error) {
	// 0. Hash email
	hashEmail, _ := crypto.GetHash(request.Email)
	fmt.Printf("Email after hashing: %s\n", hashEmail)

	// 2. check email valid
	isExist, _ := ul.sqlc.CheckUserExists(ctx, hashEmail)

	if isExist > 0 {
		return response.ExistEmailResponseCode, nil
	}
	// 3. generate otp
	var otp = random.GenerateOTP()

	fmt.Printf("Otp is :::%d\n", otp)
	// 4. save otp in redis
	var index = fmt.Sprintf("usr:%s:otp", hashEmail)
	global.Rdb.SetEx(ctx, index, otp, time.Duration(300)*time.Second).Err() // 5 minutes
	// 5. send otp
	sendemail.SendTemplateEmailOtp([]string{request.Email}, "ERP.nhacungcap1@gmail.com", "otp-auth.html", map[string]interface{}{
		"otp": strconv.Itoa(otp),
	})

	return response.SuccessResponseCode, nil
}

func (ul *UserLogin) VerifyOTP(ctx context.Context, in *requestDTO.VerifyRequest) (out responseDTO.VerifyOTPResponse, err error) {
	// logic
	hashKey, _ := crypto.GetHash(strings.ToLower(in.VerifyKey))

	// get otp
	var index = fmt.Sprintf("usr:%s:otp", hashKey)
	fmt.Println(index)
	otpFound, err := global.Rdb.Get(ctx, index).Result()
	if err != nil {
		fmt.Println(err)
		return out, err
	}

	if in.VerifyCode != otpFound {
		return out, fmt.Errorf("OTP not match")
	}

	// output
	out.Token = hashKey
	out.Message = "success"

	return out, err
}

func (s *UserLogin) UpdatePasswordRegister(ctx context.Context, in *requestDTO.UserCreateRequestModel) (userId int, err error) {
	hashKey, _ := crypto.GetHash(strings.ToLower(in.Email))

	if hashKey != in.Token {
		return 0, fmt.Errorf("wrong token or email to set password")
	}

	user := database.CreateUserParams{}
	user.ID = uuid.New().String()
	user.Email = in.Email
	userSalt, err := crypto.GenerateSalt(16)
	if err != nil {
		return response.ErrCodeAuthFailed, err
	}
	user.Salt = userSalt
	user.Password = crypto.HashPassword(in.Password, userSalt)
	user.FullName = in.FullName
	user.Role = database.NullUsersRole{UsersRole: database.UsersRole(in.Role)}
	// add userBase to user_base table
	newUser, err := s.sqlc.CreateUser(ctx, user)
	log.Println("new User::", newUser, user)
	if err != nil {
		return response.CannotCreateUser, err
	}
	user_id, err := newUser.LastInsertId()
	if err != nil {
		return response.CannotCreateUser, err
	}

	return int(user_id), nil
}

func (ul UserLogin) Login(ctx context.Context, request *requestDTO.LoginRequestModel) (rCode int, out responseDTO.LoginResponse, e error) {
	user, err := ul.sqlc.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return response.FalseEmailResponseCode, out, errors.New("user not found")
	}

	if r := crypto.MatchingPassword(user.Password, request.Password, user.Salt); !r {
		return response.FalsePasswordResponseCode, out, errors.New("invalid password")
	}

	// convert to json
	infoUserJson, err := json.Marshal(user)
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("convert to json failed: %v", err)
	}
	extTime, _ := time.ParseDuration(global.Config.JWT.JWT_EXPIRATION)

	err = global.Rdb.Set(ctx, user.UserID, infoUserJson, extTime).Err()
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}

	token, err := auth.CreateToken(user.UserID)
	if err != nil {
		return response.ErrorGenAuthCode, out, err
	}
	out.Token = token
	out.Role = string(user.Role.UsersRole)
	out.Email = string(user.Email)
	out.FullName = string(user.FullName)

	return 200, out, nil
}

func NewUserLogin(r *database.Queries) iservice.IUserLogin {
	return &UserLogin{sqlc: r}
}
