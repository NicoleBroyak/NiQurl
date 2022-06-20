package main

import (
	"github.com/nicolebroyak/niqurl/pkg/cli"
	"github.com/nicolebroyak/niqurl/tools/redishandler"
)

func main() {
	Client := redishandler.Start("niqurl-redis:6379")
	defer Client.Close()
	cli.Start()
}
