package redishandler

import (
	"fmt"
	"strconv"
)

func ReturnURL(url string) error {
	s, _, _ := RDB.ZScan(Ctx, "longurl", 0, url, 0).Result()
	x := s[1]
	i, err := strconv.Atoi(x)
	if err != nil {
		return err
	}
	j := int64(i)
	sURL := RDB.ZRange(Ctx, "shorturl", j, j).Val()
	fmt.Println("URL [" + url + "] shortened before to: " + sURL[0])
	return nil
}
