package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: 2 * time.Second, // Set the timeout to 2 seconds
	}

	resp, err := client.Head("http://google.com") // Make a HEAD request to google.com
	if err != nil {
		fmt.Println("The network is down or very slow")
		return
	}
	defer resp.Body.Close()

	// Get the first digit of the status code
	statusCode := resp.StatusCode / 100

	// Handle based on the first digit of the status code
	switch statusCode {
	case 2, 3:
		fmt.Println("HTTP connectivity is up")
	case 5:
		fmt.Println("The web proxy won't let us through")
	default:
		fmt.Println("The network is down or very slow")
	}
}

