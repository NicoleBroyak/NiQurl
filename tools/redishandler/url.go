package redishandler

import (
	"crypto/md5"
	"errors"
	"fmt"
	"path"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nicolebroyak/niqurl/tools/urlhandler"
)

var ServerPath string = "localhost:8081"

func PrintExistingShortURL(longURL *urlhandler.NiqURL) error {
	x, _ := Client.ZScan(Ctx, "longurl", 0, longURL.String(), 0).Val()
	i, err := strconv.Atoi(x[1])
	if err != nil {
		return err
	}
	shorturl := Client.ZRange(Ctx, "shorturl", int64(i), int64(i)).Val()
	fmt.Println("URL [" + longURL.String() + "] shortened before to: " + path.Join(ServerPath, shorturl[0]))
	return nil
}

func InsertURLData(longURL *urlhandler.NiqURL, shorturl, user string) {
	fmt.Println("Creating short URL for [" + longURL.String() + "]: " + path.Join(ServerPath, shorturl))
	wt, _ := getSetting("USER_WAIT_TIME")
	uc, _ := getSetting("URL_COUNT")
	Client.Incr(Ctx, "URL_COUNT")
	Client.ZAdd(Ctx, "longurl", &redis.Z{Score: float64(uc), Member: longURL.String()})
	Client.ZAdd(Ctx, "shorturl", &redis.Z{Score: float64(uc), Member: shorturl})
	Client.RPush(Ctx, "createdby", user)
	Client.Set(Ctx, user, true, time.Duration(wt*1000000000))
}

func ShortenURL(longURL *urlhandler.NiqURL) (string, error) {

	shorturl := shortURLGenerate(longURL)
	if !ExistsShortURL(shorturl) {
		return shorturl, nil
	}
	return "", errors.New(
		"can't generate short url with" +
			"specified length, please change" +
			"length using setlen command")
}

func shortURLGenerate(longURL *urlhandler.NiqURL) string {
	urllen, _ := getSetting("SHORT_URL_LEN")
	longurl := longURL.String()
	longurlbyte := []byte(longurl)
	MD5 := md5.Sum(longurlbyte)
	stringMD5 := fmt.Sprintf("%x", MD5)
	return stringMD5[:urllen]
}

func IsExistingShortURL(shorturl string) (exists bool, index int64) {

	result, _ := Client.ZScan(
		Ctx,
		"shorturl",
		0,
		shorturl,
		0,
	).Val()
	fmt.Println(result)

	if len(result) > 0 {
		index, _ := strconv.Atoi(result[1])
		return true, int64(index)
	}
	return false, 0
}

func QueryForLongURL(index int64) string {

	result := Client.ZRange(
		Ctx,
		"longurl",
		index,
		index,
	).Val()
	fmt.Println(result)
	rawurl := result[0]
	fmt.Println(rawurl)

	return rawurl
}
