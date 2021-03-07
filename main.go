package main

import (
	"go-mux/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	// Routing
	s.HandleFunc("/users", handlers.UserIndex).Methods("GET")
	s.HandleFunc("/users/{id}", handlers.UserShow).Methods("GET")
	s.HandleFunc("/users", handlers.UserCreate).Methods("POST")
	// s.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PATCH")

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	log.Fatal(srv.ListenAndServe())
}
