package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type MovieDetails struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	OpeningCrawl  string   `json:"opening_crawl"`
	CharactersIds []string `json:"characters_ids"`
}

func searchMovies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := SWAPIGet("films/?title=" + query)
	if err != nil {
		return
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	data := &SWAPIMovieSearchResponse{}
	err = json.NewDecoder(res.Body).Decode(data)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data.ToSearchResults())
	if err != nil {
	}
}

func getMovieDetails(w http.ResponseWriter, r *http.Request) {
	movieId := r.URL.Query().Get("movieId")
	if movieId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := SWAPIGet("films/" + movieId)
	if err != nil {
		return
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	data := &SWAPIMovieDetails{}
	err = json.NewDecoder(res.Body).Decode(data)
	if err != nil {
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data.ToMovieDetails())
	if err != nil {
	}
}
