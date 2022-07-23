package data

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	c "github.com/wannanbigpig/gin-layout/config"
)

var rdb *redis.Client
var ctx = context.Background()

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     c.Config.Redis.Host + ":" + c.Config.Redis.Port,
		Password: c.Config.Redis.Password,
		DB:       c.Config.Redis.Database,
	})
	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		panic("Redis connection failed：" + err.Error())
	}
}

func SetRedis(key string, value string, t int64) bool {
	expire := time.Duration(t) * time.Second
	if err := rdb.Set(ctx, key, value, expire).Err(); err != nil {
		return false
	}
	return true
}

func GetRedis(key string) string {
	result, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return result
}

func DelRedis(key string) bool {
	_, err := rdb.Del(ctx, key).Result()
	if err != nil {
		return false
	}
	return true
}

func ExpireRedis(key string, t int64) bool {
	// 延长过期时间
	expire := time.Duration(t) * time.Second
	if err := rdb.Expire(ctx, key, expire).Err(); err != nil {
		return false
	}
	return true
}
