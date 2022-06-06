package redishandler

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func ReturnURL(url string, rdb *redis.Client) error {
	s, _, _ := rdb.ZScan(Ctx, "longurl", 0, url, 0).Result()
	x := s[1]
	i, err := strconv.Atoi(x)
	if err != nil {
		return err
	}
	j := int64(i)
	sURL := rdb.ZRange(Ctx, "shorturl", j, j).Val()
	fmt.Println("URL [" + url + "] shortened before to: " + sURL[0])
	return nil
}
