package main

import (
	"path"

	"github.com/gin-gonic/gin"
	"github.com/nicolebroyak/niqurl/api"
	"github.com/nicolebroyak/niqurl/tools/redishandler"
)

func main() {
	Client := redishandler.Start()
	defer Client.Close()
	StartServer()
}

func StartServer() {
	server := gin.Default()
	tmplPath := "/" + path.Join("go", "app", "api", "templates")
	server.LoadHTMLFiles(
		path.Join(tmplPath, "404.html"),
		path.Join(tmplPath, "inspecturl.html"),
	)
	server.GET("/!:url", api.ShowURLInfo)
	server.GET("/:url", api.RedirectURL)
	server.Run(":8081")
}
