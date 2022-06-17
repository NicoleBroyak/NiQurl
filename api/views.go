package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicolebroyak/niqurl/tools/redishandler"
)

func RedirectURL(c *gin.Context) {
	if !redishandler.ExistsShortURL(c.Param("url")) {
		notFound(c)
		return
	}

	index, err := redishandler.GetIndexOfShortURL(c.Param("url"))
	if err != nil {
		notFound(c)
		return
	}

	longURL := redishandler.GetLongURL(index)

	c.Redirect(http.StatusMovedPermanently, longURL)
}

func ShowURLInfo(c *gin.Context) {
	urlInfo, err := getURLInfo(c.Param("url"))
	if err != nil {
		notFound(c)
		return
	}

	c.HTML(http.StatusOK, "inspecturl.html", urlInfo)
}

func notFound(c *gin.Context) {
	c.HTML(404, "404.html", "")
}

func getURLInfo(url string) (map[string]string, error) {
	urlInfo := map[string]string{}
	if !redishandler.ExistsShortURL(url) {
		return urlInfo, errors.New("shorturl not found")
	}
	index, _ := redishandler.GetIndexOfShortURL(url)
	urlInfo["shorturl"] = url
	urlInfo["longurl"] = redishandler.GetLongURL(index)
	urlInfo["user"] = redishandler.GetURLAuthor(index)
	return urlInfo, nil
}
