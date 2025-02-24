package global

import (
	"ecommerce_go/pkg/logger"
	"ecommerce_go/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
)
