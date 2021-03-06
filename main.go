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
	s.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	s.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	log.Fatal(srv.ListenAndServe())
}
