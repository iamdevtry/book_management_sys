package models

import (
	"devtry.net/management_book/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
	result := db.Create(&b)
	return b, result.Error
}

func GetBooks() ([]Book, error) {
	var books []Book
	result := db.Find(&books)
	return books, result.Error
}

func GetBook(id string) (*Book, error) {
	var book Book
	result := db.First(&book, id)
	return &book, result.Error
}

func DeleteBook(id string) error {
	result := db.Delete(&Book{}, id)
	return result.Error
}

func (b *Book) UpdateBook() (*Book, error) {
	result := db.Save(&b)
	return b, result.Error
}
