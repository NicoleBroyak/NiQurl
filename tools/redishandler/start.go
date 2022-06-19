package redishandler

import (
	"github.com/go-redis/redis/v8"
)

func Start(hostport string) *redis.Client {
	Client := redis.NewClient(&redis.Options{
		Addr:     hostport,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return Client
}
