package redishandler

import (
	"fmt"
	"log"
	"path"

	"github.com/nicolebroyak/niqurl/config/niqurlconfigs"
)

func printExistingShortURL(shortURL string) {
	fmt.Print("URL shortened before to: ")
	serverPath, _ := niqurlconfigs.SettingsMap["SERVER_PATH"].(string)
	fmt.Println(path.Join(serverPath, shortURL))
}

func printInsertingURLMsg(longURLstring, shorturl string) {
	fmt.Print("Creating short URL for")
	fmt.Printf(" [ %v ]: ", longURLstring)
	serverPath, _ := niqurlconfigs.SettingsMap["SERVER_PATH"].(string)
	fmt.Println(path.Join(serverPath, shorturl))
}

func PrintUserWaitTime(user string) {
	waittime := client.TTL(context, user).Val()
	log.Printf("User %v has to wait %v ms to shorten url again", user, waittime)
}

func PrintCurrentCLISettings() {
	fmt.Println("Current settings")
	fmt.Printf("short url length: %v characters\n", GetSetting("SHORT_URL_LEN"))
	fmt.Printf("user wait time: %v ms \n", GetSetting("USER_WAIT_TIME"))
}
