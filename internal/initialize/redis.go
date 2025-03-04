package initialize

import (
	"context"
	"ecommerce_go/global"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {

	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Username: r.UserName,
		Password: r.Password,
		DB:       r.Db,
		PoolSize: r.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		global.Logger.Error("redis ping error", zap.Error(err))
	}

	global.Rdb = rdb
	fmt.Println("init Redis is running")
	redisExample()
}

func redisExample() {
	err := global.Rdb.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		global.Logger.Error("redis set error: ", zap.Error(err))
		return
	}

	// get value
	value, er := global.Rdb.Get(ctx, "foo").Result()

	if er != nil {
		global.Logger.Error("redis get error: ", zap.Error(err))
		return
	}

	global.Logger.Info("redis get result: ", zap.String("foo", value))
}
