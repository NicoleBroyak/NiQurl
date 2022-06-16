package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/nicolebroyak/niqurl/tools/redishandler"
)

func RedirectURL(c *gin.Context) {
	exists, index := redishandler.IsExistingShortURL(c.Param("url"))
	if !exists {
		NotFound(c)
		return
	}
	rawurl := redishandler.QueryForLongURL(index)
	newurl, _ := url.Parse(rawurl)
	if !newurl.IsAbs() {
		newurl.Scheme = "https"
	}

	c.Redirect(http.StatusMovedPermanently, newurl.String())
}

func InspectURL(c *gin.Context) {
	x, b := FindShortURLInfo(c.Param("url"))
	if b != nil {
		NotFound(c)
		return
	}
	c.HTML(http.StatusOK, "inspecturl.html", x)
}

func NotFound(c *gin.Context) {
	c.HTML(404, "404.html", "")
}

func FindShortURLInfo(url string) (map[string]interface{}, error) {
	z := map[string]interface{}{}

	scanVal, _ := redishandler.Client.ZScan(
		redishandler.Ctx,
		"shorturl",
		0,
		url,
		0,
	).Val()

	// check if value exists
	if len(scanVal) > 0 {
		i, err := strconv.Atoi(scanVal[1])
		if err != nil {
			return z, err
		}
		z["shorturl"] = url

		z["longurl"] = redishandler.Client.ZRange(
			redishandler.Ctx,
			"longurl",
			int64(i),
			int64(i),
		).Val()[0]

		z["user"] = redishandler.Client.ZRange(
			redishandler.Ctx,
			"username",
			int64(i),
			int64(i),
		).Val()[0]
		fmt.Println(z)

		return z, nil
	}
	return z, errors.New("shorturl not found")
}
