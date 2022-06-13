package cli

import (
	"redishandler"

	"github.com/desertbit/grumble"
)

func initialize(a *grumble.App, flags grumble.FlagMap) error {
	redishandler.ManageSettings()
	GFUflag(a, flags)
	return nil
}
