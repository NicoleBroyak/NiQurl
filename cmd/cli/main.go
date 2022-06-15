package main

import (
	"github.com/nicolebroyak/niqurl/pkg/cli"
	"github.com/nicolebroyak/niqurl/tools/redishandler"
)

func main() {
	redishandler.Start()
	defer redishandler.Client.Close()
	cli.Start()
}
