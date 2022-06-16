package cli

import (
	"log"

	"github.com/desertbit/grumble"
	"github.com/nicolebroyak/niqurl/tools/randomusers"
	"github.com/nicolebroyak/niqurl/tools/redishandler"
)

func GenerateFakeUsersFlag(a *grumble.App, flags grumble.FlagMap) error {
	num := flags.Int("generate-fake-users")
	if num > 1000 || num < 1 {
		if num != 0 {
			log.Println(("Number of users has to be between 1 and 1000"))
		}
		return nil
	}
	UsersStruct, err := randomusers.GenerateFakeUsers(num)
	if err != nil {
		return err
	}
	redishandler.InsertUsers(UsersStruct)
	return nil
}
