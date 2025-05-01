package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	route := mux.NewRouter()
	s := route.PathPrefix("/api/task").Subrouter() //Base Path

	//Routes

	s.HandleFunc("", createTask).Methods("POST")
	s.HandleFunc("", getTasks).Methods("GET")
	s.HandleFunc("/{id}", updateTask).Methods("PUT")
	s.HandleFunc("/{id}", deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}