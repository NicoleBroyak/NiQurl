package cli

import (
	"errors"
	"fmt"
	"log"

	"github.com/nicolebroyak/niqurl/tools/redishandler"
	"github.com/nicolebroyak/niqurl/tools/urlhandler"

	"github.com/desertbit/grumble"
)

func commands(app *grumble.App) {
	cmdSetTime(app)
	cmdSetLen(app)
	cmdMake(app)
	cmdSettings(app)
}

func cmdSetTime(app *grumble.App) {
	app.AddCommand(&grumble.Command{
		Name: "settime",
		Help: "settime [miliseconds] ---> number between 1 and 1048576",
		Args: func(a *grumble.Args) {
			a.Int("settime", "settime")
		},
		Run: setTime,
	})
}

func cmdSetLen(app *grumble.App) {
	app.AddCommand(&grumble.Command{
		Name: "setlen",
		Help: "setlen [number] ---> number between 1 and 20",
		Args: func(a *grumble.Args) {
			a.Int("setlen", "setlen")
		},
		Run: setLen,
	})
}

func cmdMake(app *grumble.App) {
	app.AddCommand(&grumble.Command{
		Name: "make",
		Help: "make [url]",
		Args: func(a *grumble.Args) {
			a.String("url", "url")
		},
		Run: make,
	})
}

func cmdSettings(app *grumble.App) {
	app.AddCommand(&grumble.Command{
		Name: "settings",
		Run:  ShowSettings,
		Help: "shows settings",
	})
}

func ShowSettings(c *grumble.Context) error {
	redishandler.SetInvalidSettingsToDefaults()
	redishandler.PrintCurrentCLISettings()
	return nil
}

func changeSetting(cmd, setting string, min, max int, c *grumble.Context) error {
	if c.Args.Int(cmd) < min || c.Args.Int(cmd) > max {
		err := fmt.Errorf("%v variable must be between %v and %v", setting, min, max)
		return err
	}
	redishandler.ChangeSetting(setting, c.Args.Int(cmd))
	return nil
}

func setLen(c *grumble.Context) error {
	return changeSetting("setlen", "SHORT_URL_LEN", 1, 20, c)
}

func setTime(c *grumble.Context) error {
	return changeSetting("settime", "USER_WAIT_TIME", 1, 1<<20, c)
}

// Multiple functions used to perform make command
// Consolidated into one function to satisfy
// Grumble framework syntax requirements
func make(c *grumble.Context) error {
	// Validate settings and process url input
	redishandler.SetInvalidSettingsToDefaults()
	NiqURL, err := urlhandler.StringToNiqURL(c.Args.String("url"))
	if err != nil {
		return err
	}

	// assign variables used to process function cases
	NiqURL.IfEmptySchemeAddHTTPS()
	NiqURL.UserName = redishandler.GetRandomUser()
	NiqURL.GenerateShortURLPath(redishandler.GetSetting("SHORT_URL_LEN"))

	// check and processs possible conditions
	switch {
	case redishandler.ExistsLongURL(NiqURL.LongURL):
		redishandler.ProcessExistingURL(NiqURL.LongURL)
		return nil
	case redishandler.IsUserOnWaitTime(NiqURL.UserName):
		redishandler.PrintUserWaitTime(NiqURL.UserName)
		return nil
	case redishandler.ExistsShortURL(NiqURL.ShortURL):
		err := errors.New("can't generate short url with" +
			"specified length, please change" +
			"length using setlen command")
		log.Println(err)
		return err
	// finally insert URL if conditions above don't occur
	default:
		redishandler.InsertURLData(NiqURL)
		return nil
	}
}
