package query

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetLimitAndOffsetFromQuery(ctx *gin.Context) (int32, int32) {
	limit := int32(func() int {
		if v, err := strconv.Atoi(ctx.Query("limit")); err == nil {
			return v
		}
		return 20
	}())
	offset := int32(func() int {
		if v, err := strconv.Atoi(ctx.Query("offset")); err == nil {
			return v
		}
		return 0
	}())
	return limit, offset
}
