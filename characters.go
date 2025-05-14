package main

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

type CharacterDetails struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Gender      string       `json:"gender"`
	SkinColor   string       `json:"skinColor"`
	HairColor   string       `json:"hairColor"`
	Height      string       `json:"height"`
	Mass        string       `json:"mass"`
	BirthYear   string       `json:"birthYear"`
	MoviesShort []MovieShort `json:"moviesShort"`
}

func (character *CharacterDetails) FetchMoviesShort(moviesUrls []string) {
	goRoutinesMaxCount := min(len(moviesUrls), 50)
	if goRoutinesMaxCount == 0 {
		character.MoviesShort = []MovieShort{}
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(goRoutinesMaxCount)
	for _, url := range moviesUrls {
		go func() {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
			}
			defer resp.Body.Close()
			data := &SWAPIMovieDetails{}
			err = json.NewDecoder(resp.Body).Decode(data)
			if err != nil {
				return
			}
			character.MoviesShort = append(character.MoviesShort, data.ToMovieShort())
		}()
	}
	wg.Wait()
}

type CharacterShort struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func searchCharacters(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := SWAPIGet("people/?name=" + query)
	if err != nil {
		return
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
		}
	}(res.Body)
	data := &SWAPICharacterSearchResponse{}
	err = json.NewDecoder(res.Body).Decode(data)
	if err != nil {
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data.ToSearchResults())
	if err != nil {
	}
}

func getCharacterDetails(w http.ResponseWriter, r *http.Request) {
	characterId := r.URL.Query().Get("characterId")
	if characterId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := SWAPIGet("people/" + characterId)
	if err != nil {
		return
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
		}
	}(res.Body)
	data := &SWAPICharacterDetails{}
	err = json.NewDecoder(res.Body).Decode(data)
	if err != nil {
		return
	}

	character := data.ToCharacterDetails()
	character.FetchMoviesShort(data.Result.Properties.MoviesURLs)
	err = json.NewEncoder(w).Encode(character)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")

}
