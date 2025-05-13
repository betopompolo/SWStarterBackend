package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hi there, I love %s!", request.URL.Path[1:])
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
