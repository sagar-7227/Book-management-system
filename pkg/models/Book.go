package models

import (
	"github.com/sagar-7227/Book-management-system/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.ConnectDatabase()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	result := db.Where("ID=?", id).Find(&book)
	return &book, result
}

func DeleteBook(id int64) Book {
	var book Book
	db.First(&book, id)

	db.Delete(&book)
	return book
}

func (b *Book) UpdateBook() *Book {
	db.Save(&b)
	return b
}
