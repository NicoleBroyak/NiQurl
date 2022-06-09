package cli

import (
	r "redishandler"

	"github.com/desertbit/grumble"
)

func cmdMakeURL(c *grumble.Context) error {
	return cmdMakeURLGeneric(c.Args.String("url"))
}

func cmdMakeURLGeneric(url string) error {
	_, err := r.LoadSettings()
	if err != nil {
		return err
	}

	if r.SearchURL(url, "longurl") == true {
		r.ReturnURL(url)
		return nil
	}

	user, err := r.AssignRandUser()
	if err != nil {
		return err
	}

	if r.CheckWaitTime(user) == false {
		shorturl, err := r.ShortURL(url)
		if err != nil {
			return err
		}

		err2 := r.InsertURL(url, shorturl, user)
		if err2 != nil {
			return err2
		}
	}
	return nil
}
