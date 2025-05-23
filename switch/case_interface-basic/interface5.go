package main

import "fmt"

func main() {
	//var json map[string]interface{}
	var b interface{} = "3.14" // Example with a float-like string
	var a interface{} = true   // Example with a boolean value

	// Type switch to check the type of variable 'a'
	switch a.(type) {
	case int:
		fmt.Println("a is of type int")
	case string:
		fmt.Println("a is of type string")
	case bool:
		fmt.Println("a is of type bool")
	default:
		fmt.Println("a is of an unknown type")
	}

	// Type switch for the variable 'b'
	switch b.(type) {
	case int:
		fmt.Println("b is of type int")
	case string:
		fmt.Println("b is of type string")
	case float64:
		fmt.Println("b is of type float64")
	default:
		fmt.Println("b is of an unknown type")
	}
}

// a is of type bool
// b is of type string
