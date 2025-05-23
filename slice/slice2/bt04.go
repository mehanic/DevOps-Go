package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Exercise 4: An employee in the company has the following attributes: Name, Salary Coefficient, Allowance")
	fmt.Println("Create an array of employees (any number) and perform the following functions:")
	fmt.Println("- Sort employee names in ascending alphabetical order")
	fmt.Println("- Sort employees by descending salary (Salary = Salary Coefficient * 1,500,000 + Allowance)")
	fmt.Println("- Get the list of employees with the second highest salary in the employee array")
	fmt.Println("-----------------------------------------------------------------------------------------------")
	sortSliceWithFunc()
}

func sortSliceWithFunc() {
	people := []struct {
		Name      string
		Ratio     float64
		Allowance float64
	}{
		{"Gopher", 2.5, 100000},
		{"Alice", 2.6, 100000},
		{"Vera", 2.6, 110000},
		{"Bob", 2.7, 110000},
	}
	fmt.Println("Initial list of employees: ", people)
	fmt.Println()

	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("List of employees sorted by name:", people)
	fmt.Println()

	// NOTE: The salary formula here should be Ratio * 1,500,000 + Allowance
	sort.Slice(people, func(i, j int) bool {
		salaryI := people[i].Ratio*1500000 + people[i].Allowance
		salaryJ := people[j].Ratio*1500000 + people[j].Allowance
		return salaryI > salaryJ
	})
	fmt.Println("List of employees sorted by descending salary:", people)
	fmt.Println()

	fmt.Println("Employee(s) with the second highest salary:", people[1])
	fmt.Println()
}
