package main

import (
	"flag"
	"log"
	"net/http"
	"tempograph/api"
)

func main() {
	httpAddr := flag.String("http", ":8080", "HTTP service address")
	flag.Parse()

	log.Println("Starting TempoGraph on", *httpAddr)
	if err := http.ListenAndServe(*httpAddr, api.NewRouter()); err != nil {
		log.Fatal("HTTP server error:", err)
	}
}
