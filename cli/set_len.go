package cli

import (
	"errors"
	"fmt"
	"os"
	"redishandler"

	"github.com/desertbit/grumble"
	"github.com/go-redis/redis/v8"
)

// cli command "setlen"
func SetLen(RDB *redis.Client) {
	App.AddCommand(&grumble.Command{
		Name: "setlen",
		Help: "set SHORT_URL_LEN variable, must be > 0",
		Args: func(a *grumble.Args) {
			a.Int("setlen", "setlen")
		},
		Run: func(c *grumble.Context) error {
			v := c.Args.Int("setlen")
			if v < 1 {
				err := errors.New("SHORT_URL_LEN variable, must be > 0")
				return err
			}
			RDB.Set(redishandler.Ctx, "SHORT_URL_LEN", v, 0)
			fmt.Fprintf(os.Stdout, "SHORT_URL_LEN set to %d\n", v)
			LoadSettings(RDB)
			return nil
		},
	})
}
