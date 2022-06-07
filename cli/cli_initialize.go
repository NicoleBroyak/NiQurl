package cli

import (
	"errors"
	"os/exec"

	"github.com/desertbit/grumble"
)

func CLIInitialize(a *grumble.App, flags grumble.FlagMap) error {
	err := SetVariables()
	if err != nil {
		return err
	}
	GenerateFakeUsers(a, flags)
	LoadSettings()
	if USER_COUNT < 1 {
		return errors.New("No users detected, exiting app" +
			" please initialize app with " +
			"-generate-fake-users flag")
	}
	return nil
}

// equivalent of CLIInitialize to run tests
func cliInitialize(USER_COUNT int) error {
	if USER_COUNT < 1 {
		exe := exec.Command("niqurl", "-z", "10")
		exe.Run()
		return errors.New("No users detected, exiting app" +
			" please initialize app with " +
			"-generate-fake-users flag")
	}
	return nil
}
