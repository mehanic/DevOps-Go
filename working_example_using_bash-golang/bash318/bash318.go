package main

import (
	"fmt"
	"io/ioutil"
//	"os"
	"time"
)

func main() {
	// Set timezone to America/Los_Angeles
	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println("Error loading timezone:", err)
		return
	}
	fmt.Println("Your Timezone is = America/Los_Angeles")
	fmt.Println("Current time in America/Los_Angeles:", time.Now().In(loc).Format(time.RFC1123))

	// Set timezone to Asia/Tokyo
	loc, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("Error loading timezone:", err)
		return
	}
	fmt.Println("Your Timezone is = Asia/Tokyo")
	fmt.Println("Current time in Asia/Tokyo:", time.Now().In(loc).Format(time.RFC1123))

	// Unset TZ equivalent
	// Not applicable in Go, just ignoring

	// Read the system timezone from /etc/timezone
	timezoneFile := "/etc/timezone"
	data, err := ioutil.ReadFile(timezoneFile)
	if err != nil {
		fmt.Println("Error reading /etc/timezone:", err)
		return
	}
	fmt.Printf("Your Timezone is = %s", data)

	// Get the current time in the system timezone
	fmt.Println("\nCurrent time in the system timezone:")
	fmt.Println(time.Now().Format(time.RFC1123))
}

