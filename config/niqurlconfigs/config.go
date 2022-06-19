package niqurlconfigs

import (
	"fmt"
	"strings"
)

var DefaultAPISource = "https://randomuser.me/api/?results=[value]&inc=login,name,email,registered"
var ServerPort = "8081"
var ServerPath = "localhost:" + ServerPort
var SettingsMap = map[string]int{
	"SHORT_URL_LEN":  4,
	"USER_WAIT_TIME": 30000,
	"URL_COUNT":      0,
	"USER_COUNT":     0,
}

func CreateAPISourceFromDefault(num int) string {
	APISource := strings.Replace(DefaultAPISource, "[value]", fmt.Sprintf("%v", num), 1)
	return APISource
}

var RedisHost = "localhost:6379"
