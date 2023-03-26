package services

import (
	"book-api/app/entity"
	"book-api/app/repo/db"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(ctx *gin.Context) {

	result, err := db.GetAllData()

	if err != nil {
		panic(err)
	}

	if len(result) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Tidak ada data buku tersimpan",
		})
	}

	if len(result) != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Success",
			"data":    result,
		})
	}
}

func GetBookByID(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))

	result, err := db.GetDataByID(bookID)
	if err != nil {
		panic(err)
	}

	if len(result) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusBadRequest,
			"message": fmt.Sprintf("Buku dengan id %d tidak ditemukan", bookID),
		})
	}

	if len(result) != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Success",
			"data":    result,
		})
	}
}

func AddBook(ctx *gin.Context) {
	var requestBody entity.Book

	if err := ctx.BindJSON(&requestBody); err != nil {
		panic(err)
	}

	var book = entity.Book{}

	book.Title = requestBody.Title
	book.Author = requestBody.Author
	book.Description = requestBody.Description

	hasil, err := db.InsertData(book)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    hasil,
	})
}

func EditBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	var requestBody entity.Book

	if err := ctx.BindJSON(&requestBody); err != nil {
		panic(err)
	}

	var book = entity.Book{}

	book.Title = requestBody.Title
	book.Author = requestBody.Author
	book.Description = requestBody.Description

	hasil, err := db.UpdateData(bookID, book)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    hasil,
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	hasil, err := db.DeleteData(bookID)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    hasil,
	})
}
