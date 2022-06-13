package main

import (
	"api"
	"redishandler"
)

func main() {
	redishandler.Start()
	defer redishandler.RDB.Close()
	api.StartServer()
}
