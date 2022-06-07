package cli

import (
	"fmt"
	"math/rand"
	"redishandler"
	"time"

	"github.com/go-redis/redis/v8"
)

func ShortenURL(longurl string, RDB *redis.Client) string {

	shorturl := ""
	for b, i := true, 0; b; i++ {
		u := make([]byte, SHORT_URL_LEN)
		for i := range u {
			rand.Seed(time.Now().UTC().UnixNano())
			u[i] = ALLOWED_URL_CHARS[rand.Intn(len(ALLOWED_URL_CHARS))]
		}

		shorturl = string(u)
		q := redishandler.SearchURL(shorturl, "shorturl", RDB)
		if q == true {
			continue
		}

		// If loop can't find available url with specific length, increase length and try again
		if i == 100 {
			SHORT_URL_LEN += 1
			fmt.Printf("SHORT_URL_LEN increased to %v\n", SHORT_URL_LEN)
			RDB.Set(redishandler.Ctx, "SHORT_URL_LEN", SHORT_URL_LEN, 0)
			continue
		}
		break
	}
	return shorturl
}
