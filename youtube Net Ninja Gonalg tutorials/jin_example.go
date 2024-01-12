package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var items = []Item{
	{ID: 1, Title: "Item 1"},
	{ID: 2, Title: "Item 2"},
}

func main() {
	r := gin.Default()

	r.GET("/items", getItems)
	r.GET("/items/:id", getItem)
	r.POST("/items", postItem)
	r.PUT("/items/:id", updateItem)
	r.DELETE("/items/:id", deleteItem)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func getItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, a := range items {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func postItem(c *gin.Context) {
	var newItem Item

	// Call BindJSON to bind the received JSON to newItem.
	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	// Add the new item to the slice.
	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

func updateItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedItem Item

	if err := c.BindJSON(&updatedItem); err != nil {
		return
	}

	for i, a := range items {
		if a.ID == id {
			items[i] = updatedItem
			c.IndentedJSON(http.StatusOK, updatedItem)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func deleteItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, a := range items {
		if a.ID == id {
			items = append(items[:i], items[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "item deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}
