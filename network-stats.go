package main

import (
	"net/http"
)

type NetworkStats struct {
	url        string
	usageCount int
}

func WithLogging(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlQueue <- r.URL.Path
		handler(w, r)
	}
}

func computeNetworkStats(c chan string) {
	db.Update(c)
}
