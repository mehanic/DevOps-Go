package main

import (
	"net/http"
)

type urlstatus struct {
	url    string
	status bool
}

func checkurl(url string, c chan urlstatus) {
	_, err := http.Get(url)
	if err != nil {
		c <- urlstatus{url, false}
	} else {
		c <- urlstatus{url, true}

	}

}

func main() {
	c := make(chan urlstatus)
	urls := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"https://google.com",
		"https://google.com",
	}
	for _, url := range urls {
		go checkurl(url, c)
	}

	results := make([]urlstatus, len(urls))

	for _, result := range results {
		result = <-c
		if result.status {
			println(result.url, "is up")

		} else {
			println(result.url, "is down!!!!!!!!!!!")

		}

	}
}

// http://localhost:8081 is down!!!!!!!!!!!
// http://localhost:8082 is down!!!!!!!!!!!
// https://google.com is up
// https://google.com is up
