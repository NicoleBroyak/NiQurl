package niqurlconfigs

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
var SettingsMap = LoadEnvVarsIntoMap(SettingsSlice...)

func LoadEnvVarsIntoMap(settings ...string) map[string]interface{} {
	settingsMap := map[string]interface{}{}
	for setting := range settings {
		setting := fmt.Sprintf("%v", setting)
		envString := os.Getenv(setting)
		settingAsInt, err := strconv.Atoi(envString)
		if err != nil {
			settingsMap[setting] = envString
		} else {
			settingsMap[setting] = settingAsInt
		}
	}
	log.Println(settingsMap)
	return settingsMap
}

func CreateAPISourceFromDefault(num int) string {
	apiSource, _ := SettingsMap["DEFAULT_API_SOURCE"].(string)
	APISource := strings.Replace(apiSource, "[value]", fmt.Sprintf("%v", num), 1)
	return APISource
}
