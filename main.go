package main

import (
	"api"
	"cli"
	"fmt"
	"redis"
)

func main() {
	// 1. Query database if user_count > 0
	// False: panic
	// True: continue
	url, shorturl, err := cli.MakeURL("https://google.com")
	redis.MakeURL(url, shorturl, "Example")
	fmt.Println(url, shorturl, err)
	api.StartRouter()
}
