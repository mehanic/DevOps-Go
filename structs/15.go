package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (p *Post) ShowInfo() {
	message := fmt.Sprintf("Id:%d,Title:%s", p.Id, p.Title)
	fmt.Println(message)
}

func main() {
	//send request and get response
	url := "http://jsonplaceholder.typicode.com/posts/1"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	//response body convert to bytes
	jsonData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	//bytes from response convert to needed struct
	var d map[string]interface{}
	err = json.Unmarshal(jsonData, &d)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(reflect.TypeOf(d["id"]))
}
