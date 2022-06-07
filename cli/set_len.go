package cli

import (
	"errors"
	"fmt"
	"os"
	r "redishandler"

	"github.com/desertbit/grumble"
)

// cli command "setlen"
func SetLen(App *grumble.App) {
	App.AddCommand(&grumble.Command{
		Name: "setlen",
		Help: "set SHORT_URL_LEN variable, must be > 0",
		Args: func(a *grumble.Args) {
			a.Int("setlen", "setlen")
		},
		Run: SetLenFunc,
	})
}

func SetLenFunc(c *grumble.Context) error {
	RDB := r.RedisStart()
	defer RDB.Close()

	v := c.Args.Int("setlen")
	if v < 1 {
		err := errors.New("SHORT_URL_LEN variable, must be > 0")
		return err
	}
	RDB.Set(r.Ctx, "SHORT_URL_LEN", v, 0)
	fmt.Fprintf(os.Stdout, "SHORT_URL_LEN set to %d\n", v)
	LoadSettings()
	return nil
}

// equivalent of SetLenFunc to run tests
func setLenFunc(len int) error {
	if len < 1 {
		err := errors.New("SHORT_URL_LEN variable, must be > 0")
		return err
	}
	return nil
}
