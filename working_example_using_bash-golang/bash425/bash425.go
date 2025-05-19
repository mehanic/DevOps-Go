package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	url := "http://google.com"
	client := http.Client{
		Timeout: 2 * time.Second, // Set a timeout for the HTTP request
	}

	// Perform a HEAD request to the URL
	resp, err := client.Head(url)
	if err != nil {
		fmt.Println("The network is down or very slow")
		return
	}
	defer resp.Body.Close() // Ensure the response body is closed

	// Check the HTTP status code
	switch resp.StatusCode / 100 { // Divide by 100 to get the category of the status code
	case 2, 3:
		fmt.Println("HTTP connectivity is up")
	case 5:
		fmt.Println("The web proxy won't let us through")
	default:
		fmt.Println("The network is down or very slow")
	}
}

