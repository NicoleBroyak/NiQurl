package redishandler

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func InsertURL(url, shorturl, user string) error {

	waittime, err := loadSetting("USER_WAIT_TIME")
	if err != nil {
		return err
	}
	url_count, err := loadSetting("URL_COUNT")
	if err != nil {
		return err
	}

	fmt.Println("Creating short URL for [" + url + "]: " + shorturl)
	RDB.Incr(Ctx, "URL_COUNT")
	RDB.ZAdd(Ctx, "longurl", &redis.Z{
		Score:  float64(url_count),
		Member: url,
	})
	RDB.ZAdd(Ctx, "shorturl", &redis.Z{
		Score:  float64(url_count),
		Member: shorturl,
	})
	RDB.RPush(Ctx, "createdby", user)
	RDB.Set(Ctx, user, true, time.Duration(waittime*1000000000))

	return nil
}
