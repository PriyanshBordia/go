package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/PriyanshBordia/go/pkg/models"
)


var book models.Book


func home(writter http.ResponseWriter, request *http.Request) {
	return
}

func add_book(writter http.ResponseWriter, request *http.Request) {
	return
}

func get_books(writter http.ResponseWriter, request *http.Request) {
	books := models.get_books()
	res, _ := json.Marshal(books)
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(http.StatusOK)
	writter.Write(res)
}	

func update_book(writter http.ResponseWriter, request *http.Request) {
	return
}

func delete_book(writter http.ResponseWriter, request *http.Request) {
	return
}
