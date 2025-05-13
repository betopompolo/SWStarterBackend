package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type CharacterDetails struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	SkinColor string `json:"skinColor"`
	HairColor string `json:"hairColor"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	BirthYear string `json:"birthYear"`
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
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data.ToCharacterDetails())
	if err != nil {
	}
}
