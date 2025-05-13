package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func searchCharacters(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := http.Get(os.Getenv("SW_BASE_URL") + "people/?name=" + query)
	if err != nil {
		return
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	data := &SWAPICharacterSearchResponse{}
	err = json.NewDecoder(res.Body).Decode(data)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data.ToSearchResults())
	if err != nil {
		return
	}
}
