package query

import (
	"net/http"

	"tempograph/graph"
)

func HandleSnapshot(w http.ResponseWriter, r *http.Request) {
	graph.HandleSnapshot(w, r)
}

func HandleDiff(w http.ResponseWriter, r *http.Request) {
	graph.HandleDiff(w, r)
}

func HandleHistory(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	history := graph.GetHistory(id)
	if history == nil {
		http.Error(w, "Entity not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(history)
}
