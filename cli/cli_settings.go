package cli

import (
	"fmt"
	"os"
	"redis"

	"github.com/desertbit/grumble"
)

func CLISettings() {
	App.AddCommand(&grumble.Command{
		Name: "settings",
		Help: "show settings",
		Run: func(c *grumble.Context) error {
			fmt.Println("Current settings")
			rdb := redis.RedisStart()
			defer rdb.Close()
			urllen, _ := rdb.Get(redis.Ctx, "SHORT_URL_LEN").Result()
			waittime, _ := rdb.Get(redis.Ctx, "USER_WAIT_TIME").Result()
			fmt.Fprintf(os.Stdout, "short url length: %q characters\n", urllen)
			fmt.Fprintf(os.Stdout, "user wait time: %q ms \n", waittime)
			c.App.Println("generate-fake-users:", c.Flags.Int("generate-fake-users"))
			return nil
		},
	})
}
