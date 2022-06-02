package cli

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func Start() {
	// set args for examples sake
	os.Args = []string{"greet", "--generate-bash-completion"}

	app := cli.NewApp()
	app.Name = "greet"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:        "describeit",
			Aliases:     []string{"d"},
			Usage:       "use it to see a description",
			Description: "This is how we describe describeit the function",
			Action: func(c *cli.Context) error {
				fmt.Printf("i like to describe things")
				return nil
			},
		}, {
			Name:        "next",
			Usage:       "next example",
			Description: "more stuff to see when generating bash completion",
			Action: func(c *cli.Context) error {
				fmt.Printf("the next example")
				return nil
			},
		},
	}

	_ = app.Run(os.Args)
}
