package responseDTO

type LoginResponse struct {
	Token    string `json:"token"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type VerifyOTPResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
