package cli

import (
	"github.com/desertbit/grumble"
)

func GenerateFakeUsers(a *grumble.App, flags grumble.FlagMap) error {
	fakeusers := flags.Int("generate-fake-users")
	if fakeusers > 0 {
		QueryFakeUsers(fakeusers)
	}
	return nil
}
