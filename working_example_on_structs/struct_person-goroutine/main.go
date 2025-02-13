package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

const domain = "localhost:8080"

type person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *address `json:"address,omitempty"`
}
type address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []person

func main() {
	addPerson()
	addConfigServer()
}

func addPerson() {
	people = append(people, person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &address{City: "City X", State: "State X"}})
	people = append(people, person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &address{City: "City Z", State: "State Y"}})
	people = append(people, person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
}

func addConfigServer() {
	log.Println("Server Starting!")
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/contact", findAllContact).Methods("GET")
	router.HandleFunc("/contact/{id}", findByIDContact).Methods("GET")
	router.HandleFunc("/contact", createContact).Methods("POST")
	router.HandleFunc("/contact/{id}", deleteContact).Methods("DELETE")
	http.Handle("/", router)

	log.Printf("Add routers on: http://%s/", domain)

	log.Fatal(http.ListenAndServe(domain, router))

	srv := &http.Server{
		Addr:         domain,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	// Running the server in a goroutine.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// Setting the keys "Ctrl + C" (SIGNINT) as command to shutdown the server.
	// SIGKILL, SIGQUIT ou SIGTERM (Ctrl + /) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Locking the system until receiving the shutdown signal.
	<-c

	// Setting a waiting period so that all connections are safely closed.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	// Optionally, you can execute "srv.Shutdown" on a goroutine and lock on "<-ctx.Done ()"
	// if your application should wait for other services to terminate based on context cancellation.
	log.Println("Server Shutting Down!")
	os.Exit(0)
}

func findAllContact(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func findByIDContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&person{})
}

func createContact(w http.ResponseWriter, r *http.Request) {
	var person person
	_ = json.NewDecoder(r.Body).Decode(&person)
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

// 2025/02/13 19:17:49 Server Starting!
// 2025/02/13 19:17:49 Add routers on: http://localhost:8080/
