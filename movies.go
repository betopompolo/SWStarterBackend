package main

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

type MovieDetails struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	OpeningCrawl    string           `json:"opening_crawl"`
	CharactersShort []CharacterShort `json:"characters"`
}

type MovieShort struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (movie *MovieDetails) FetchCharactersShort(charsUrls []string) {
	goRoutinesMaxCount := min(len(charsUrls), 50)
	wg := sync.WaitGroup{}
	wg.Add(goRoutinesMaxCount)
	for _, url := range charsUrls {
		go func() {
			defer wg.Done()
			res, err := http.Get(url)
			if err != nil {
			}
			defer res.Body.Close()
			data := &SWAPICharacterDetails{}
			err = json.NewDecoder(res.Body).Decode(data)
			if err != nil {
				return
			}
			movie.CharactersShort = append(movie.CharactersShort, data.ToCharactersShort())
		}()
	}
	wg.Wait()
}

func FetchMovie(id string) (*MovieDetails, error) {
	res, err := SWAPIGet("films/" + id)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	movie := data.ToMovieDetails()
	movie.FetchCharactersShort(data.Result.Properties.CharactersURLs)

	return movie, nil
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

	movieDetails, err := FetchMovie(movieId)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(movieDetails)
}
