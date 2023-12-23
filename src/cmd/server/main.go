package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restfulGoCart/src/internal/api"
)

func main() {
	r := mux.NewRouter()

	api.SetupRoutes(r)

	// Start server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
