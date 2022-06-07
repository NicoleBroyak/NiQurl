package cli

import (
	"errors"
	"fmt"
	"os"
	r "redishandler"

	"github.com/desertbit/grumble"
)

func CLISettings(App *grumble.App) {
	App.AddCommand(&grumble.Command{
		Name:    "settings",
		Aliases: []string{"show settings", "setup", "config"},
		Help:    "show settings",
		Run:     ShowSettings,
	})
}

func ShowSettings(c *grumble.Context) error {
	RDB := r.RedisStart()
	defer RDB.Close()

	fmt.Println("Current settings")

	b, _ := RDB.Get(r.Ctx, "SHORT_URL_LEN").Result()
	if len(b) == 0 {
		return errors.New("SHORT_URL_LEN not found")
	}
	urllen, _ := RDB.Get(r.Ctx, "SHORT_URL_LEN").Int()

	b, _ = RDB.Get(r.Ctx, "USER_WAIT_TIME").Result()
	if len(b) == 0 {
		return errors.New("USER_WAIT_TIME not found")
	}
	waittime, _ := RDB.Get(r.Ctx, "USER_WAIT_TIME").Int()

	fmt.Fprintf(os.Stdout, "short url length: %v characters\n", urllen)
	fmt.Fprintf(os.Stdout, "user wait time: %v s \n", waittime)
	return nil
}

// equivalent of ShowSettings just to run tests
func showSettings() error {
	RDB := r.RedisStart()
	defer RDB.Close()

	fmt.Println("Current settings")

	b, _ := RDB.Get(r.Ctx, "SHORT_URL_LEN").Result()
	if len(b) == 0 {
		return errors.New("SHORT_URL_LEN not found")
	}
	urllen, _ := RDB.Get(r.Ctx, "SHORT_URL_LEN").Int()

	b, _ = RDB.Get(r.Ctx, "USER_WAIT_TIME").Result()
	if len(b) == 0 {
		return errors.New("USER_WAIT_TIME not found")
	}
	waittime, _ := RDB.Get(r.Ctx, "USER_WAIT_TIME").Int()

	fmt.Fprintf(os.Stdout, "short url length: %v characters\n", urllen)
	fmt.Fprintf(os.Stdout, "user wait time: %v s \n", waittime)
	return nil
}
