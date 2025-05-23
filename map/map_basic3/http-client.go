package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	classOne := [2]string{"Mandy", "Collins"}
	classTwo := [2]string{"Emeli", "Nesland"}
	classThree := [2]string{"Ariana", "Wood"}
	classFour := [2]string{"Norma", "Patina"}

	fmt.Println(classOne)
	fmt.Println(classTwo)
	fmt.Println(classThree)
	fmt.Println(classFour)

	for index, item := range classOne {
		fmt.Println(index, item)
	}

	for index, item := range classThree {
		fmt.Println(index, item)
	}

	for index, item := range classFour {
		fmt.Println(index, item)
	}

	for index, item := range classTwo {
		fmt.Println(index, item)
	}

	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Errorf("error: %v", err)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Current time %s", time)

	// Create HTTP client with timeout
	client := &http.Client{
		//Timeout: 30 * time.Second,
	}

	// Make request
	response, err := client.Get("https://randomuser.me/api?results=50")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Copy data from the response to standard output
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Number of bytes copied to STDOUT:", n)
}

