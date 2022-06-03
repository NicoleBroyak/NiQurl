package cli

import (
	"redis"
	"strconv"
)

var SHORT_URL_LEN int8 = 4
var USER_WAIT_TIME int = 1000

const ALLOWED_URL_CHARS string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Settings() {
	rdb := redis.RedisStart()
	defer rdb.Close()
	l, _ := rdb.Get(redis.Ctx, "SHORT_URL_LEN").Result()
	t, _ := rdb.Get(redis.Ctx, "USER_WAIT_TIME").Result()
	ls, _ := strconv.Atoi(l)
	var lp *int8 = &SHORT_URL_LEN
	var lt *int = &USER_WAIT_TIME
	*lp = int8(ls)
	*lt, _ = strconv.Atoi(t)
}
