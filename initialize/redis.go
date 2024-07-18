package initialize

import (
	"blog-backend/global"
	"context"
	"github.com/redis/go-redis/v9"
)

func initRedis() (rdb *redis.Client) {
	config := global.CONFIG.RedisConfig
	if len(config.Addr) == 0 {
		return nil
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil
	}

	return rdb
}
