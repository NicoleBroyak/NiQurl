package cli

import (
	"fmt"
	"log"

	"github.com/nicolebroyak/niqurl/tools/redishandler"
	"github.com/nicolebroyak/niqurl/tools/urlhandler"

	"github.com/desertbit/grumble"
)

func commands(app *grumble.App) {
	settime(app)
	setlen(app)
	makeurl(app)
	settings(app)
}

func settime(app *grumble.App) {
	app.AddCommand(&grumble.Command{
		Name:    "settime",
		Aliases: []string{"SETTIME"},
		Help:    "set USER_WAIT_TIME variable, value in ms, must be > 0",
		Args: func(a *grumble.Args) {
			a.Int("settime", "settime")
		},
		Run: cmdSetTime,
	})
}

func setlen(app *grumble.App) {
	app.AddCommand(&grumble.Command{
		Name:    "setlen",
		Aliases: []string{"SETLEN"},
		Help:    "set SHORT_URL_LEN variable, must be > 0",
		Args: func(a *grumble.Args) {
			a.Int("setlen", "setlen")
		},
		Run: cmdSetLen,
	})
}

func makeurl(app *grumble.App) {
	app.AddCommand(&grumble.Command{
		Name:    "make",
		Aliases: []string{"MAKE"},
		Help:    "shorten url",
		Args: func(a *grumble.Args) {
			a.String("url", "url")
		},
		Run: cmdMake,
	})
}

func settings(app *grumble.App) {
	app.AddCommand(&grumble.Command{
		Name: "settings",
		Aliases: []string{"setup",
			"config",
			"SETTINGS",
			"SETUP",
			"CONFING"},
		Help: "show settings",
		Run:  cmdSettings,
	})
}

func cmdSettings(c *grumble.Context) error {
	redishandler.CheckSettings()
	redishandler.PrintCLISettings()
	return nil
}

func cmdChangeSetting(cmd, setting string, min, max int, c *grumble.Context) error {
	if c.Args.Int(cmd) < min || c.Args.Int(cmd) > max {
		err := fmt.Errorf("%v variable must be between %v and %v", setting, min, max)
		return err
	}
	redishandler.ChangeSetting(setting, c.Args.Int(cmd))
	return nil
}

func cmdSetLen(c *grumble.Context) error {
	return cmdChangeSetting("setlen", "SHORT_URL_LEN", 1, 20, c)
}

func cmdSetTime(c *grumble.Context) error {
	return cmdChangeSetting("settime", "USER_WAIT_TIME", 1, 1<<20, c)
}

func cmdMake(c *grumble.Context) error {
	longURL, err := processMakeArg(c.Args.String("url"))
	if err != nil {
		log.Println(err)
		return err
	}

	redishandler.CheckSettings()

	if redishandler.ExistsLongURL(longURL.String()) {
		redishandler.PrintExistingShortURL(longURL)
		return nil
	}

	username := redishandler.RandomUser()
	if redishandler.IsUserLimited(username) {
		redishandler.PrintUserWaitTime(username)
		return nil
	}

	shorturl, err := redishandler.ShortenURL(longURL)
	if err != nil {
		log.Println(err)
		return err
	}
	redishandler.InsertURLData(longURL, shorturl, username)
	return nil
}

func processMakeArg(arg string) (longURL *urlhandler.NiqURL, err error) {
	longURL, err = urlhandler.InputStringToNiqURL(arg)
	if err != nil {
		log.Println(err)
		return longURL, err
	}
	return longURL.AddHTTPSSchemeIfNonAbs(), nil
}
