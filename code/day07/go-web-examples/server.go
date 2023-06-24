package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Title string `json:"title"`
	Page  string `json:"page"`
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/books/:title/page/:page", GetBookPageById)

	return router
}

func GetBookPageById(c *gin.Context) {
	title := c.Param("title")
	page := c.Param("page")

	book := Book{Title: title, Page: page}

	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := setupRouter()
	router.Run("localhost:8080")
}
