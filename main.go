package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var movies []Movie

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {
	movies = append(movies, Movie{ID: "1", Isbn: "12345", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "67890", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	router := mux.NewRouter()
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("Server is running on port 8000")
	log.Println(http.ListenAndServe(":8000", router))
}
