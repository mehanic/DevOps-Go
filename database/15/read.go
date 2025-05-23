package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/kirigaikabuto/Golang1300Lessons/15/users"
)

var (
	hostRead = "localhost"
	portRead = "6379"
)

func main() {
	client := redis.NewClient(&redis.Options{Addr: hostRead + ":" + portRead})
	err := client.Ping().Err()
	if err != nil {
		panic(err)
		return
	}
	//get data from redis
	key := "one"
	valJson, err := client.Get(key).Result()
	if err != nil {
		panic(err)
		return
	}
	data := &users.User{}
	err = json.Unmarshal([]byte(valJson), &data)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(data)

}
