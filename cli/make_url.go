package cli

import (
	"fmt"
	"redis"

	"github.com/desertbit/grumble"
)

func MakeURL() {
	App.AddCommand(&grumble.Command{
		Name: "make",
		Help: "shorten url",
		Args: func(a *grumble.Args) {
			a.String("url", "url")
		},
		Run: func(c *grumble.Context) error {
			url := c.Args.String("url")
			b, _ := redis.QueryURL(url)
			if b {
				shorturl, _ := redis.ReturnURL(url)
				fmt.Println("URL [" + url + "] shortened before to: " + shorturl)
				return nil
			}
			user := AssignRandUser()
			longurl, shorturl, _ := ShortenURL(c.Args.String("url"))
			redis.MakeURL(longurl, shorturl, user)
			fmt.Println("Creating short URL for [" + longurl + "]: " + shorturl)
			// redis.WaitTime()
			return nil
		},
	})
}
