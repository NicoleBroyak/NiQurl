package main

import (
	"cli"
	"redishandler"
)

func main() {
	redishandler.Start()
	cli.Start()
	defer redishandler.RDB.Close()
}
