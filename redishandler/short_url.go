package redishandler

import (
	"fmt"
	"math/rand"
	"time"
)

func ShortURL(longurl string) (string, error) {

	var shorturl string
	const allowedchars string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	urllen, err := loadSetting("SHORT_URL_LEN")
	if err != nil {
		return "", err
	}

	for b, i := true, 0; b; i++ {
		// If loop can't find available url with specific length, increase length and try again
		if i == 100 {
			shortURLIncrLen(urllen)
			continue
		}
		shorturl = shortURLGenerate(urllen)
		if SearchURL(shorturl, "shorturl") == true {
			continue
		}
		break
	}
	return shorturl, nil
}

func shortURLIncrLen(urllen int) {
	urllen += 1
	fmt.Printf("SHORT_URL_LEN increased to %v\n", urllen)
	RDB.Set(Ctx, "SHORT_URL_LEN", urllen, 0)
}

func shortURLGenerate(urllen int) string {
	const allowedchars string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	u := make([]byte, urllen)
	for i := range u {
		rand.Seed(time.Now().UTC().UnixNano())
		u[i] = allowedchars[rand.Intn(len(allowedchars))]
	}

	return string(u)
}
