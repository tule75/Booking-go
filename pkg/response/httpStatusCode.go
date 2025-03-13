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
	CannotCreateRoomCode         = 40003
	CannotCreateReviewCode       = 40004
	NotFoundResponseCode         = 50000
	CannotGetPropertyByOwnerCode = 50001
	CannotGetPropertyCode        = 50002
	CannotGetRoomByIDCode        = 50003
	CannotListRoomByPropertyCode = 50004
	CannotDeletePropertyCode     = 60001
	CannotDeleteRoomCode         = 60002
	CannotDeleteReviewCode       = 60003
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
	UnauthorizedResponseCode:  "Unauthorized",
	CannotCreatePropertyCode:  "Cannot Create Property",
	CannotCreateRoomCode:      "Cannot Create Room",
	CannotCreateReviewCode:    "Cannot Create review request",

	NotFoundResponseCode:         "Not Found Response",
	CannotGetPropertyByOwnerCode: "Cannot get property by owner ID: %s",
	CannotGetPropertyCode:        "Cannot get property",
	CannotGetRoomByIDCode:        "Cannot get room by ID ",
	CannotListRoomByPropertyCode: "Cannot list room by property",

	CannotDeletePropertyCode: "Cannot delete property by PropertyID",
	CannotDeleteRoomCode:     "Cannot delete room by ID",
	CannotDeleteReviewCode:   "Cannot delete review by ID",
}
