package cli

import (
	"fmt"
	"redishandler"

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
		"            NiQurl - Simple URL Shortening App\n\n\n" +
		"NiQurl is running. \nTo shorten url use command " +
		"\"make https://example.com\"\nUse help command to read documentation",
	Prompt:                "NiQurl>",
	PromptColor:           color.New(color.BgMagenta, color.Bold, color.FgBlack),
	HelpHeadlineColor:     color.New(color.FgMagenta),
	HelpHeadlineUnderline: true,
	HelpSubCommands:       true,
	Flags: func(f *grumble.Flags) {
		f.Int("z", "generate-fake-users", 0, "help string")
	},
})

func init() {
	RDB := redishandler.RedisStart()
	Ctx := redishandler.Ctx
	MakeURL(RDB)
	SetLen(RDB)
	SetTime(RDB)
	CLISettings(RDB, Ctx)
	App.OnInit(CLIInitialize)
}

func Start() {
	fmt.Println(App.Config().Description)
	App.OnInit(CLIInitialize)
	grumble.Main(App)
}
