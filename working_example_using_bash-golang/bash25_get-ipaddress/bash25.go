package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getCountry(ipAddress string) (string, error) {
	// Prepare the API URL (new URL for ip-api)
	url := "http://ip-api.com/json/" + ipAddress

	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the JSON response
	var responseJson map[string]interface{}
	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}

	// Extract the countryCode from the response
	countryCode, ok := responseJson["countryCode"].(string)
	if !ok {
		return "", fmt.Errorf("could not find countryCode in response")
	}

	return countryCode, nil
}

func main() {
	ipAddress := "50.78.253.58"
	country, err := getCountry(ipAddress)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Country Code:", country)
	}
}

