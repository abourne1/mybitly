package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github/abourne1/mybitly/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/link/", handlers.LinkHandler)
	r.HandleFunc("/stats/", handlers.StatsHandler)
	r.PathPrefix("/{slug}").HandlerFunc(handlers.RedirectHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
