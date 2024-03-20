package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func InitDB(url string) (*redis.Client, error) {

	redisClient := redis.NewClient(&redis.Options{
		Addr: url,
	})

	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		panic(err)
		return nil, err
	}

	return redisClient, nil
}
