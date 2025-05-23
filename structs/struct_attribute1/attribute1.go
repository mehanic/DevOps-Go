package main

import "fmt"

// Define the Emp struct
type Emp struct {
	name   string
	salary int
}

// Constructor function equivalent to __init__ in Python
func NewEmp(name string, salary int) *Emp {
	return &Emp{
		name:   name,
		salary: salary,
	}
}

// Method to display employee details (similar to the display method in Python)
func (e Emp) Display() {
	fmt.Printf("The name is: %s\nThe salary is: %d\n", e.name, e.salary)
}

func main() {
	// Create employee instances using the constructor
	emp1 := NewEmp("Ramu", 56000)
	emp2 := NewEmp("Naren", 90000)

	// Call the display method to print details of emp1
	emp1.Display()
  emp2.Display()
	// Uncomment below to display emp2 details
	// emp2.Display()
}

// The name is: Ramu
// The salary is: 56000
// The name is: Naren
// The salary is: 90000
