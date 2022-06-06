package redishandler

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func InsertURL(url, shorturl, user string, RDB *redis.Client, waittime, url_count int) error {

	fmt.Println("Creating short URL for [" + url + "]: " + shorturl)
	RDB.Do(Ctx, "incr", "URL_COUNT")
	RDB.Do(Ctx, "ZADD", "longurl", url_count, url)
	RDB.Do(Ctx, "ZADD", "shorturl", url_count, shorturl)
	RDB.Do(Ctx, "RPUSH", "createdby", user)
	RDB.Set(Ctx, user, true, time.Duration(waittime*1000000000))

	return nil
}
