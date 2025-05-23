package main

import "fmt"


func main() {
	var result interface{} = random()
	// Type Assertions
	var resultString string = result.(string) //Could you call me Yoso?
	fmt.Println(resultString)
}

func random() interface{} {
	return "Could you call me Yoso?"
}

func main1() {
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
}

func random1() interface{} {
	return "Could you call me Yoso?"
}

