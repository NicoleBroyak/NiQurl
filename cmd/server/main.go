package main

import (
	"path"

	"github.com/gin-gonic/gin"
	"github.com/nicolebroyak/niqurl/api"
	"github.com/nicolebroyak/niqurl/tools/redishandler"
)

func main() {
	redishandler.Start()
	defer redishandler.Client.Close()
	StartServer()
}

func StartServer() {
	server := gin.Default()
	tmplPath := "/" + path.Join("go", "app", "api", "templates")
	server.LoadHTMLFiles(
		path.Join(tmplPath, "404.html"),
		path.Join(tmplPath, "inspecturl.html"),
	)
	server.GET("/!:url", api.InspectURL)
	server.GET("/:url", api.RedirectURL)
	server.NoRoute(api.NotFound)
	server.Run(":8081")
}
