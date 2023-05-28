package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
A GET /books route that lists all books. ( Done )
A GET /books/:id route that shows details of a specific book. (Done)
A POST /books route that allows adding a new book. (Done)
A PUT /books/:id route that allows updating the existing book. (Done)
A DELETE /books/:id route that allows deleting a specific book. (Done)
*/

func main() {
	router := gin.Default()

	router.GET("/", homePage)
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", removeBook)
	router.POST("/books", addNewBook)

	router.Run()
}

func homePage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to My Library"})
}

func getBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"books": books})
}

func getBook(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, val := range books {
		if val.ID == id {
			ctx.JSON(http.StatusOK, val)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func updateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, val := range books {
		if val.ID == id {
			books[i] = updatedBook
			ctx.JSON(http.StatusOK, gin.H{"message": "Book updated"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func removeBook(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, val := range books {
		if val.ID == id {
			books = append(books[:i], books[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func addNewBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	books = append(books, newBook)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Book added"})
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "12 Rules of Life", Author: "Jordan Peterson"},
	{ID: "2", Title: "Harry Potter", Author: "J.K. Rowling"},
}

