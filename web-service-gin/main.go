package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  int32  `json:"price"`
}

var albums = []album{
	{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 5700},
	{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 1800},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 4000},
}

func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func getAlbum(context *gin.Context) {
	id := context.Param("id")
	for _, album := range albums {
		if album.Id != id {
			continue
		}
		context.IndentedJSON(http.StatusOK, album)
		return
	}

	message := fmt.Sprintf("album(id=%v) not found", id)
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": message})
}

func postAlbum(context *gin.Context) {
	var newAlbum album
	err := context.BindJSON(&newAlbum)
	if err != nil {
		return
	}
	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbum)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8000")
}
