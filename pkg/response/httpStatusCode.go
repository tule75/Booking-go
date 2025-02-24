package response

const (
	SuccessResponseCode       = 20000
	FalseEmailResponseCode    = 20001
	FalsePasswordResponseCode = 20002
	InvalidTokenResponseCode  = 20003
)

var msg = map[int]string{
	SuccessResponseCode:       "Thành công",
	FalseEmailResponseCode:    "False Email",
	FalsePasswordResponseCode: "False Password",
	InvalidTokenResponseCode:  "Invalid Token response",
}
