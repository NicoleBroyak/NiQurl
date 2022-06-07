package cli

import (
	r "redishandler"

	"github.com/desertbit/grumble"
)

func MakeURL(App *grumble.App) {
	App.AddCommand(&grumble.Command{
		Name: "make",
		Help: "shorten url",
		Args: func(a *grumble.Args) {
			a.String("url", "url")
		},
		Run: MakeURLFunc,
	})
}

func MakeURLFunc(c *grumble.Context) error {
	RDB := r.RedisStart()
	defer RDB.Close()
	LoadSettings()
	url := c.Args.String("url")

	b := r.SearchURL(c.Args.String("url"), "longurl", RDB)
	if b {
		r.ReturnURL(url, RDB)
		return nil
	}

	user, _ := r.AssignRandUser(RDB)
	b = r.CheckWaitTime(user, RDB)

	if b == false {
		shorturl := ShortenURL(c.Args.String("url"), RDB)
		r.InsertURL(url, shorturl, user, RDB, USER_WAIT_TIME, URL_COUNT)
	}
	return nil
}

// Equivalent of MakeURL to run tests
func makeURLFunc(url string) error {
	RDB := r.RedisStart()
	defer RDB.Close()
	err := LoadSettings()
	if err != nil {
		return err
	}

	b := r.SearchURL(url, "longurl", RDB)
	if b {
		r.ReturnURL(url, RDB)
		return nil
	}

	user, _ := r.AssignRandUser(RDB)
	b = r.CheckWaitTime(user, RDB)

	if b == false {
		shorturl := ShortenURL(url, RDB)
		err := r.InsertURL(url, shorturl, user, RDB, USER_WAIT_TIME, URL_COUNT)
		if err != nil {
			return err
		}
	}
	return nil
}
