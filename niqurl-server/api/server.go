package api

import (
	"os"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	server := gin.Default()
	path, _ := os.Getwd()
	server.LoadHTMLFiles(path+"/api/templates/404.html", path+"/api/templates/viewurl.html")
	server.GET("/!:url", viewURL)
	server.GET("/:url", redirectURL)
	server.NoRoute(notFound)
	server.Run(":8081")
}
