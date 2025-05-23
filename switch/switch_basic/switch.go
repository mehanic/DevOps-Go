package main

import "fmt"

func main() {
	money := 50000
	switch money {
	case 0:
		fmt.Println("Not lucky yet")

	case 1000, 2000, 5000:
		fmt.Println("Not bad")

	case 10000, 20000:
		fmt.Println("Lucky")

	case 50000, 100000:
		fmt.Println("Very lucky")

	default:
		fmt.Println("Maybe you are not participating in the lottery")
	}

	main1()
	main2()
	main3()
}

func main1() {
	firstName := ""

	// Switch with a short statement; Condition
	// Short statement: " length := firstName "
	switch length := firstName; length == "Prasetiyo" {
	case true:
		fmt.Println("Prasetiyo")
	case false:
		fmt.Println("You are not recognized by the system")
	}
}

func main2() {
	length := 7
	switch {
	case length < 6:
		fmt.Println("Minimum characters required: 6")
	case length >= 6 && length <= 16:
		fmt.Println("Format is correct")
	case length > 16:
		fmt.Println("Character limit exceeded (max: 16)")
	}
}

func main3() {
	value := 42
	switch value {
	case 100:
		fmt.Println("Value is 100")
		fallthrough
	case 42:
		fmt.Println("Value is 42") // value is 42
		fallthrough
	case 1:
		fmt.Println("Value is 1") // value is 1
		fallthrough
	default:
		fmt.Println("Default case") // default
	}
}


// Very lucky
// You are not recognized by the system
// Format is correct
// Value is 42
// Value is 1
// Default case
