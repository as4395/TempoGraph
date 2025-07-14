package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"tempograph/graph"
	"tempograph/query"
	"tempograph/storage"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()

	// Graph mutation handlers
	router.HandleFunc("/graph", graph.HandleGraphMutation).Methods("POST", "DELETE")

	// Query handlers
	router.HandleFunc("/snapshot", query.HandleSnapshot).Methods("GET")
	router.HandleFunc("/diff", query.HandleDiff).Methods("GET")
	router.HandleFunc("/history/{id}", query.HandleHistory).Methods("GET")

	return router
}
