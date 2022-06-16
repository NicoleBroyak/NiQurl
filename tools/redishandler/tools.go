package redishandler

import "log"

// return true if val exists in set
func ExistsValInZSET(val, set string) bool {
	query, _, _ := Client.ZScan(Ctx, set, 0, val, 0).Result()
	return len(query) != 0
}

func ExistsShortURL(shorturl string) bool {
	return ExistsValInZSET(shorturl, "shorturl")
}

func ExistsLongURL(longurl string) bool {
	return ExistsValInZSET(longurl, "longurl")
}

func PrintUserWaitTime(user string) {
	isuserlimited, _ := Client.Get(Ctx, user).Bool()
	if isuserlimited {
		waittime := Client.TTL(Ctx, user).Val()
		log.Printf("User %v has to wait %v to shorten url again", user, waittime)
	}
}

func IsUserLimited(user string) bool {
	IsUserLimited, _ := Client.Get(Ctx, user).Bool()
	return IsUserLimited
}

func ChangeSetting(setting string, value int) {
	Client.Set(Ctx, setting, value, 0)
	log.Printf("%v set to %v\n", setting, value)
}
