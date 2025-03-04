package repo

type IUserRepo interface {
	GetUserByEmail(email string) bool
	GetUserInfo() string
}

type UserRepo struct{}

// Register implements IUserRepo.
func (ur *UserRepo) GetUserByEmail(email string) bool {
	return false
}

func NewUserRepo() IUserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetUserInfo() string {
	return "tule"
}
