package cli

import (
	"fmt"
	"os"
	"redis"

	"github.com/desertbit/grumble"
)

func SetTime() {
	App.AddCommand(&grumble.Command{
		Name: "settime",
		Help: "set USER_WAIT_TIME variable, value in ms, must be > 0",
		Args: func(a *grumble.Args) {
			a.Int("settime", "settime")
		},
		Run: func(c *grumble.Context) error {
			v := c.Args.Int("settime")
			if v < 1 {
				fmt.Println("USER_WAIT_TIME variable, must be > 0")
				return nil
			}
			rdb := redis.RedisStart()
			defer rdb.Close()
			rdb.Set(redis.Ctx, "USER_WAIT_TIME", v, 0)
			fmt.Fprintf(os.Stdout, "USER_WAIT_TIME set to %d\n", v)
			Settings()
			return nil
		},
	})
}
