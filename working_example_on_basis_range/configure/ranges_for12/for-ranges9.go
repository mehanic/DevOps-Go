package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// Helper function to get random range
func randomRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed random number generator

	domain := []string{"gmail", "yahoo", "hotmail", "rediffmail"}
	courses := []string{
		"B Tech(CSE)", "B Tech(IT)", "B Tech(ME)", "B Tech(CIVIL)",
		"B Tech(EEE)", "B Tech(ECE)", "B Pharm", "BBA", "BBA(HM)", "BCA",
		"B.Sc(IT)", "B.Sc(Agriculture)", "B.Sc(Forestry)", "B.Com", "BSC(PCM)",
		"BSC(CBZ)", "BSC(Horti)", "MCA", "M Tech(CSE)", "M Tech(ME)", "M Tech(EEE)",
		"M Tech(ECE)", "D.Pharam", "Diploma(CS)", "Diploma(EE)", "Diploma(EC)", "Diploma(ME)",
	}
	queries := []string{
		"How to apply?", "What is the last date to apply?", "Please call & explain the things.",
	}

	// Open the CSV file
	file, err := os.Open("names.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Read CSV file
	reader := csv.NewReader(file)
	names, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	for i := 0; i < 100; i++ {
		a := strconv.Itoa(randomRange(7, 10))
		b := strconv.Itoa(randomRange(7, 10))
		c := strconv.Itoa(randomRange(70, 100))
		d := strconv.Itoa(randomRange(111111, 999999))

		name := names[i][0]
		contact := a + b + c + d
		email := strings.ToLower(name) + strconv.Itoa(randomRange(1, 9999)) + "@" + domain[rand.Intn(len(domain))] + ".com"
		course := courses[rand.Intn(len(courses))]
		query := queries[rand.Intn(len(queries))]

		data := url.Values{}
		data.Set("name", name)
		data.Set("contact", contact)
		data.Set("email", email)
		data.Set("courses", course)
		data.Set("query", query)

		urlStr := "http://www.dbgidoon.ac.in/registration-form/submit1.php"
		req, err := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
		if err != nil {
			log.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Request failed: %v", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Printf("%d : Success! Status code: %d\n", i, resp.StatusCode)
		} else if resp.StatusCode == http.StatusNotFound {
			fmt.Printf("Error 404: %s\n", urlStr)
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Printf("%d : Received response: %s\n", i, string(body))
		}
	}
}

