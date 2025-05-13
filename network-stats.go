package main

import "net/http"

func WithLogging(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		print("logging " + r.URL.Path)
		handler(w, r)
	}
}
