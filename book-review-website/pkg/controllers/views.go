package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PriyanshBordia/go/pkg/models"
	"github.com/PriyanshBordia/go/pkg/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)


var print = fmt.Println
// var book models.Book


func Home(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(http.StatusOK)
	writter.Write(make([]byte, 1))
}

func AddBook(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	book := &models.Book{}
	utils.ParseBody(request, book)
	book = models.AddBook(book)
	res, _ := json.Marshal(book)
	writter.WriteHeader(http.StatusOK)
	writter.Write(res)
}

func GetBook(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		print(err.Error())
		writter.WriteHeader(http.StatusForbidden)
		writter.Write([]byte(err.Error()))
		return
	}
	book, db := models.GetBook(id)
	// print(db.Error.Error() == gorm.ErrRecordNotFound.Error())
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		print(db.Error)
		writter.WriteHeader(http.StatusForbidden)
		writter.Write([]byte(gorm.ErrRecordNotFound.Error()))
		return
	}
	res, _ := json.Marshal(book)
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
	updated_book := &models.Book{}
	utils.ParseBody(request, updated_book)
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		writter.WriteHeader(http.StatusForbidden)
		writter.Write([]byte(err.Error()))
		return
	}
	book := models.UpdateBook(id, updated_book)
	res, _ := json.Marshal(book)
	writter.WriteHeader(http.StatusOK)
	writter.Write(res)
}

func DeleteBook(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		writter.WriteHeader(http.StatusForbidden)
		writter.Write([]byte(err.Error()))
		return
	}
	_ = models.DeleteBook(id)
	res, _ := json.Marshal("Book deleted.")
	writter.WriteHeader(http.StatusOK)
	writter.Write(res)
}
