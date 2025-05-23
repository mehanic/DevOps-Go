package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Go tutorial for POST request")
	postJsonRequest()
}

func postJsonRequest() {
	const myurl = "http://localhost:8000/post"

	req := strings.NewReader(`
		{
			"coursename":"test",
			"price":0,
			"platform":"rohan.com"
		}
	`)
	res, err := http.Post(myurl, "application/json", req)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}
	defer res.Body.Close()

	cont, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(cont))
}
