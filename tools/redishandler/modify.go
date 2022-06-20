package redishandler

import (
	"log"

	"github.com/nicolebroyak/niqurl/config/niqurlconfigs"
	"github.com/nicolebroyak/niqurl/tools/randomusers"
)

// used regularly in CLI to constantly provide valid settings
func SetInvalidSettingsToDefaults() {
	for setting, defaultValue := range niqurlconfigs.SettingsMap {
		if !isValidSetting(setting) {
			ChangeSetting(setting, defaultValue)
			if setting == "USER_COUNT" {
				UsersStruct := randomusers.GenerateFakeUsers(niqurlconfigs.CreateAPISourceFromDefault(5), 5)
				InsertUsers(UsersStruct)
			}
		}
	}
}

func ChangeSetting(setting string, value string) {
	client.Set(context, setting, value, 0)
	log.Printf("%v set to %v\n", setting, value)
}

func incrementUsersCount() {
	client.Incr(context, "USER_COUNT")
}

func incrementURLCount() {
	client.Incr(context, "URL_COUNT")
}
