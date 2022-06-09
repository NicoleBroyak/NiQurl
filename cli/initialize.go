package cli

import (
	r "redishandler"

	"github.com/desertbit/grumble"
)

func initialize(a *grumble.App, flags grumble.FlagMap) error {
	r.CheckVariables()
	r.GenerateFakeUsers(a, flags)
	return nil
}
