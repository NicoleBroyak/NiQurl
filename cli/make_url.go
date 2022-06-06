package cli

import (
	r "redishandler"

	"github.com/go-redis/redis/v8"

	"github.com/desertbit/grumble"
)

// cli command "make"
func MakeURL(RDB *redis.Client) {
	App.AddCommand(&grumble.Command{
		Name: "make",
		Help: "shorten url",
		Args: func(a *grumble.Args) {
			a.String("url", "url")
		},
		Run: func(c *grumble.Context) error {
			LoadSettings(RDB)
			url := c.Args.String("url")

			b, _ := r.SearchURL(c.Args.String("url"), "longurl", RDB)
			if b {
				r.ReturnURL(url, RDB)
				return nil
			}

			user := r.AssignRandUser(RDB)
			b = r.CheckWaitTime(user, RDB)

			if b == false {
				shorturl := ShortenURL(c.Args.String("url"), RDB)
				r.InsertURL(url, shorturl, user, RDB, USER_WAIT_TIME, URL_COUNT)
			}
			return nil
		},
	})
}
