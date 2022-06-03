package redis

import (
	"fmt"
	"strconv"
	"strings"
)

func ReturnURL(longurl string) (string, error) {
	rdb := RedisStart()
	defer rdb.Close()

	query, err := rdb.Do(Ctx, "zscan", "longurl", "0", "match", longurl).Result()
	if err != nil {
		return "", err
	}

	s := fmt.Sprintf("%v", query)
	i := strings.Index(s, longurl)
	j := strings.Index(s, "]]")
	crop := s[i+len(longurl)+1 : j]
	i, err = strconv.Atoi(crop)
	if err != nil {
		return "", err
	}

	query, err = rdb.Do(Ctx, "zrange", "shorturl", i, i).Result()
	if err != nil {
		return "", err
	}

	s = fmt.Sprintf("%v", query)
	return s[1 : len(s)-1], nil
}
