package handler

import (
	"book-api/app/services"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	//get book by id
	router.GET("/books/", services.GetAllBooks)          //get all books
	router.GET("/books/:bookID", services.GetBookByID)   //get book by id
	router.POST("/books/", services.AddBook)             //create book
	router.PUT("/books/:bookID", services.EditBook)      //update book by id
	router.DELETE("/books/:bookID", services.DeleteBook) //delete book by id
	return router
}
