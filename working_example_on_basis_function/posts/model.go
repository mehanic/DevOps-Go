package posts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	var p1 Post
	err = json.Unmarshal(jsonData, &p1)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(p1.Title)
	fmt.Println(p1.Body)
}

func main1() {
	//send request and get response
	url := "http://jsonplaceholder.typicode.com/posts"
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
	var postsData []Post
	err = json.Unmarshal(jsonData, &postsData)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, v := range postsData {
		v.ShowInfo()
	}
}

func main2() {
	url := "http://jsonplaceholder.typicode.com/posts"
	p1 := Post{
		UserId: 12,
		Title:  "hello from yerassyl",
		Body:   "1321212312",
	}
	jsonData, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
		return
	}
	body := bytes.NewReader(jsonData)
	response, err := http.Post(url, "application/json", body)
	if err != nil {
		log.Fatal(err)
		return
	}
	jsonResponseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	p2 := Post{}
	err = json.Unmarshal(jsonResponseData, &p2)
	if err != nil {
		log.Fatal(err)
		return
	}
	p2.ShowInfo()
}

