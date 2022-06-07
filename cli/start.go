package cli

import (
	"fmt"

	"github.com/desertbit/grumble"
)

func Start() {
	App := MakeApp()
	fmt.Println(App.Config().Description)
	grumble.Main(App)
}
