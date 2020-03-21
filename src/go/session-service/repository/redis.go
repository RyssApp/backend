package repository

import (
	"time"
	"github.com/ryssapp/backend/src/go/session-service/session"
	redis "github.com/go-redis/redis/v7"
)

type redisRepository struct {
	client redis.Client
	expiration time.Duration
}

func New(client *redis.Client, expiration time.Duration) *session.Repository {
	&redisRepository{
		client: client,
		expiration: expiration,
	}
}

func (r *redisRepository) SetToken (id string, token string) error {
	return r.client.Set(id, token, r.expiration).Err()
}

func (r *redisRepository) GetToken(id string) (string, error) {
	return r.client.Get(id).Result()
}
