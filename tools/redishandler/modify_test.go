package redishandler

import (
	"testing"

	"github.com/nicolebroyak/niqurl/config/niqurlconfigs"
)

func TestChangeSetting(t *testing.T) {
	urlLenOld, err := client.Get(context, "SHORT_URL_LEN").Int()
	if err != nil {
		t.Fatalf("Error from redis client")
	}
	urlLenNew := urlLenOld + 2
	ChangeSetting("SHORT_URL_LEN", urlLenNew)
	urlLenNew, err = client.Get(context, "SHORT_URL_LEN").Int()
	client.Set(context, "SHORT_URL_LEN", urlLenOld, 0)
	if err != nil {
		t.Fatalf("Error from redis client")
	}
	if urlLenNew != urlLenOld+2 {
		t.Fatalf(`Error: %v want match for %v`, urlLenNew, urlLenOld+2)
	}
}

func TestSetInvalidSettingsToDefaults(t *testing.T) {
	urlLenBackup, _ := client.Get(context, "SHORT_URL_LEN").Int()
	waitTimeBackup, _ := client.Get(context, "USER_WAIT_TIME").Int()
	urlCountBackup, _ := client.Get(context, "URL_COUNT").Int()
	userCountBackup, _ := client.Get(context, "USER_COUNT").Int()
	client.Del(
		context,
		"SHORT_URL_LEN",
		"USER_WAIT_TIME",
		"URL_COUNT",
		"USER_COUNT",
	)
	restorePreviousSettings := func() {
		client.Set(context, "SHORT_URL_LEN", urlLenBackup, 0)
		client.Set(context, "USER_WAIT_TIME", waitTimeBackup, 0)
		client.Set(context, "URL_COUNT", urlCountBackup, 0)
		client.Set(context, "USER_COUNT", userCountBackup+5, 0)
	}
	SetInvalidSettingsToDefaults()
	PrintCurrentCLISettings()
	for setting, defValue := range niqurlconfigs.SettingsMap {
		_, ok := defValue.(int)
		if ok {
			afterSetup, _ := client.Get(context, setting).Int()
			defValueInt, _ := defValue.(int)
			if setting != "USER_COUNT" {
				if afterSetup != defValueInt {
					t.Fatalf(`Error in setting %q: %v want match for %v`, setting, afterSetup, defValueInt)
					restorePreviousSettings()
				}
			} else {
				if afterSetup != defValueInt+5 {
					t.Fatalf(`Error in setting %q: %v want match for %v`, setting, afterSetup, defValueInt+5)
					restorePreviousSettings()
				}
			}
		}
		restorePreviousSettings()
	}
}
