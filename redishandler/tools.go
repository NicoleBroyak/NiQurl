package redishandler

import "log"

// return true if val exists in set
func CheckZSet(val, set string) bool {
	l, _, _ := RDB.ZScan(Ctx, set, 0, val, 0).Result()
	if len(l) == 0 {
		return false
	}
	return true
}

// check if user has to wait to shorten url again
func CheckWaitTime(user string) bool {
	v, _ := RDB.Get(Ctx, user).Int()
	if v == 1 {
		x := RDB.TTL(Ctx, user).Val()
		log.Printf("User %v has to wait %v to shorten url again", user, x)
		return true
	}
	return false
}
