package repo

import (
	"ecommerce_go/global"
	"ecommerce_go/internal/database"
	"fmt"
)

type IUserRepo interface {
	GetUserByEmail(email string) bool
	GetUserInfo() string
	CheckEmailAndPassword(email string, password string) (database.GetUserByEmailRow, error)
}

type UserRepo struct {
	sqlc *database.Queries
}

func NewUserRepo() IUserRepo {
	return &UserRepo{
		sqlc: database.New(global.Mdbc),
	}
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

// Check email and password
func (ur *UserRepo) CheckEmailAndPassword(email string, password string) (database.GetUserByEmailRow, error) {
	user, err := ur.sqlc.GetUserByEmail(ctx, email)

	return user, err
}

func (ur *UserRepo) GetUserInfo() string {
	return "tule"
}
