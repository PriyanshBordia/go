package models

import (
	"github.com/PriyanshBordia/go/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func AddBook(book *Book) *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func GetBook(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", id).Find(&book)
	return &book, db
}

func GetBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func UpdateBook(id int64) (Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", id).Update(book)
	return book, db
}
// book.Name = updated_book.Name
	// book.Author = updated_book.Author

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}