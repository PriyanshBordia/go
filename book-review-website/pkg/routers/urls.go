package routers

import (
	"github.com/gorilla/mux"
)

var BookStoreApp = func(router *mux.Router) {
	router.HandleFunc("/", controllers.home).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.get_book).Methods("GET")
	router.HandleFunc("/books/", controllers.get_books).Methods("GET")
	router.HandleFunc("/book/add", controllers.add_book).Methods("POST")
	router.HandleFunc("/book/update/{id}", controllers.update_book).Methods("PUT")
	router.HandleFunc("/book/delete/{id}", controllers.delete_book).Methods("DELETE")
}

// "github.com/PriyanshBordia/go/book-review-website/pkg/controllers"
