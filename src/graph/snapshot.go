package graph

import (
	"encoding/json"
	"net/http"

	"tempograph/storage"
	"tempograph/util"
)

// HandleSnapshot reconstructs the graph state at a specific time
func HandleSnapshot(w http.ResponseWriter, r *http.Request) {
	tsStr := r.URL.Query().Get("ts")
	ts, err := util.ParseTime(tsStr)
	if err != nil {
		http.Error(w, "Invalid timestamp", http.StatusBadRequest)
		return
	}

	state := storage.Snapshot(ts)
	json.NewEncoder(w).Encode(state)
}
