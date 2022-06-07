package cli

import (
	"errors"

	"github.com/desertbit/grumble"
)

func GenerateFakeUsers(a *grumble.App, flags grumble.FlagMap) error {
	fakeusers := flags.Int("generate-fake-users")
	if fakeusers > 1000 {
		return errors.New("You can generate max 1000 users one time")
	}
	if fakeusers > 0 {
		err := QueryFakeUsers(fakeusers)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Number must be higher than 0")
}

// Equivalent of GenerateFakeUsers to run tests
func generateFakeUsers(fakeusers int) error {
	if fakeusers > 1000 {
		return errors.New("You can generate max 1000 users one time")
	}
	if fakeusers > 0 {
		err := QueryFakeUsers(fakeusers)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Number must be higher than 0")
}
