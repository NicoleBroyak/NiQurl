package redishandler

import (
	"github.com/go-redis/redis/v8"
)

var RDB = RedisStart()
var Ctx = RDB.Context()

func RedisStart() *redis.Client {
	RDB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return RDB
}
