package redis

import (
	"context"
	"ecommerce_go/global"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func GetSearchResultsFromRedis(ctx context.Context, pre string, query string) (string, error) {
	cacheKey := fmt.Sprintf("%s:%s", pre, query)

	val, err := global.Rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("Không có dữ liệu cache cho query: %s", query)
	} else if err != nil {
		return "", fmt.Errorf("Lỗi khi truy vấn Redis: %v", err)
	}

	return val, nil
}

func CacheStore(ctx context.Context, pre string, query string, value interface{}, t time.Duration) {
	cacheKey := fmt.Sprintf("%s:%s", pre, query)
	now := time.Now()
	year, week := now.ISOWeek()
	global.Rdb.ZIncrBy(ctx, fmt.Sprintf("%s-%d-%02d", pre, year, week), 1, query)

	rank, _ := global.Rdb.ZScore(ctx, fmt.Sprintf("%s-%d-%02d", pre, year, week), query).Result()

	if rank > 10 {
		jsonData, _ := json.Marshal(value)
		global.Rdb.Set(ctx, cacheKey, jsonData, t)
	}
}

func DeleteCache(ctx context.Context, pre string, query string) error {
	cacheKey := fmt.Sprintf("%s:%s", pre, query)
	_, err := global.Rdb.Del(ctx, cacheKey).Result()
	return err
}
