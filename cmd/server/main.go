package main

import (
	"fmt"
	"log"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/nicolebroyak/niqurl/api"
	"github.com/nicolebroyak/niqurl/config/niqurlconfigs"
	"github.com/nicolebroyak/niqurl/tools/redishandler"
)

func main() {
	Client := redishandler.Start("niqurl-redis:6379")
	defer Client.Close()
	StartServer()
}

func StartServer() {
	server := gin.Default()
	tmplPath := "/" + path.Join("niqurl", "api", "templates")
	server.LoadHTMLFiles(
		path.Join(tmplPath, "404.html"),
		path.Join(tmplPath, "inspecturl.html"),
	)
	server.GET("/!:url", api.ShowURLInfo)
	server.GET("/:url", api.RedirectURL)
	port := fmt.Sprintf(":%v", niqurlconfigs.SettingsMap["SERVER_PORT"])
	log.Println(port)
	server.Run(port)
}
