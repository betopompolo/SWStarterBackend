package main

import "sync"

type InMemoryDB struct {
	mux          sync.Mutex
	networkStats []NetworkStats
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		mux:          sync.Mutex{},
		networkStats: make([]NetworkStats, 0),
	}
}

func (db *InMemoryDB) Update(c chan string) {
	db.mux.Lock()
	defer db.mux.Unlock()

	urlsCount := len(c)
	if urlsCount == 0 {
		return
	}

	var counterMap = map[string]int{}
	for _, ns := range db.networkStats {
		counterMap[ns.URL] = ns.UsageCount
	}
	for range urlsCount {
		url := <-c
		if url == "" {
			continue
		}
		counterMap[url] += 1
	}

	var stats []NetworkStats
	for key, value := range counterMap {
		if value == 0 {
			continue
		}
		stats = append(stats, NetworkStats{
			URL:        key,
			UsageCount: value,
		})
	}
	db.networkStats = stats
}

func (db *InMemoryDB) ReadNetworkStats() []NetworkStats {
	db.mux.Lock()
	defer db.mux.Unlock()
	if db.networkStats == nil {
		db.networkStats = make([]NetworkStats, 0)
	}
	return db.networkStats
}
