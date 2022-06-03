package cli

import "github.com/desertbit/grumble"

func CLIInitialize(a *grumble.App, flags grumble.FlagMap) error {
	SetVariables()
	GenerateFakeUsers(a, flags)
	return nil
}
