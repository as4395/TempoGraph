package graph

import (
	"encoding/json"
	"net/http"
	"time"

	"tempograph/storage"
	"tempograph/util"
)

type Mutation struct {
	Type       string            `json:"type"` // node or edge
	ID         string            `json:"id"`
	Label      string            `json:"label"`
	Source     string            `json:"source,omitempty"` // for edges
	Target     string            `json:"target,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

// HandleGraphMutation accepts insert/update/delete events with timestamps
func HandleGraphMutation(w http.ResponseWriter, r *http.Request) {
	tsStr := r.URL.Query().Get("ts")
	ts, err := util.ParseTime(tsStr)
	if err != nil {
		http.Error(w, "Invalid timestamp", http.StatusBadRequest)
		return
	}

	var mut Mutation
	if err := json.NewDecoder(r.Body).Decode(&mut); err != nil {
		http.Error(w, "Malformed JSON", http.StatusBadRequest)
		return
	}

	// Append event to the log for durability
	if err := storage.AppendEvent(ts, mut, r.Method); err != nil {
		http.Error(w, "Failed to log mutation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
