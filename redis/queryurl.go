package redis

import (
	"fmt"
)

func QueryURL(url string) (bool, error) {
	rdb := RedisStart()
	defer rdb.Close()
	result, _ := rdb.Do(Ctx, "zscan", "longurl", "0", "match", url).Result()
	s := fmt.Sprintf("%v", result)
	if s == "[0 []]" {
		return false, nil
	}
	return true, nil
}
