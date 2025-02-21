package main

import (
	"fmt"
	"os"
)

// Struct definition
type Person struct {
	Name        string
	RotWeiseRot int
	Country     string
}

// Function definition
func (p Person) Greet() {
	fmt.Printf("Hello Herr Nikolai, my name is %s and I am %d years in Europa. Iam from %s.\n", p.Name, p.RotWeiseRot, p.Country)
}

func main() {

	// Variable declaration and initialization
	var num1 int = 5
	num2 := 10

	// Conditional statement
	if num1 < num2 {
		fmt.Println("num1 is less than num2")
	} else {
		fmt.Println("num1 is greater than or equal to num2")
	}
	// Loop
	for i := 0; i < 5; i++ {
		fmt.Println("Loop iteration:", i)
	}
	// Array
	numbers := [5]int{1, 2, 3, 4, 5}
	// Loop over array
	for _, number := range numbers {
		fmt.Println("Number:", number)
	}
	// Slice
	names := []string{"Mikola", "Petro", "Dmitro"}
	// Loop over slice
	for _, name := range names {
		fmt.Println("Name:", name)
	}
	// Struct instantiation
	person := Person{Name: "Petro", RotWeiseRot: 2, Country: "Ukraine"}
	// Method call
	person.Greet()
	// File handling
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	file.WriteString("Hello, Ухилянти !")
	fmt.Println("Script execution completed.")
}

