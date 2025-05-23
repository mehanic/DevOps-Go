package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://golang.org",
		"https://www.github.com",
	}
	c := NewClient(http.DefaultClient, len(urls))
	FetchAll(urls, c)

	for i := 0; i < len(urls); i++ {
		select {
		case resp := <-c.Resp:
			fmt.Printf("Status received for %s: %d\n", resp.Request.URL, resp.StatusCode)
		case err := <-c.Err:
			fmt.Printf("Error received: %s\n", err)
		}
	}
}

//---

// NewClient creates a new client and sets its appropriate channels
func NewClient(client *http.Client, bufferSize int) *Client {
	respch := make(chan *http.Response, bufferSize)
	errch := make(chan error, bufferSize)
	return &Client{
		Client: client,
		Resp:   respch,
		Err:    errch,
	}
}

// Client stores a client and has two channels to aggregate
// responses and errors
type Client struct {
	*http.Client
	Resp chan *http.Response
	Err  chan error
}

// AsyncGet performs a Get then returns
// the resp/error to the appropriate channel
func (c *Client) AsyncGet(url string) {
	resp, err := c.Get(url)
	if err != nil {
		c.Err <- err
		return
	}
	c.Resp <- resp
}

//----

// FetchAll grabs a list of urls
func FetchAll(urls []string, c *Client) {
	for _, url := range urls {
		go c.AsyncGet(url)
	}
}

// go run async.go                                                                                                                                               ──(Fri,Jan17)─┘
// Status received for https://github.com/: 200
// Status received for https://www.google.com: 200
// Status received for https://go.dev/: 200
