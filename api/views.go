package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()
	path, _ := os.Getwd()
	router.LoadHTMLFiles(path + "/api/templates/404.html")
	// router.GET("/albums", getAlbums)
	// router.GET("/albums/:id", getAlbumByID)
	// router.GET("/url", viewURL)
	// router.GET("/:url", redirectURL)
	router.GET("/NotFound", notFound)
	// router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func redirectURL(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func viewURL(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
}

func notFound(c *gin.Context) {
	c.HTML(http.StatusOK, "404.html", "")
}

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
