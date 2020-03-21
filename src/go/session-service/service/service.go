package service

import (
	redis "github.com/go-redis/redis/v7"
	"github.com/ryssapp/backend/src/go/user-service/config"
)

func Start() {
	c := config.Load()
	db := initRedis(c.RedisAddress, c.RedisPassword, c.RedisDatabase)
}

func initRedis(addr string, pw string, db int) *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pw,
		DB:       db,
	})
	_, err = c.Pong().Result()
	if err != nil {
		zap.L().Fatal("Failed to ping redis server.", zap.Error(err))
	}
	defer c.Close()
	return c
}
