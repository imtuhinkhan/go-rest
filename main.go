package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Author   string  `json:"author"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2, Price: 100},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5, Price: 150},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6, Price: 130},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}
