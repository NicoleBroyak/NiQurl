package redishandler

import (
	"fmt"
)

func CheckWaitTime(user string) bool {
	v, _ := RDB.Get(Ctx, user).Int()
	if v == 1 {
		x := RDB.TTL(Ctx, user).String()
		fmt.Printf("User %v has to wait %v to shorten url again", user, x)
		return true
	}
	return false
}
