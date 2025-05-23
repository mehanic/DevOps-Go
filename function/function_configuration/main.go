package main

import "fmt"

func main() {
	counters := make(map[string]int, 10)

	modelToMake := map[string]string{
		"prius":    "toyota",
		"chevelle": "chevy",
	}

	carMake := modelToMake["chevelle"]
	fmt.Println(carMake) // Prints "chevy"

	carMake = modelToMake["outback"]
	fmt.Println(carMake) // Prints an empty string since "outback" is not found

	if carMake, ok := modelToMake["outback"]; ok {
		fmt.Printf("Car model \"outback\" has make %q\n", carMake)
	} else {
		fmt.Printf("Car model \"outback\" has an unknown make\n")
	}

	modelToMake["outback"] = "subaru"
	counters["pageHits"] = 10

	for key, val := range modelToMake {
		fmt.Printf("Car model %q has make %q\n", key, val)
	}
}
