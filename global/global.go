package global

import (
	"database/sql"
	"ecommerce_go/pkg/logger"
	"ecommerce_go/pkg/setting"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

var (
	Config        setting.Config
	Logger        *logger.LoggerZap
	Mdb           *gorm.DB
	Rdb           *redis.Client
	Mdbc          *sql.DB
	KafkaProducer *kafka.Writer
)
