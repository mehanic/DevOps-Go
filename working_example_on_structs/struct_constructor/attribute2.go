package main

import "fmt"

// Define the Person struct
type Person struct {
	name string
	age  int
}

// Constructor function for Person
func NewPerson(name string, age int) *Person {
	fmt.Println("An object has been created")
	return &Person{
		name: name,
		age:  age,
	}
}

// Method to simulate object deletion (cleanup)
func (p *Person) Cleanup() {
	fmt.Println("Object has been deleted")
}

func main() {
	// Create an instance of Person
	per1 := NewPerson("Jhon", 26)
	
	// Simulate cleanup
	defer per1.Cleanup()
	
	// Do other operations here if needed
}

// An object has been created
// Object has been deleted
