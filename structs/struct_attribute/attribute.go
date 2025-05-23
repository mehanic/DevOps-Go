package main

import "fmt"

// Defining the Emp struct
type Emp struct {
	name   string
	age    int
	salary int
}

// Variable to keep track of employee count (similar to class-level variable)
var empCount int = 0

// Method to set employee details
func (e *Emp) GetNameAgeSalary(name string, age, salary int) {
	e.name = name
	e.age = age
	e.salary = salary
	e.IncreaseCountForEmp()
}

// Method to increase the employee count
func (e *Emp) IncreaseCountForEmp() {
	empCount++
}

// Method to display employee details
func (e Emp) DisplayDetails() {
	fmt.Printf("The name is: %s\nThe age is: %d\nThe salary is: %d\n", e.name, e.age, e.salary)
}

func main() {
	// Creating two employee instances
	emp1 := Emp{}
	emp2 := Emp{}

	// Setting employee details
	emp1.GetNameAgeSalary("John", 34, 45000)
	emp2.GetNameAgeSalary("Cliton", 25, 54000)

	// Displaying the global employee count
	fmt.Println(empCount)

	// Displaying the employee details
	emp1.DisplayDetails()
	emp2.DisplayDetails()

	// Displaying the global employee count
	fmt.Println("Total Employees:", empCount)
}


// 2
// The name is: John
// The age is: 34
// The salary is: 45000
// The name is: Cliton
// The age is: 25
// The salary is: 54000
// Total Employees: 2


