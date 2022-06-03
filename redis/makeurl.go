package redis

import (
	"strconv"
	"time"
)

func MakeURL(longurl, shorturl, user string) error {
	rdb := RedisStart()
	defer rdb.Close()

	id, err := rdb.Get(Ctx, "url_count").Result()
	if err != nil {
		return err
	}

	a, err := rdb.Get(Ctx, "USER_WAIT_TIME").Result()
	if err != nil {
		return err
	}
	i, err := strconv.Atoi(a)
	if err != nil {
		return err
	}

	rdb.Do(Ctx, "incr", "url_count")
	rdb.Do(Ctx, "ZADD", "longurl", id, longurl)
	rdb.Do(Ctx, "ZADD", "shorturl", id, shorturl)
	rdb.Do(Ctx, "RPUSH", "createdby", user)
	rdb.Set(Ctx, user, true, time.Duration(i*1000000000))
	return nil
}
