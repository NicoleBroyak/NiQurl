package redishandler

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

func PrintShortURL(url string) error {
	x, _ := RDB.ZScan(Ctx, "longurl", 0, url, 0).Val()
	i, err := strconv.Atoi(x[1])
	if err != nil {
		return err
	}
	shorturl := RDB.ZRange(Ctx, "shorturl", int64(i), int64(i)).Val()
	fmt.Println("URL [" + url + "] shortened before to: " + shorturl[0])
	return nil
}

func InsertURL(url, shorturl, user string) {
	fmt.Println("Creating short URL for [" + url + "]: " + shorturl)
	wt, _ := getSetting("USER_WAIT_TIME")
	uc, _ := getSetting("URL_COUNT")
	RDB.Incr(Ctx, "URL_COUNT")
	RDB.ZAdd(Ctx, "longurl", &redis.Z{Score: float64(uc), Member: url})
	RDB.ZAdd(Ctx, "shorturl", &redis.Z{Score: float64(uc), Member: shorturl})
	RDB.RPush(Ctx, "createdby", user)
	RDB.Set(Ctx, user, true, time.Duration(wt*1000000000))
}

func ShortURL(url string) string {

	var shrt string
	n, _ := getSetting("SHORT_URL_LEN")

	for i := 0; true; i++ {
		// If loop can't find available url with specific length
		// increase length and try again
		if i == 100 {
			shortURLIncrLen(n)
			i = 0
			continue
		}
		shrt = shortURLGenerate(n)
		if CheckZSet(shrt, "shorturl") == true {
			continue
		}
		break
	}
	return shrt
}

func shortURLIncrLen(n int) {
	n += 1
	fmt.Printf("SHORT_URL_LEN increased to %v\n", n)
	RDB.Set(Ctx, "SHORT_URL_LEN", n, 0)
}

// returns random
func shortURLGenerate(n int) string {
	chr := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	u := make([]byte, n)
	for i := range u {
		rand.Seed(time.Now().UTC().UnixNano())
		u[i] = chr[rand.Intn(len(chr))]
	}

	return string(u)
}

func TrimURL(url string) string {
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "https://www.")
	url = strings.TrimPrefix(url, "www.")
	url = strings.TrimSuffix(url, ".")
	parts := strings.Split(url, ".")
	parts[0] = strings.ToLower(parts[0])
	parts[0] = "https://" + parts[0]
	return strings.Join(parts, ".")
}

func VerifyURL(url string) bool {
	if strings.HasSuffix(url, ".") {
		return false
	}
	if strings.Contains(url, "..") {
		return false
	}
	if !strings.Contains(url, ".") {
		return false
	}
	return true
}
