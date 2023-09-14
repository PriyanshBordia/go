package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/PriyanshBordia/go/pkg/models"
)


var book models.Book


func Home(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(http.StatusOK)
	writter.Write(make([]byte, 1))
}

func AddBook(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
}

func GetBook(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	books := models.GetBooks()
	res, _ := json.Marshal(books)
	writter.WriteHeader(http.StatusOK)
	writter.Write(res)
}	


func GetBooks(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	books := models.GetBooks()
	res, _ := json.Marshal(books)
	writter.WriteHeader(http.StatusOK)
	writter.Write(res)
}	

func UpdateBook(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")

}

func DeleteBook(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")

}
