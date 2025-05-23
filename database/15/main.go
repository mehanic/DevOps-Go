package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kirigaikabuto/Golang1300Lessons/15/mdw"
	"github.com/kirigaikabuto/Golang1300Lessons/15/redis_connect"
	"github.com/kirigaikabuto/Golang1300Lessons/15/users"
	"log"
	"net/http"
)

var (
	PORT = "8080"
)

func main() {
	redisConfig := redis_connect.RedisConfig{
		Host: "localhost",
		Port: "6379",
	}
	redisConnect, err := redis_connect.NewRedisConnectStore(redisConfig)
	if err != nil {
		log.Fatal(err)
		return
	}
	mongoConfig := users.MongoConfig{
		Host:           "mongodb://localhost:27017",
		Port:           "27017",
		Database:       "login_example",
		CollectionName: "users",
	}
	mongoUsersStore, err := users.NewUsersStore(mongoConfig)
	if err != nil {
		log.Fatal(err)
		return
	}
	usersHttpEndpoints := users.NewUsersHttpEndpoints(mongoUsersStore, redisConnect)
	middleware := mdw.NewMiddleware(redisConnect)
	r := mux.NewRouter()
	r.Methods("POST").Path("/login").HandlerFunc(usersHttpEndpoints.Login())
	r.Methods("POST").Path("/register").HandlerFunc(usersHttpEndpoints.Register())
	r.Methods("GET").Path("/get-info").HandlerFunc(middleware.LoginMiddleware(usersHttpEndpoints.GetInfo()))
	fmt.Println("Server is running on port " + PORT)
	http.ListenAndServe(":"+PORT, r)
}
