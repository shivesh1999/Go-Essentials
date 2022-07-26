package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Add your model
type movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isdn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//global variables
var movies []movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	//setting up the response to json
	w.Header().Set("Content-Type", "application/json")
	//encoding the response to a json
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// getting parameters from the link
	params := mux.Vars(r)
	for index, item := range movies {
		// params["id"] to get the field id from the param
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // append all of the other at index
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//making a variable
	var movie movie
	// decoding the body to the movie struct refernce
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// strconv - for converting into string
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // append all of the other at index
			var movie movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func main() {
	movies = append(movies, movie{ID: "1", Isbn: "12345", Title: "Movie 1", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, movie{ID: "2", Isbn: "21354", Title: "Movie 2", Director: &Director{FirstName: "Steve", LastName: "Smith"}})

	//Make a router
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	// use HandleFunc for assiningning func to the the request, add method you want to use
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("starting server at Port 8000")
	// Declaring the port
	log.Fatal(http.ListenAndServe(":8000", r))
}
