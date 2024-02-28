package main

import (
	"example/GoTutorial/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting  server")

	router := gin.Default()
	router.GET("/books", models.GetBooks)
	router.GET("/books/:id", models.GetBooksById)
	router.POST("/book", models.AddNewBookToInventory)
	router.PATCH("/books/checkout", models.CheckoutBook)
	router.PATCH("/books/return", models.ReturnBook)

	router.Run("localhost:8090")

}
