package response

const (
	SuccessResponseCode       = 20000
	FalseEmailResponseCode    = 20001
	FalsePasswordResponseCode = 20002
	InvalidTokenResponseCode  = 20003
	ErrorBindingParam         = 30001
	ExistEmailResponseCode    = 20004
	ErrorGenAuthCode          = 30002
	ErrCodeAuthFailed         = 30003
	CannotCreateUser          = 30004
)

var msg = map[int]string{
	SuccessResponseCode:       "Thành công",
	FalseEmailResponseCode:    "False Email",
	FalsePasswordResponseCode: "False Password",
	InvalidTokenResponseCode:  "Invalid Token response",
	ErrorBindingParam:         "Error Binding parameter",
	ExistEmailResponseCode:    "Exists Email",
	ErrorGenAuthCode:          "Error Gen JWT Authentication token",
	ErrCodeAuthFailed:         "Error Auth",
	CannotCreateUser:          "Cannot Create User",
}
