package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Define a struct that matches the JSON structure
type Number struct {
	Number int `json:"number"`
}

type Fruit struct {
	Fruit string `json:"fruit"`
}

type Data struct {
	ArrayOfNums   []Number `json:"arrayOfNums"`
	ArrayOfFruits []Fruit  `json:"arrayOfFruits"`
}

func main() {
	// JSON string
	jsonString := `{"arrayOfNums":[{"number":0},{"number":1},{"number":2}],"arrayOfFruits":[{"fruit":"apple"},{"fruit":"banana"},{"fruit":"pear"}]}`

	// Parse the JSON string into the 'Data' struct
	var jsonObj Data
	err := json.Unmarshal([]byte(jsonString), &jsonObj)
	if err != nil {
		log.Fatal(err)
	}

	// Access the elements similarly to Python's get() method
	fmt.Println(jsonObj.ArrayOfNums)
	fmt.Println(jsonObj.ArrayOfNums[1])
	sum := jsonObj.ArrayOfNums[1].Number + jsonObj.ArrayOfNums[2].Number
	fmt.Println(sum)
	fmt.Println(jsonObj.ArrayOfFruits[2].Fruit)
}

