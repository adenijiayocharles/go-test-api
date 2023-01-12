package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct {
	ID 			string 	`json: "id`
	Title 		string	`json: "title`
	Author 		string	`json: "author`
	Quantity 	int		`json: quantity`
}

var books = []book {
	{ID: "1", Title: "In Search of Truth", Author: "Marcel Porcoit", Quantity: 2},
	{ID: "2", Title: "In Search of Dare and Policing", Author: "Baba Blu", Quantity: 23},
	{ID: "3", Title: "Peter Obi, the man of the people", Author: "Adamu Garba", Quantity: 29},
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context){
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main () {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}