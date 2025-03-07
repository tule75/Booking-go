package initialize

import (
	"ecommerce_go/global"
	"fmt"

	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("mysql ", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Config ok roif: ", zap.String("Ok", "success"))
	InitMySQL()
	InitMySQLC()
	InitRedis()
	InitKafka()
	r := InitRouter()

	r.Run(fmt.Sprintf(":%v", global.Config.Server.Port))
}
