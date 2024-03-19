package cache

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type RedisClient struct {
	*redis.Client
}

var Redis *RedisClient = nil

func GetRedis() *RedisClient {
	if Redis == nil {
		Redis = &RedisClient{NewRedis()}
	}
	return Redis
}

func NewRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func (rc *RedisClient) Set(key string, value interface{}) error {
	return rc.Client.Set(context.Background(), key, value, time.Hour*24).Err()
}

func (rc *RedisClient) Get(key string) (string, error) {
	return rc.Client.Get(context.Background(), key).Result()
}
