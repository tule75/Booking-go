package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	// message == "" set msg[code]
	if message == "" {
		message = msg[code]
	}
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func BadResponse(c *gin.Context, code int) {
	c.JSON(http.StatusNotFound, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

func UnauthorizedResponse(c *gin.Context, code int) {
	c.JSON(http.StatusUnauthorized, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
	c.Abort()
}
