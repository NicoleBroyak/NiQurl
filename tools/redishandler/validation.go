package redishandler

import "log"

func isValidSetting(setting string) bool {
	query, err := client.Get(context, setting).Result()
	if err != nil {
		log.Printf("Value %v not found\n", setting)
		return false
	}

	// Sometimes redis client can return
	// empty string without signaling error

	if len(query) == 0 {
		log.Printf("Value %v not found\n", setting)
		return false
	}

	return true
}

// urlversion "shorturl" or "longurl"
//
// Generic function
func existsURL(urlversion, url string) bool {

	// query returns slice of strings and if user doesn't exist
	// query returns slice with lenght == 0

	query, _ := client.ZScan(context, urlversion, 0, url, 0).Val()
	return len(query) != 0
}

func ExistsShortURL(shorturl string) bool {
	return existsURL("shorturl", shorturl)
}

func ExistsLongURL(longurl string) bool {
	return existsURL("longurl", longurl)
}

func existsUser(username string) bool {

	// query returns slice of strings and if user doesn't exist
	// query returns slice with lenght == 0

	query, _ := client.ZScan(context, "username", 0, username, 0).Val()
	return len(query) != 0
}

//
func IsUserOnWaitTime(user string) bool {
	isUserLimited, _ := client.Get(context, user).Bool()
	return isUserLimited
}
