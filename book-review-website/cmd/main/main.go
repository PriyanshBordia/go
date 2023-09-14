package main

import (
	"log"
	"net/http"

	"github.com/PriyanshBordia/go/pkg/routers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routers.BookStoreApp(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", router))
}