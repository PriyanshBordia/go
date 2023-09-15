package routers

import (
	"github.com/PriyanshBordia/go/pkg/controllers"
	"github.com/gorilla/mux"
)

var BookStoreApp = func(router *mux.Router) {
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/add", controllers.AddBook).Methods("POST")
	router.HandleFunc("/book/update/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/delete/{id}", controllers.DeleteBook).Methods("DELETE")
}