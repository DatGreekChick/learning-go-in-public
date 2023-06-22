package main

import (
	"github.com/gin-gonic/gin"

	"example.com/api"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/albums", api.GetAlbums)
	router.GET("/albums/:id", api.GetAlbumByID)
	router.POST("/albums", api.PostAlbums)

	return router
}

func main() {
	router := setupRouter()
	router.Run("localhost:8080")
}
