package initialize

import (
	"ecommerce_go/global"
	"ecommerce_go/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.LogSettings)
}
