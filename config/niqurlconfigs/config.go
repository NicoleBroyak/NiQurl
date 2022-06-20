package niqurlconfigs

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var SettingsSlice = []string{
	"SHORT_URL_LEN",
	"URL_COUNT",
	"USER_COUNT",
	"USER_WAIT_TIME",
	"DEFAULT_API_SOURCE",
	"REDIS_PORT",
	"SERVER_PORT",
	"SERVER_PATH",
}
var SettingsMap = LoadEnvVarsIntoMap(SettingsSlice)

func LoadEnvVarsIntoMap(settings []string) map[string]string {
	settingsMap := map[string]string{}
	for setting := range settings {
		settingsMap[fmt.Sprintf("%v", setting)] = os.Getenv(fmt.Sprintf("%v", setting))
	}
	log.Println(settingsMap)
	return settingsMap
}

func CreateAPISourceFromDefault(num int) string {
	apiSource := SettingsMap["DEFAULT_API_SOURCE"]
	apiSource = strings.Replace(apiSource, "[value]", fmt.Sprintf("%v", num), 1)
	return apiSource
}
