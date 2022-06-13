package main

import (
	"cli"
	"redishandler"
)

func main() {
	redishandler.Start()
	defer redishandler.RDB.Close()
	cli.Start()
}
