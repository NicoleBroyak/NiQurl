package cli

import (
	"github.com/nicolebroyak/niqurl/tools/redishandler"

	"github.com/desertbit/grumble"
)

func initialize(a *grumble.App, flags grumble.FlagMap) error {
	redishandler.SetInvalidSettingsToDefaults()
	generateFakeUsersFlag(a, flags)
	return nil
}
