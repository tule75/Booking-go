package repo

import (
	"ecommerce_go/global"
	"ecommerce_go/internal/database"
	"fmt"
	"time"
)

type IUserAuthRepo interface {
	AddOTP(email string, otp string, time int64) error
}

type UserAuthRepo struct {
	sqlc *database.Queries
}

// AddOTP implements IUserAuthRepo.
func (u *UserAuthRepo) AddOTP(email string, otp string, timeExpiration int64) error {
	var index = fmt.Sprintf("usr:%s:otp", email)
	return global.Rdb.SetEx(ctx, index, otp, time.Duration(timeExpiration)*time.Second).Err()
}

func NewUserAuthRepo() IUserAuthRepo {
	return &UserAuthRepo{
		sqlc: database.New(global.Mdbc),
	}
}
