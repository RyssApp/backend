package repository

import (
	redis "github.com/go-redis/redis/v7"
	"github.com/ryssapp/backend/src/go/session-service/session"
	"time"
)

type redisRepository struct {
	client     *redis.Client
	expiration time.Duration
}

func New(client *redis.Client, expiration time.Duration) session.Repository {
	return &redisRepository{
		client:     client,
		expiration: expiration,
	}
}

func (r *redisRepository) SetToken(id string, token string) error {
	return r.client.Set(id, token, r.expiration).Err()
}

func (r *redisRepository) GetToken(id string) (string, error) {
	return r.client.Get(id).Result()
}

func (r *redisRepository) DelToken(id string) error {
	return r.client.Del(id).Err()
}
