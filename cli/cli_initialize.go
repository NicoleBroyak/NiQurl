package cli

import (
	"log"
	r "redishandler"

	"github.com/desertbit/grumble"
)

func CLIInitialize(a *grumble.App, flags grumble.FlagMap) error {
	RDB := r.RedisStart()
	SetVariables()
	GenerateFakeUsers(a, flags)
	LoadSettings(RDB)
	if USER_COUNT < 1 {
		log.Fatalln("No users detected, exiting app, please run app with -generate-fake-users flag")
	}
	return nil
}
