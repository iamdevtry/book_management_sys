package controllers

import (
	"net/http"

	"devtry.net/management_book/pkg/models"
	"devtry.net/management_book/pkg/utils"

	"github.com/gin-gonic/gin"
)

var NewBook models.Book

func CreateBook(c *gin.Context) {
	err := utils.ParseBody(c.Request, &NewBook)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	book, err := NewBook.CreateBook()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, book)
}

func GetBooks(c *gin.Context) {
	books, err := models.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	book, err := models.GetBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	book, err := models.GetBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err = utils.ParseBody(c.Request, book)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	book, err = book.UpdateBook()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	err := models.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book successfully deleted"})
}
