package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"

	"github.com/gorilla/mux"

	"github/abourne1/mybitly/config"
	dbLib "github/abourne1/mybitly/db"
	"github/abourne1/mybitly/handlers"
)

func main() {
	// Connect to Postgres SB
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", 
		config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatalf("Open err: %s", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Ping err: %s", err.Error())
	}
	defer db.Close()

	// Create handler
	dbImpl := dbLib.New(db)
	handler := handlers.New(dbImpl)

	// Define routes and serve
	r := mux.NewRouter()
	r.HandleFunc("/link/", handler.Link)
	r.HandleFunc("/stats/", handler.Stats)
	r.PathPrefix("/{slug}").HandlerFunc(handler.Redirect)
	log.Fatal(http.ListenAndServe(":8080", r))
}
