package redis

import (
	"context"
	"ecommerce_go/global"
	constant "ecommerce_go/pkg"
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

func CheckSoftLock(roomID, userID string, checkIn, checkOut time.Time) bool {
	ctx := context.Background()
	for d := checkIn; !d.After(checkOut); d = d.AddDate(0, 0, 1) {
		lockID, err := global.Rdb.Get(ctx, fmt.Sprintf("%s:%s-%s", constant.PreSoftLock, roomID, d.Format("2006-01-02"))).Result()
		if err == redis.Nil {
			continue
		} else if err != nil {
			fmt.Println("Redis error:", err)
			return false
		} else if err == nil {
			fmt.Printf("Room %s is locked by %s\n", roomID, lockID)
			return true
		}
	}

	return false
}

func SoftLock(roomID, userID string, checkIn, checkOut time.Time) {
	ctx := context.Background()
	for d := checkIn; !d.After(checkOut); d = d.AddDate(0, 0, 1) {
		err := global.Rdb.Set(ctx, fmt.Sprintf("%s:%s-%s", constant.PreSoftLock, roomID, d.Format("2006-01-02")), roomID, 10*time.Minute)
		if err != nil {
			fmt.Errorf("error when soft lock room %s: %v", roomID, err)
		}
	}
}
