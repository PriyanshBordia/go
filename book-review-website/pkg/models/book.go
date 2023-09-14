package models

import (
	"github.com/PriyanshBordia/go/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.model
	Name string `gorm:""json:"name"`
	Author string `json:"isbn"`
}

func init() {
	config.Connect()
	db = config.get_db()
	db.AutoMigrate(&Book{})
}


func add_book(book *Book) *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func get_book(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", id).Find(&book)
	return &book, db
}

func get_books() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func delete_book(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}