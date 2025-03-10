package initialize

import (
	"ecommerce_go/global"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() *gin.Engine {
	LoadConfig()
	fmt.Println("mysql ", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Config ok roif: ", zap.String("Ok", "success"))
	InitMySQL()
	InitMySQLC()
	InitRedis()
	InitKafka()
	InitServices()
	r := InitRouter()
	return r
}
