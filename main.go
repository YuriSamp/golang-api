package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
}

var albums = []album{
	{ID: "1", Title: "Jubile", Artist: "Japonese Breakfast"},
	{ID: "2", Title: "Whole lotta red", Artist: "playboi Carti"},
	{ID: "3", Title: "Bo Jackson", Artist: "Boldy James"},
	{ID: "4", Title: "Call me if you get lost", Artist: "Tyler the creator"},
	{ID: "5", Title: "Circles", Artist: "Mac Miller"},
	{ID: "6", Title: "After hours", Artist: "The weeknd"},
	{ID: "7", Title: "Pray for paris", Artist: "Westside Gunn"},
	{ID: "8", Title: "The Forever Story", Artist: "JID"},
	{ID: "9", Title: "LP!", Artist: "JPEGMAFIA"},
	{ID: "10", Title: "Never Enough", Artist: "Daniel Caesar"},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}


func postAlbums(c *gin.Context) {
	var newAlbum album


	if err := c.BindJSON(&newAlbum); err != nil {
			return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
			if a.ID == id {
					c.IndentedJSON(http.StatusOK, a)
					return
			}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
