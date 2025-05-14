package main

import "sync"

type InMemoryDB struct {
	mux          sync.Mutex
	networkStats []NetworkStats
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		mux:          sync.Mutex{},
		networkStats: []NetworkStats{},
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
		counterMap[ns.url] = ns.usageCount
	}
	for range urlsCount {
		url := <-c
		counterMap[url] += 1
	}

	var stats = make([]NetworkStats, len(counterMap))
	for key, value := range counterMap {
		stats = append(stats, NetworkStats{
			url:        key,
			usageCount: value,
		})
	}
	db.networkStats = stats
}

func (db *InMemoryDB) ReadNetworkStats() []NetworkStats {
	db.mux.Lock()
	defer db.mux.Unlock()
	return db.networkStats
}
