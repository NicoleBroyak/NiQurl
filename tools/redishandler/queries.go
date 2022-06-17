package redishandler

import (
	"math/rand"
	"strconv"
	"time"
)

// function assumes that setting was validated before invoking
func GetSetting(setting string) int {
	settingValue, _ := client.Get(context, setting).Int()
	return settingValue
}

// urlversion "shorturl" or "longurl"
//
// Generic function of GetLongURL and GetShortURL
func getURL(urlversion string, index int64) string {
	query := client.ZRange(context, urlversion, index, index).Val()
	return query[0]
}

func GetLongURL(index int64) string {
	return getURL("longurl", index)
}

func getShortURL(index int64) string {
	return getURL("shorturl", index)
}

func GetURLAuthor(index int64) string {
	query := client.ZRange(context, "createdby", index, index).Val()
	return query[0]
}

func getUserName(index int64) string {
	query := client.ZRange(context, "username", index, index).Val()
	return query[0]
}

// urlversion "shorturl" or "longurl"
//
// Generic function of getIndexOfLongURL and GetIndexOfShortURL
func getIndexOfURL(urlversion, url string) (int64, error) {
	query, _ := client.ZScan(context, urlversion, 0, url, 0).Val()
	index, err := strconv.Atoi(query[1])
	if err != nil {
		return 0, err
	}
	return int64(index), nil
}

func GetIndexOfShortURL(shorturl string) (int64, error) {
	return getIndexOfURL("shorturl", shorturl)

}

func getIndexOfLongURL(longurl string) (int64, error) {
	return getIndexOfURL("longurl", longurl)
}

func GetRandomUser() string {
	rand.Seed(time.Now().UTC().UnixNano())
	userCount := GetSetting("USER_COUNT")
	index := rand.Intn(userCount)
	return getUserName(int64(index))
}

func ProcessExistingURL(longURL string) {
	index, _ := getIndexOfLongURL(longURL)
	shortURL := getShortURL(index)
	printExistingShortURL(shortURL)
}
