package redishandler

import (
	"errors"
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
func getValueFromSortedSet(sortedSet string, index int64) string {
	query := client.ZRange(context, sortedSet, index, index).Val()
	if len(query) == 0 {
		return ""
	}
	return query[0]
}

func GetLongURL(index int64) string {
	return getValueFromSortedSet("longurl", index)
}

func getShortURL(index int64) string {
	return getValueFromSortedSet("shorturl", index)
}

func getValueFromList(listName string, index int64) string {
	query := client.LRange(context, listName, index, index).Val()
	if len(query) == 0 {
		return ""
	}
	return query[0]
}

func GetURLAuthor(index int64) string {
	return getValueFromList("createdby", index)
}

func getUserName(index int64) string {
	return getValueFromSortedSet("username", index)
}

// urlversion "shorturl" or "longurl"
//
// Generic function of getIndexOfLongURL and GetIndexOfShortURL
func getIndexOfValueFromSortedSet(sortedSetName, value string) (int64, error) {
	query, _ := client.ZScan(context, sortedSetName, 0, value, 0).Val()
	if len(query) == 0 {
		return -1, errors.New("value not found")
	}
	index, _ := strconv.Atoi(query[1])
	return int64(index), nil
}

func GetIndexOfShortURL(shorturl string) (int64, error) {
	return getIndexOfValueFromSortedSet("shorturl", shorturl)

}

func getIndexOfLongURL(longurl string) (int64, error) {
	return getIndexOfValueFromSortedSet("longurl", longurl)
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
