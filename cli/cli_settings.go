package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"

	"github.com/desertbit/grumble"
)

// cli command "settings"
func CLISettings(RDB *redis.Client, Ctx context.Context) {
	App.AddCommand(&grumble.Command{
		Name: "settings",
		Help: "show settings",
		Run: func(c *grumble.Context) error {
			fmt.Println("Current settings")
			urllen, _ := RDB.Get(Ctx, "SHORT_URL_LEN").Result()
			waittime, _ := RDB.Get(Ctx, "USER_WAIT_TIME").Result()
			fmt.Fprintf(os.Stdout, "short url length: %q characters\n", urllen)
			fmt.Fprintf(os.Stdout, "user wait time: %q s \n", waittime)
			return nil
		},
	})
}
