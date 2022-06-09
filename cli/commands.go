package cli

import "github.com/desertbit/grumble"

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
		Run: cmdMakeURL,
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
