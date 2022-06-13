package api

import (
	"net/http"
	"redishandler"

	"github.com/gin-gonic/gin"
)

func redirectURL(c *gin.Context) {
	url, err := redishandler.RedirectAPI(c.Param("url"))
	if err != nil {
		notFound(c)
		return
	}
	c.Redirect(http.StatusMovedPermanently, url)
}

func viewURL(c *gin.Context) {
	x, b := redishandler.ViewAPI(c.Param("url"))
	if b != nil {
		notFound(c)
		return
	}
	c.HTML(http.StatusOK, "viewurl.html", x)
}

func notFound(c *gin.Context) {
	c.HTML(404, "404.html", "")
}
