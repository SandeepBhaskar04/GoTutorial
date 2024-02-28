package models

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `jason:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

func createBooksInventory() []*Book {
	// function to create books
	books := []*Book{
		{ID: "1", Title: "Atomic Habits", Author: "SomeGreat person", Quantity: 3},
		{ID: "2", Title: "Rich Dad Poor Dad", Author: "Robert Kioski", Quantity: 5},
		{ID: "3", Title: "The power of positive assertions", Author: "SomeOne great", Quantity: 8},
	}

	return books
}

var booksInventory = createBooksInventory()

func GetBooks(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, booksInventory)
}

func fetchBookById(id string) (*Book, error) {
	for _, book := range booksInventory {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, errors.New("book not found")
}

func GetBooksById(c *gin.Context) {
	bookIdToGet := c.Param("id")
	book, err := fetchBookById(bookIdToGet)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func AddNewBookToInventory(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	booksInventory = append(booksInventory, &newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func CheckoutBook(c *gin.Context) {

	book, err := checkBookInInventory(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func ReturnBook(c *gin.Context) {

	book, err := checkBookInInventory(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func checkBookInInventory(c *gin.Context) (*Book, error) {
	bookIdToGet, found := c.GetQuery("id")
	if !found {
		return nil, errors.New("id param missing to get the book")
	}

	book, err := fetchBookById(bookIdToGet)
	if err != nil {
		return nil, errors.New("Book not found")
	}
	return book, nil
}
