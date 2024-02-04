package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tttulepbergen/tsis1/internal/handlers"
)

func main() {
	log.Println("Starting Movie API server")
	router := mux.NewRouter()

	log.Println("Creating routes")
	router.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	router.HandleFunc("/movies/title/{title}", handlers.GetMovieByTitle).Methods("GET")
	router.HandleFunc("/movies/{id}", handlers.GetMovieById).Methods("GET")
	router.HandleFunc("/health-checking", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/movies/genre/{genre}", handlers.GetMoviesByGenre).Methods("GET")
	router.HandleFunc("/movies/director/{director}", handlers.GetMoviesByDirector).Methods("GET")
	http.Handle("/", router)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
