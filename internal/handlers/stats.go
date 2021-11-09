package handlers

import (
	"encoding/json"
	"net/http"
	"runtime"
)

type Stats struct{}

func (s Stats) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var memStat runtime.MemStats
	runtime.ReadMemStats(&memStat)

	stats := struct {
		Alloc        uint64 `json:"alloc"`
		TotalAlloc   uint64 `json:"total_alloc"`
		Sys          uint64 `json:"sys"`
		Mallocs      uint64 `json:"mallocs"`
		Frees        uint64 `json:"frees"`
		LiveObjects  uint64 `json:"live_objects"`
		PauseTotalNs uint64 `json:"pause_total_ns"`
		NumGC        uint32 `json:"num_gc"`
		NumGoroutine int    `json:"num_goroutines"`
	}{
		Alloc:        memStat.Alloc,
		TotalAlloc:   memStat.TotalAlloc,
		Sys:          memStat.Sys,
		Mallocs:      memStat.Mallocs,
		Frees:        memStat.Frees,
		LiveObjects:  memStat.Mallocs - memStat.Frees,
		PauseTotalNs: memStat.PauseTotalNs,
		NumGC:        memStat.NumGC,
		NumGoroutine: runtime.NumGoroutine(),
	}

	bytes, _ := json.Marshal(stats)
	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}
