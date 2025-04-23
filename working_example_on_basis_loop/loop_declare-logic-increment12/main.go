package main

import "fmt"

// Iterate through a slice and print each element
func main() {
	mySlice := []string{"Prasetiyo", "Angga", "Diana"}

	for i := 0; i < len(mySlice); i++ {
		fmt.Println(mySlice[i])
	}

	// Output:
	// Prasetiyo
	// Angga
	// Diana

	main1()
	main2()
	main3()
	main4()
	main5()
	fmt.Println("-------------")
	main6()
}

// Iterate through a string and print each character
func main1() {
	myString := "Sirclo"
	for _, eachString := range myString {
		fmt.Println(string(eachString))
	}

	// Output:
	// S
	// i
	// r
	// c
	// l
	// o
}

// Iterate through an array and print each element
func main2() {
	myArray := [5]string{"Prasetiyo", "Diana", "Angga", "Satria", "Ramdani"}
	for _, eachArray := range myArray {
		fmt.Println(eachArray)
	}

	// Output:
	// Prasetiyo
	// Diana
	// Angga
	// Satria
	// Ramdani
}

// Iterate through a slice and print each element
func main3() {
	mySlice := []string{"Prasetiyo", "Diana", "Angga", "Satria", "Ramdani"}
	for _, eachSlice := range mySlice {
		fmt.Println(eachSlice)
	}

	// Output:
	// Prasetiyo
	// Diana
	// Angga
	// Satria
	// Ramdani
}

// Iterate through a map and print each value
func main4() {
	myMap := map[string]string{
		"Name":  "Prasetiyo",
		"Hobby": "Coding",
	}

	for _, eachMap := range myMap {
		fmt.Println(eachMap)
	}

	// Output:
	// Prasetiyo
	// Coding
}

// Find and print the vowels in a string with their indices
func main5() {
	title := "I am learning the fundamentals of Golang"

	for index, letter := range title {
		letterString := string(letter)
		if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
			fmt.Println("index:", index, "letter:", letterString)
		}
	}
}

// Use a switch-case to find vowels in a string and print their indices
func main6() {
	title := "I am learning the fundamentals of Golang"

	for index, letter := range title {
		letterString := string(letter)
		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("index:", index, " letter:", letterString)
		}
	}
}
