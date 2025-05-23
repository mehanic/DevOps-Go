package main

import "net/http"

func main() {
	urls := []string{"https://ebmsuite.com", "https://youtube.com", "http://127.0.0.1"}

	c := make(chan string)
	for _, url := range urls {

		go func() {
			HeavyTask(c, url)

		}()
	}

	// var responses []string
	for i := range c {
		println(i)
	}

	close(c)

}

func HeavyTask(c chan string, url string) {

	c <- "Processing: " + url

	client := &http.Client{}
	method := "GET"

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		c <- url + " error on request" + err.Error()
	}

	if res, err := client.Do(req); err != nil {

		c <- url + " response error" + err.Error()

	} else {
		c <- url + ": " + res.Status
	}

}

// ─λ go run main.go                                                                        0 (0.001s) < 15:19:01
// Processing: http://127.0.0.1
// Processing: https://ebmsuite.com
// Processing: https://youtube.com
// http://127.0.0.1 response errorGet "http://127.0.0.1": dial tcp 127.0.0.1:80: connect: connection refused
// https://youtube.com: 200 OK
// https://ebmsuite.com: 200 OK
