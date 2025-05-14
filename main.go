package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"
)

var urlQueue = make(chan string, 100)
var db = NewInMemoryDB()

func getFirst(n int, ns []NetworkStats) []NetworkStats {
	if len(ns) < n {
		n = len(ns)
	}

	return ns[:n]
}

const defaultRecalcTimeInMinutes = 5

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/searchMovies", WithLogging(searchMovies))
	http.HandleFunc("/searchCharacters", WithLogging(searchCharacters))
	http.HandleFunc("/getMovieDetails", WithLogging(getMovieDetails))
	http.HandleFunc("/getCharacterDetails", WithLogging(getCharacterDetails))
	http.HandleFunc("/getNetworkStats", WithLogging(func(writer http.ResponseWriter, request *http.Request) {
		dbData := db.ReadNetworkStats()

		sort.Slice(dbData, func(i, j int) bool {
			return dbData[i].UsageCount > dbData[j].UsageCount
		})

		err = json.NewEncoder(writer).Encode(getFirst(5, dbData))
	}))

	recalcMins, err := strconv.Atoi(os.Getenv("RECALCULATE_NETWORK_STATS_MINUTES"))
	if err != nil || recalcMins <= 0 {
		recalcMins = defaultRecalcTimeInMinutes
	}
	ticker := time.NewTicker(time.Duration(recalcMins) * time.Minute)
	defer ticker.Stop()
	go func() {
		for {
			select {
			case <-ticker.C:
				computeNetworkStats(urlQueue)
			}
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
