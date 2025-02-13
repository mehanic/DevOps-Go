package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

// Define the Point struct
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Function to serialize a custom object
func serializeInstance(obj interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	// Add the class name of the object
	m["__classname__"] = reflect.TypeOf(obj).Elem().Name()

	// Use reflection to get all fields and their values
	val := reflect.ValueOf(obj).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		m[field.Name] = val.Field(i).Interface()
	}
	return m
}

// Function to unserialize an object
func unserializeObject(m map[string]interface{}) interface{} {
	clsname, ok := m["__classname__"].(string)
	if !ok {
		return m
	}

	if clsname == "Point" {
		p := &Point{}
		if x, ok := m["X"].(float64); ok { // float64 is the default type for numbers in JSON
			p.X = int(x)
		}
		if y, ok := m["Y"].(float64); ok {
			p.Y = int(y)
		}
		return p
	}
	return m
}

// Function to read and write JSON data
func rwJSON() {
	data := map[string]interface{}{
		"name":   "ACME",
		"shares": 100,
		"price":  542.23,
	}

	// Serialize to JSON string
	jsonStr, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error serializing JSON: %v", err)
	}

	// Deserialize JSON string back into Go map
	var deserializedData map[string]interface{}
	err = json.Unmarshal(jsonStr, &deserializedData)
	if err != nil {
		log.Fatalf("Error deserializing JSON: %v", err)
	}

	// Writing JSON data to file
	err = ioutil.WriteFile("data.json", jsonStr, 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	// Reading data back from file
	fileData, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatalf("Error reading from file: %v", err)
	}
	var fileContent map[string]interface{}
	err = json.Unmarshal(fileData, &fileContent)
	if err != nil {
		log.Fatalf("Error deserializing from file: %v", err)
	}

	fmt.Println("File content:", fileContent)

	// Using Ordered Map equivalent in Go
	// Go doesn't support ordered maps by default, but we can simulate it using slices or specific libraries if needed
	// In this case, just print out the order from the original deserialized data
	fmt.Println("Original Data:", deserializedData)

	// Serialize and deserialize custom objects (Point)
	p := &Point{X: 2, Y: 3}
	serializedPoint, err := json.Marshal(serializeInstance(p))
	if err != nil {
		log.Fatalf("Error serializing custom object: %v", err)
	}
	fmt.Println("Serialized Point:", string(serializedPoint))

	var unserializedPoint map[string]interface{}
	err = json.Unmarshal(serializedPoint, &unserializedPoint)
	if err != nil {
		log.Fatalf("Error deserializing custom object: %v", err)
	}

	pointObject := unserializeObject(unserializedPoint).(*Point)
	if pointObject == nil {
		log.Fatalf("Failed to unserialize object")
	}
	fmt.Printf("Unserialized Point: X=%d, Y=%d\n", pointObject.X, pointObject.Y)
}

func main() {
	rwJSON()
}


// File content: map[name:ACME price:542.23 shares:100]
// Original Data: map[name:ACME price:542.23 shares:100]
// Serialized Point: {"X":2,"Y":3,"__classname__":"Point"}
// Unserialized Point: X=2, Y=3
