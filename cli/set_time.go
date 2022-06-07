package cli

import (
	"errors"
	"fmt"
	"os"
	r "redishandler"

	"github.com/desertbit/grumble"
)

func SetTime(App *grumble.App) {
	App.AddCommand(&grumble.Command{
		Name: "settime",
		Help: "set USER_WAIT_TIME variable, value in ms, must be > 0",
		Args: func(a *grumble.Args) {
			a.Int("settime", "settime")
		},
		Run: SetTimeFunc,
	})
}

func SetTimeFunc(c *grumble.Context) error {
	RDB := r.RedisStart()
	defer RDB.Close()

	v := c.Args.Int("settime")
	if v < 1 {
		err := errors.New("USER_WAIT_TIME variable must be > 0")
		return err
	}
	RDB.Set(r.Ctx, "USER_WAIT_TIME", v, 0)
	fmt.Fprintf(os.Stdout, "USER_WAIT_TIME set to %d\n", v)
	LoadSettings()
	return nil
}

// equivalent of SetLenFunc to run tests
func setTimeFunc(time int) error {
	if time < 1 {
		err := errors.New("USER_WAIT_TIME variable, must be > 0")
		return err
	}
	return nil
}
