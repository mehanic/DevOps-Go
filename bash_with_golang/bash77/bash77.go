package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Set up an HTTP client with a timeout
	client := http.Client{
		Timeout: 2 * time.Second, // 2 seconds timeout
	}

	// Make a HEAD request to google.com
	resp, err := client.Head("http://google.com")
	if err != nil {
		fmt.Println("The network is down or very slow")
		return
	}
	defer resp.Body.Close()

	// Get the status code and check its first digit
	statusCode := resp.StatusCode
	switch {
	case statusCode >= 200 && statusCode < 300:
		fmt.Println("HTTP connectivity is up")
	case statusCode >= 500 && statusCode < 600:
		fmt.Println("The web proxy won't let us through")
	default:
		fmt.Println("The network is down or very slow")
	}
}

