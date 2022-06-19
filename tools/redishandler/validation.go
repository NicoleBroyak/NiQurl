package redishandler

import "log"

func isValidSetting(setting string) bool {
	_, err := client.Get(context, setting).Result()
	if err != nil {
		log.Printf("Value %v not found\n", setting)
		return false
	}

	return true
}

// Generic function
func existsValueInSortedSet(sortedSetName, url string) bool {

	// query returns slice of strings and if url doesn't exist
	// query returns slice with lenght == 0

	query, _ := client.ZScan(context, sortedSetName, 0, url, 0).Val()
	return len(query) != 0
}

func ExistsShortURL(shorturl string) bool {
	return existsValueInSortedSet("shorturl", shorturl)
}

func ExistsLongURL(longurl string) bool {
	return existsValueInSortedSet("longurl", longurl)
}

func existsUser(username string) bool {
	return existsValueInSortedSet("username", username)
}

func IsUserOnWaitTime(user string) bool {
	isUserLimited, _ := client.Get(context, user).Bool()
	return isUserLimited
}
