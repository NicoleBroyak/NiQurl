package cli

import (
	"errors"
	"fmt"
	"os"
	"redishandler"

	"github.com/desertbit/grumble"
	"github.com/go-redis/redis/v8"
)

// cli command "settime"
func SetTime(RDB *redis.Client) {
	App.AddCommand(&grumble.Command{
		Name: "settime",
		Help: "set USER_WAIT_TIME variable, value in ms, must be > 0",
		Args: func(a *grumble.Args) {
			a.Int("settime", "settime")
		},
		Run: func(c *grumble.Context) error {
			v := c.Args.Int("settime")
			if v < 1 {
				err := errors.New("USER_WAIT_TIME variable must be > 0")
				return err
			}
			RDB.Set(redishandler.Ctx, "USER_WAIT_TIME", v, 0)
			fmt.Fprintf(os.Stdout, "USER_WAIT_TIME set to %d\n", v)
			LoadSettings(RDB)
			return nil
		},
	})
}
