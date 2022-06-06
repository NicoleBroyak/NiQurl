package redishandler

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func RedisStart() *redis.Client {
	RDB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return RDB
}
