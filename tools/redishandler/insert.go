package redishandler

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nicolebroyak/niqurl/tools/randomusers"
	"github.com/nicolebroyak/niqurl/tools/urlhandler"
)

func InsertUsers(UsersStruct *randomusers.UsersStruct) {
	user_index := 0
	for user_index < len(UsersStruct.Results) {
		if !existsUser(UsersStruct.Results[user_index].Login.Username) {
			insertUserData(UsersStruct, user_index)
			incrementUsersCount()
		}
		user_index++
	}
}

func insertUserData(UsersData *randomusers.UsersStruct, user_index int) {
	insertUserName(UsersData.Results[user_index].Login.Username)
	insertEmail(UsersData.Results[user_index].Email)
	insertFirstName(UsersData.Results[user_index].Name.First)
	insertLastName(UsersData.Results[user_index].Name.Last)
	insertRegistrationDate(UsersData.Results[user_index].Registered.Date)
}

// urlversion "shorturl" or "longurl"

// generic function for insertLongURL and insertShortURL
func insertURL(urlversion, url string) {
	urlCount := GetSetting("URL_COUNT")
	index := float64(urlCount)
	client.ZAdd(
		context,
		urlversion,
		&redis.Z{Score: index, Member: url},
	)
}

func insertLongURL(url string) {
	insertURL("longurl", url)
}

func insertShortURL(url string) {
	insertURL("shorturl", url)
}

func insertURLAuthor(user string) {
	client.RPush(context, "createdby", user)
}

func insertIntoList(listname string, value interface{}) {
	client.RPush(context, listname, value)
}

func insertUserName(username string) {
	insertToZSet("USER_COUNT", "username", username)
}

func insertEmail(email string) {
	insertToZSet("USER_COUNT", "email", email)
}

func insertFirstName(firstname string) {
	insertIntoList("firstname", firstname)
}

func insertLastName(lastname string) {
	insertIntoList("lastname", lastname)
}

func insertRegistrationDate(regdate time.Time) {
	insertIntoList("regdate", regdate)
}

func InsertURLData(NiqURL *urlhandler.NiqURL) {
	printInsertingURLMsg(NiqURL.String(), NiqURL.ShortURL)

	insertLongURL(NiqURL.String())
	insertShortURL(NiqURL.ShortURL)
	insertURLAuthor(NiqURL.UserName)

	incrementURLCount()
	insertUserWaitTime(NiqURL.UserName)
}

func insertUserWaitTime(user string) {
	waitTime := GetSetting("USER_WAIT_TIME")
	waitTime *= 1000000 // nanoseconds to miliseconds to satisft client.Set function
	client.Set(context, user, true, time.Duration(int64(waitTime)))
}

func insertToZSet(setting, zSetName string, value interface{}) {
	index := GetSetting(setting)
	client.ZAdd(context, zSetName, &redis.Z{
		Score:  float64(index),
		Member: value,
	})
}
