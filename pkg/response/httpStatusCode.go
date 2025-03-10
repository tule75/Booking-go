package response

const (
	SuccessResponseCode          = 20000
	FalseEmailResponseCode       = 20001
	FalsePasswordResponseCode    = 20002
	InvalidTokenResponseCode     = 20003
	ErrorBindingParam            = 30001
	UnauthorizedResponseCode     = 30006 //
	ExistEmailResponseCode       = 20004
	ErrorGenAuthCode             = 30002
	ErrCodeAuthFailed            = 30003
	CannotCreateUser             = 40001
	CannotCreatePropertyCode     = 40002
	CannotGetPropertyByOwnerCode = 50001
	CannotGetPropertyCode        = 50002
)

var msg = map[int]string{
	SuccessResponseCode:          "Thành công",
	FalseEmailResponseCode:       "False Email",
	FalsePasswordResponseCode:    "False Password",
	InvalidTokenResponseCode:     "Invalid Token response",
	ErrorBindingParam:            "Error Binding parameter",
	ExistEmailResponseCode:       "Exists Email",
	ErrorGenAuthCode:             "Error Gen JWT Authentication token",
	ErrCodeAuthFailed:            "Error Auth",
	CannotCreateUser:             "Cannot Create User",
	UnauthorizedResponseCode:     "Unauthorized",
	CannotCreatePropertyCode:     "Cannot Create Property",
	CannotGetPropertyByOwnerCode: "Cannot get property by owner ID: %s",
	CannotGetPropertyCode:        "Cannot get property",
}
