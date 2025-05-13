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

	http.HandleFunc("/searchMovies", searchMovies)
	http.HandleFunc("/searchCharacters", searchCharacters)
	http.HandleFunc("/getMovieDetails", getMovieDetails)
	http.HandleFunc("/getCharacterDetails", getCharacterDetails)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
