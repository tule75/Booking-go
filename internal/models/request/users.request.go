package requestDTO

type UserRole string

const (
	CUSTOMER UserRole = "CUSTOMER"
	HOST     UserRole = "HOST"
)

type RegisterRequestModel struct {
	Email string `json:"email", binding:"required,email"`
	Code  string `json:"code"`
}

type LoginRequestModel struct {
	Email    string `json:"email", binding:"required,email"`
	Password string `json:"password", binding:"required,password"`
}

type VerifyRequest struct {
	VerifyKey  string `json:"verify_key"`
	VerifyCode string `json:"verify_code"`
}

type UserCreateRequestModel struct {
	Email    string   `json:"email", binding:"required,email,min=1"`
	Token    string   `json:"token", binding:"required"`
	Password string   `json:"password", binding:"required,password,min=8"`
	FullName string   `json:"full_name", binding:"required"`
	Role     UserRole `json:"role", binding:"required"`
}
