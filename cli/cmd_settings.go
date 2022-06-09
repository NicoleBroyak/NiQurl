package cli

import (
	"errors"
	"fmt"
	"os"
	r "redishandler"

	"github.com/desertbit/grumble"
)

func cmdSettings(c *grumble.Context) error {
	return cmdSettingsGeneric()
}

func cmdSettingsGeneric() error {
	fmt.Println("Current settings")

	b, _ := r.RDB.Get(r.Ctx, "SHORT_URL_LEN").Result()
	if len(b) == 0 {
		return errors.New("SHORT_URL_LEN not found")
	}
	urllen, _ := r.RDB.Get(r.Ctx, "SHORT_URL_LEN").Int()

	b, _ = r.RDB.Get(r.Ctx, "USER_WAIT_TIME").Result()
	if len(b) == 0 {
		return errors.New("USER_WAIT_TIME not found")
	}
	waittime, _ := r.RDB.Get(r.Ctx, "USER_WAIT_TIME").Int()

	fmt.Fprintf(os.Stdout, "short url length: %v characters\n", urllen)
	fmt.Fprintf(os.Stdout, "user wait time: %v s \n", waittime)
	return nil
}
