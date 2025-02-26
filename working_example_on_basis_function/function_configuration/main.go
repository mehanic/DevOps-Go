package main

import "fmt"

func main() {
    var counters = make(map[string]int, 10)

    modelToMake := map[string]string{
        "prius":    "toyota",
        "chevelle": "chevy",
    }

    // Retrieve a value as follows:
    carMake := modelToMake["chevelle"]
    fmt.Println(carMake) // Prints "chevy"

    // But what happens if the key isn't in the map? In that case, we will receive the zero value of
    // the data type:
    carMake = modelToMake["outback"]
    fmt.Println(carMake) // Prints an empty string since "outback" is not found

    // Check if the key exists in the map and handle accordingly
    if carMake, ok := modelToMake["outback"]; ok {
        fmt.Printf("Car model \"outback\" has make %q\n", carMake)
    } else {
        fmt.Printf("Car model \"outback\" has an unknown make\n")
    }

    // Add a new entry to the map:
    modelToMake["outback"] = "subaru"
    counters["pageHits"] = 10

    // Iterate over the map and print the key-value pairs:
    for key, val := range modelToMake {
        fmt.Printf("Car model %q has make %q\n", key, val)
    }
}
