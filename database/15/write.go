package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/kirigaikabuto/Golang1300Lessons/15/users"
	"time"
)

var (
	host = "localhost"
	port = "6379"
)

func main() {
	client := redis.NewClient(&redis.Options{Addr: host + ":" + port})
	err := client.Ping().Err()
	if err != nil {
		panic(err)
		return
	}
	//save data to redis
	key := "one"
	data := &users.User{
		Id:       "1",
		Username: "kirito",
		Password: "1234",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
		return
	}
	err = client.Set(key, jsonData, 3*time.Minute).Err()
	if err != nil {
		panic(err)
		return
	}
}
