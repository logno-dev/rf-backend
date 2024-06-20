package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type item struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

var items = []item{
	{ID: "1", Title: "HTMX 2.0", Description: "Vintage floppy disk with printed HTMX 2.0 logo", Price: 9.99},
	{ID: "2", Title: "NeoVim", Description: "Vintage floppy disk with printed NeoVim logo", Price: 9.99},
	{ID: "3", Title: "Solid.js", Description: "Vintage floppy disk with printed Solid.js logo", Price: 9.99},
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func main() {
	router := gin.Default()
	router.GET("/items", getItems)

	router.Run("localhost:3001")

}
