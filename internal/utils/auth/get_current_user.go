package auth

import (
	"ecommerce_go/global"
	"ecommerce_go/internal/database"
	"ecommerce_go/pkg/response"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func GetCurrentUserId(c *gin.Context) string {
	id, exists := c.Get("currentUser")
	if !exists {
		response.UnauthorizedResponse(c, response.UnauthorizedResponseCode)
		return ""
	}

	return id.(string)
}
func GetCurrentUser(c *gin.Context) (database.GetUserByEmailRow, error) {
	var user database.GetUserByEmailRow
	userID := GetCurrentUserId(c)

	infoUserJson, err := global.Rdb.Get(c, userID).Result()

	if err == redis.Nil {
		fmt.Println("User not found in Redis")
		response.UnauthorizedResponse(c, response.UnauthorizedResponseCode)
		return database.GetUserByEmailRow{}, err
	} else if err != nil {
		fmt.Println("Error getting user from Redis:", err)
		response.UnauthorizedResponse(c, response.UnauthorizedResponseCode)
		return database.GetUserByEmailRow{}, err
	} else {
		err = json.Unmarshal([]byte(infoUserJson), &user)
		if err != nil {
			fmt.Println("Error unmarshalling user:", err)
			response.UnauthorizedResponse(c, response.UnauthorizedResponseCode)
			return database.GetUserByEmailRow{}, err
		}
	}

	return user, nil
}
