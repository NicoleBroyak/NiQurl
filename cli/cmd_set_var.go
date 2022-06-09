package cli

import (
	"fmt"
	"log"
	r "redishandler"

	"github.com/desertbit/grumble"
)

// generic function to implement db variables
// Example: cmdSetVar("setlen", "SHORT_URL_LEN", 1, 20, c)
func cmdSetVar(cmd, v string, min, max int, c *grumble.Context) error {
	if c.Args.Int(cmd) < min || c.Args.Int(cmd) > max {
		err := fmt.Errorf("%q variable must be between %d and %d", v, min, max)
		return err
	}
	r.RDB.Set(r.Ctx, v, c.Args.Int(cmd), 0)
	log.Printf("%q set to %d\n", v, c.Args.Int(cmd))
	return nil
}

func cmdSetLen(c *grumble.Context) error {
	return cmdSetVar("setlen", "SHORT_URL_LEN", 1, 20, c)
}

func cmdSetTime(c *grumble.Context) error {
	return cmdSetVar("settime", "USER_WAIT_TIME", 1, 1<<20, c)
}
