package cli

import (
	"fmt"

	"github.com/desertbit/grumble"
	"github.com/fatih/color"
)

var App = grumble.New(&grumble.Config{
	Name: "NiQurl",
	Description: "\n\n*    .##....##.####..#######..##.....##.########..##......\n" +
		"*    .###...##..##..##.....##.##.....##.##.....##.##......\n" +
		"*    .####..##..##..##.....##.##.....##.##.....##.##......\n" +
		"*    .##.##.##..##..##.....##.##.....##.########..##......\n" +
		"*    .##..####..##..##..##.##.##.....##.##...##...##......\n" +
		"*    .##...###..##..##....##..##.....##.##....##..##......\n" +
		"*    .##....##.####..#####.##..#######..##.....##.########\n\n" +
		"            NiQurl - Simple URL Shortening App\n\n\n            ",
	Prompt:                "NiQurl>",
	PromptColor:           color.New(color.BgMagenta, color.Bold, color.FgBlack),
	HelpHeadlineColor:     color.New(color.FgMagenta),
	HelpHeadlineUnderline: true,
	HelpSubCommands:       true,
	Flags: func(f *grumble.Flags) {
		f.Int("z", "generate-fake-users", 0, "help string")
	},
})

// Create CLI commands
func init() {
	MakeURL()
	SetLen()
	SetTime()
	CLISettings()
	App.OnInit(CLIInitialize)
}

func Start() {
	fmt.Println(App.Config().Description)
	grumble.Main(App)
}
