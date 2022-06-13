package api

import (
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	server := gin.Default()
	serverPath, _ := os.Getwd()
	tmplPath := path.Join(serverPath, "api", "templates")
	server.LoadHTMLFiles(
		path.Join(tmplPath, "404.html"),
		path.Join(tmplPath, "viewurl.html"),
	)
	server.GET("/!:url", viewURL)
	server.GET("/:url", redirectURL)
	server.NoRoute(notFound)
	server.Run(":8081")
}
