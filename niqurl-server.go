package main

import (
	"api"
	"redishandler"
)

func main() {
	redishandler.Start()
	api.StartServer()
	defer redishandler.RDB.Close()
}
