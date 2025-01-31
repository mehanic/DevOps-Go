package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

//This Go program reads names from a CSV file (names.csv), generates random emails and phone numbers, and makes HTTP POST requests to http://www.jbitdoon.com/index.php with the generated data.



func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	domain := []string{"gmail", "yahoo", "hotmail", "rediffmail"}

	// Open the CSV file
	file, err := os.Open("names.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read names from the CSV file
	var names []string
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Extract names from CSV rows
	for _, line := range lines {
		names = append(names, line[0]) // assuming names are in the first column
	}

	// Generate random values and make HTTP requests
	for i := 0; i < 100; i++ {
		a := strconv.Itoa(rand.Intn(3) + 7) // Random number from 7 to 9
		b := strconv.Itoa(rand.Intn(3) + 7) // Random number from 7 to 9
		c := strconv.Itoa(rand.Intn(30) + 70) // Random number from 70 to 99
		d := strconv.Itoa(rand.Intn(888888) + 111111) // Random number from 111111 to 999999

		// Assuming names list has enough elements
		if i >= len(names) {
			break
		}
		name := names[i]
		email := name + strconv.Itoa(rand.Intn(9999)) + "@" + domain[rand.Intn(len(domain))] + ".com"
		phoneno := a + b + c + d

		values := url.Values{
			"name":    {name},
			"email":   {email},
			"phoneno": {phoneno},
		}

		url := "http://www.jbitdoon.com/index.php"
		resp, err := http.PostForm(url, values)
		if err != nil {
			fmt.Println("Request failed:", err)
			continue
		}

		fmt.Printf("%d : %d\n", i, resp.StatusCode)

		resp.Body.Close() // Close the response body
	}
}

