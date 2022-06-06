package redishandler

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func CheckWaitTime(user string, RDB *redis.Client) bool {
	v, _ := RDB.Get(Ctx, user).Int()
	if v == 1 {
		x, _ := RDB.TTL(Ctx, user).Result()
		fmt.Println(fmt.Sprintf("User %v has to wait %v to shorten url again", user, x.String()))
		return true
	}
	return false
}
