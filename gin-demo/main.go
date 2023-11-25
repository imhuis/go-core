package main

import (
	"gin-demo/repo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	route := gin.Default()
	route.GET("/albums", getAlbums)
	route.GET("/albums/:id", getAlbumByID)
	route.POST("/albums", postAlbums)

	route.Run("localhost:8080")
}

var (
	albums = []repo.Album{
		{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
)

func getAlbums(g *gin.Context) {
	g.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(context *gin.Context) {
	var newAlbum repo.Album
	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
}

func getAlbumByID(context *gin.Context) {
	key := context.Param("id")
	id, _ := strconv.ParseInt(key, 10, 64)
	for _, album := range albums {
		if album.ID == id {
			context.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
