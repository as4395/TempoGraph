package graph

import (
	"encoding/json"
	"net/http"
	"time"

	"tempograph/storage"
	"tempograph/util"
)

// HandleDiff compares two timestamps and returns changes
func HandleDiff(w http.ResponseWriter, r *http.Request) {
	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")

	t1, err1 := util.ParseTime(fromStr)
	t2, err2 := util.ParseTime(toStr)
	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid time range", http.StatusBadRequest)
		return
	}

	events := storage.DiffEvents(t1, t2)
	json.NewEncoder(w).Encode(events)
}
