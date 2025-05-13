package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/searchMovies", WithLogging(searchMovies))
	http.HandleFunc("/searchCharacters", WithLogging(searchCharacters))
	http.HandleFunc("/getMovieDetails", WithLogging(getMovieDetails))
	http.HandleFunc("/getCharacterDetails", WithLogging(getCharacterDetails))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
