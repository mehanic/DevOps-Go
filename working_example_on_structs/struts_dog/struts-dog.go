package main

import (
	"fmt"
	"strings"
)

// Define the Dog struct
type Dog struct {
	name string
	age  int
}

// Constructor to initialize the Dog object
func NewDog(name string, age int) *Dog {
	return &Dog{
		name: name,
		age:  age,
	}
}

// Method to simulate the dog sitting
func (d *Dog) Sit() {
	fmt.Println(strings.Title(d.name) + " is now sitting.")
}

// Method to simulate the dog rolling over
func (d *Dog) RollOver() {
	fmt.Println(strings.Title(d.name) + " rolled over!")
}

func main() {
	// Create instances of Dog
	myDog := NewDog("willie", 6)
	yourDog := NewDog("lucy", 3)

	// Display info for myDog
	fmt.Println("My dog's name is " + strings.Title(myDog.name) + ".")
	fmt.Println("My dog is", myDog.age, "years old.")
	myDog.Sit()
	myDog.RollOver()

	// Display info for yourDog
	fmt.Println("\nYour dog's name is " + strings.Title(yourDog.name) + ".")
	fmt.Println("Your dog is", yourDog.age, "years old.")
	yourDog.Sit()
}

// My dog's name is Willie.
// My dog is 6 years old.
// Willie is now sitting.
// Willie rolled over!

// Your dog's name is Lucy.
// Your dog is 3 years old.
// Lucy is now sitting.


