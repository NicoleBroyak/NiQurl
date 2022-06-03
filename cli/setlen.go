package cli

import (
	"fmt"
	"os"
	"redis"

	"github.com/desertbit/grumble"
)

func SetLen() {
	App.AddCommand(&grumble.Command{
		Name: "setlen",
		Help: "set SHORT_URL_LEN variable, must be > 0",
		Args: func(a *grumble.Args) {
			a.Int("setlen", "setlen")
		},
		Run: func(c *grumble.Context) error {
			v := c.Args.Int("setlen")
			if v < 1 {
				fmt.Println("SHORT_URL_LEN variable, must be > 0")
				return nil
			}
			rdb := redis.RedisStart()
			defer rdb.Close()
			rdb.Set(redis.Ctx, "SHORT_URL_LEN", v, 0)
			fmt.Fprintf(os.Stdout, "SHORT_URL_LEN set to %d\n", v)
			Settings()
			return nil
		},
	})
}
