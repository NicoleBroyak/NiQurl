package cli

import (
	"fmt"
	r "redishandler"

	"github.com/desertbit/grumble"
)

func Start() {
	app := makeApp()
	fmt.Println(app.Config().Description)
	grumble.Main(app)
	r.RDB.Close()
}
