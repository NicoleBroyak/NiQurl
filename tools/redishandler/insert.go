package redishandler

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nicolebroyak/niqurl/tools/randomusers"
	"github.com/nicolebroyak/niqurl/tools/urlhandler"
)

func insertIntoList(listname string, value interface{}) {
	client.RPush(context, listname, value)
}

func insertIntoSortedSet(zSetName string, index float64, value interface{}) {
	client.ZAdd(context, zSetName, &redis.Z{
		Score:  index,
		Member: value,
	})
}

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

func insertLongURL(longURL string) {
	urlCount := GetSetting("URL_COUNT")
	insertIntoSortedSet("longurl", float64(urlCount), longURL)
}

func insertShortURL(shortURL string) {
	urlCount := GetSetting("URL_COUNT")
	insertIntoSortedSet("shorturl", float64(urlCount), shortURL)
}

func insertURLAuthor(user string) {
	client.RPush(context, "createdby", user)
}

func insertUserName(username string) {
	userCount := GetSetting("USER_COUNT")
	insertIntoSortedSet("username", float64(userCount), username)
}

func insertEmail(email string) {
	userCount := GetSetting("USER_COUNT")
	insertIntoSortedSet("email", float64(userCount), email)
}

func insertFirstName(firstname string) {
	insertIntoList("firstname", firstname)
}

func insertLastName(lastname string) {
	insertIntoList("lastname", lastname)
}

func insertRegistrationDate(regdate time.Time) {
	insertIntoList("regdate", regdate.String())
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
