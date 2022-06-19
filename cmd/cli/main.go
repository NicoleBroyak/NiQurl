package main

import (
	"github.com/nicolebroyak/niqurl/config/niqurlconfigs"
	"github.com/nicolebroyak/niqurl/pkg/cli"
	"github.com/nicolebroyak/niqurl/tools/redishandler"
)

func main() {
	Client := redishandler.Start(niqurlconfigs.RedisHost)
	defer Client.Close()
	cli.Start()
}
