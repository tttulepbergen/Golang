package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/tttulepbergen/tsis1/api"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(api.Movies)
}

func GetMovieByTitle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieTitle := params["title"]

	movieTitle = strings.ToLower(movieTitle)

	for _, movie := range api.Movies {
		if strings.ToLower(movie.Title) == movieTitle {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	http.Error(w, "Movie not found", http.StatusNotFound)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]

	for _, movie := range api.Movies {
		if strconv.Itoa(movie.Id) == movieId {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	http.Error(w, "Movie not found", http.StatusNotFound)
}

func GetMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	genre := params["genre"]

	genre = strings.ToLower(genre)
	var matchingMovies []api.Movie

	for _, movie := range api.Movies {
		if strings.ToLower(movie.Genre) == genre {
			matchingMovies = append(matchingMovies, movie)
		}
	}

	if len(matchingMovies) > 0 {
		json.NewEncoder(w).Encode(matchingMovies)
	} else {
		http.Error(w, "Movies with the specified genre not found", http.StatusNotFound)
	}
}

func GetMoviesByDirector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	director := params["director"]

	director = strings.ToLower(director)
	var matchingMovies []api.Movie

	for _, movie := range api.Movies {
		if strings.ToLower(movie.Director) == director {
			matchingMovies = append(matchingMovies, movie)
		}
	}

	if len(matchingMovies) > 0 {
		json.NewEncoder(w).Encode(matchingMovies)
	} else {
		http.Error(w, "Movies by the specified director not found", http.StatusNotFound)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	healthStatus := "App is healthy!\n\n"

	appDescription := "This is a simple web API providing information about movies. It allows users to retrieve movie details, perform searches, and check the overall health of the application.\n"
	authorInfo := "\n\nAuthor: Bahauddin\n"

	healthCheckResponse := healthStatus + appDescription + authorInfo

	w.Write([]byte(healthCheckResponse))
}
