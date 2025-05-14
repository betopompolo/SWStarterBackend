package main

import (
	"net/http"
	"os"
)

type SWAPIMovieProperties struct {
	Title          string   `json:"title"`
	OpeningCrawl   string   `json:"opening_crawl"`
	CharactersURLs []string `json:"characters"`
}

type SWAPICharacterProperties struct {
	Name       string   `json:"name"`
	Gender     string   `json:"gender"`
	SkinColor  string   `json:"skin_color"`
	HairColor  string   `json:"hair_color"`
	Height     string   `json:"height"`
	EyeColor   string   `json:"eye_color"`
	Mass       string   `json:"mass"`
	BirthYear  string   `json:"birth_year"`
	MoviesURLs []string `json:"movies"`
}
type SWAPIMovieSearchResult struct {
	UID        string               `json:"uid"`
	Properties SWAPIMovieProperties `json:"properties"`
}
type SWAPIMovieSearchResponse struct {
	Results []SWAPIMovieSearchResult `json:"result"`
}

func (res SWAPIMovieSearchResponse) ToSearchResults() []SearchResult {
	var results []SearchResult
	for _, swResult := range res.Results {
		results = append(results, SearchResult{
			Id:   swResult.UID,
			Name: swResult.Properties.Title,
			Type: "movie",
		})
	}
	return results
}

type SWAPICharacterSearchResult struct {
	UID        string                   `json:"uid"`
	Properties SWAPICharacterProperties `json:"properties"`
}

type SWAPICharacterSearchResponse struct {
	Results []SWAPICharacterSearchResult `json:"result"`
}

func (res SWAPICharacterSearchResponse) ToSearchResults() []SearchResult {
	var results []SearchResult
	for _, swResult := range res.Results {
		results = append(results, SearchResult{
			Id:   swResult.UID,
			Name: swResult.Properties.Name,
			Type: "character",
		})
	}
	return results
}

type SWAPIMovieDetailsResult struct {
	UID        string               `json:"uid"`
	Properties SWAPIMovieProperties `json:"properties"`
}
type SWAPIMovieDetails struct {
	Result SWAPIMovieDetailsResult `json:"result"`
}

func (md SWAPIMovieDetails) ToMovieDetails() *MovieDetails {
	return &MovieDetails{
		ID:           md.Result.UID,
		Name:         md.Result.Properties.Title,
		OpeningCrawl: md.Result.Properties.OpeningCrawl,
	}
}

func (md SWAPIMovieDetails) ToMovieShort() MovieShort {
	return MovieShort{
		ID:   md.Result.UID,
		Name: md.Result.Properties.Title,
	}
}

type SWAPICharacterDetails struct {
	Result struct {
		UID        string                   `json:"uid"`
		Properties SWAPICharacterProperties `json:"properties"`
	}
}

func (c SWAPICharacterDetails) ToCharacterDetails() CharacterDetails {
	return CharacterDetails{
		ID:        c.Result.UID,
		Name:      c.Result.Properties.Name,
		Gender:    c.Result.Properties.Gender,
		SkinColor: c.Result.Properties.SkinColor,
		HairColor: c.Result.Properties.HairColor,
		Height:    c.Result.Properties.Height,
		Mass:      c.Result.Properties.Mass,
		BirthYear: c.Result.Properties.BirthYear,
	}
}

func (c SWAPICharacterDetails) ToCharactersShort() CharacterShort {
	return CharacterShort{
		ID:   c.Result.UID,
		Name: c.Result.Properties.Name,
	}
}

func SWAPIGet(path string) (res *http.Response, err error) {
	return http.Get(os.Getenv("SW_BASE_URL") + path)
}
