package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	First string `json:"first"`
	Last string `json:"last"`
}

var print = fmt.Println

var directors = make(map[string]Director)
var movies = make(map[string]Movie)

func get_directors(writter http.ResponseWriter, router *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writter).Encode(directors)
}

func get_director(writter http.ResponseWriter, router *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	director := directors[params["first"] + params["last"]]
	json.NewEncoder(writter).Encode(director)
	return
}

func add_director(writter http.ResponseWriter, router *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	director := Director{}
	_ = json.NewDecoder(router.Body).Decode(&director)
	directors[director.First + director.Last] = director
	json.NewEncoder(writter).Encode(director)
}

func update_director(writter http.ResponseWriter, router *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	first := params["first"]
	last := params["last"]
	delete(directors, first + last)
	director := Director{}
	_ = json.NewDecoder(router.Body).Decode(&director)
	directors[director.First + director.Last] = director
	json.NewEncoder(writter).Encode(director)
}

func delete_director(writter http.ResponseWriter, router *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	first := string(params["first"])
	last := string(params["last"])
	delete(directors, first + last)
}

func get_movies(writter http.ResponseWriter, router *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writter).Encode(movies)
}

func get_movie(writter http.ResponseWriter, router *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	movie := movies[params["id"]]
	json.NewEncoder(writter).Encode(movie)
	return
}

func add_movie(writter http.ResponseWriter, router *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	movie := Movie{}
	_ = json.NewDecoder(router.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100))
	movies[movie.ID] = movie
	json.NewEncoder(writter).Encode(movie)
}

func update_movie(writter http.ResponseWriter, router *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	id := params["id"]
	delete(movies, id)
	movie := Movie{}
	_ = json.NewDecoder(router.Body).Decode(&movie)
	movie.ID = id
	movies[id] = movie
	json.NewEncoder(writter).Encode(movie)
}

func delete_movie(writter http.ResponseWriter, router *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	id := string(params["id"])
	delete(movies, id)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/directors", get_directors).Methods("GET")
	router.HandleFunc("/director/{first}/{last}", get_director).Methods("GET")
	router.HandleFunc("/director/add", add_director).Methods("POST")
	router.HandleFunc("/director/update/{first}/{last}", update_director).Methods("PUT")
	router.HandleFunc("/director/delete/{first}/{last}", delete_director).Methods("DELETE")

	router.HandleFunc("/movies", get_movies).Methods("GET")
	router.HandleFunc("/movie/{id}", get_movie).Methods("GET")
	router.HandleFunc("/movie/add", add_movie).Methods("POST")
	router.HandleFunc("/movie/update/{id}", update_movie).Methods("PUT")
	router.HandleFunc("/movie/delete/{id}", delete_movie).Methods("DELETE")

	print("Starting server...")
	http.ListenAndServe(":8000", router)
}