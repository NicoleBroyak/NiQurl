package cli

import (
	"math/rand"
	"redis"
	"time"
)

func ShortenURL(longurl string) (string, string, error) {
	b := make([]byte, SHORT_URL_LEN)
	for i := range b {
		rand.Seed(time.Now().UTC().UnixNano())
		b[i] = ALLOWED_URL_CHARS[rand.Intn(len(ALLOWED_URL_CHARS))]
	}
	rdb := redis.RedisStart()
	defer rdb.Close()

	return longurl, string(b), nil

}
