package redishandler

import (
	"github.com/go-redis/redis/v8"
)

var RDB = Start()
var Ctx = RDB.Context()

func Start() *redis.Client {
	RDB := redis.NewClient(&redis.Options{
		Addr:     "niqurl-redis-1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return RDB
}
