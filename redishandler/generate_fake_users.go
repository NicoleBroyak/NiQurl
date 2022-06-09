package redishandler

import (
	"fmt"

	"github.com/desertbit/grumble"
)

func GenerateFakeUsers(a *grumble.App, flags grumble.FlagMap) error {
	return generateFakeUsersGeneric(flags.Int("generate-fake-users"), 1, 1000)
}

func generateFakeUsersGeneric(num, min, max int) error {
	if num > max || num < min {
		return fmt.Errorf("Num %v is out of range %v and %v", num, min, max)
	}
	err := QueryFakeUsers(num)
	if err != nil {
		return err
	}
	return nil
}
