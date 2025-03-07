package repo

import (
	"ecommerce_go/global"
	"ecommerce_go/internal/database"
	"fmt"
)

type IUserRepo interface {
	GetUserByEmail(email string) bool
	GetUserInfo() string
}

type UserRepo struct {
	sqlc *database.Queries
}

// Register implements IUserRepo.
func (ur *UserRepo) GetUserByEmail(email string) bool {
	us, err := ur.sqlc.GetUserByEmail(ctx, email)
	if err != nil {
		fmt.Printf("GetUserByEmail error: %v\n", err)
		return false
	}

	return us.Email == email
}

func NewUserRepo() IUserRepo {
	return &UserRepo{
		sqlc: database.New(global.Mdbc),
	}
}

func (ur *UserRepo) GetUserInfo() string {
	return "tule"
}
