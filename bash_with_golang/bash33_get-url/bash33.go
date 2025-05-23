package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// URL to send the GET request to
	url := "http://github.com/zach-king/"

	// Perform the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		// Handle error if the request fails
		fmt.Println("Error fetching the page:", err)
		return
	}
	defer resp.Body.Close() // Ensure the response body is closed

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Handle error if reading the body fails
		fmt.Println("Error reading the response body:", err)
		return
	}

	// Print the content of the page (similar to page.content in Python)
	fmt.Println(string(body))
}

