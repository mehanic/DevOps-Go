package maps

import "fmt"

func IterateMaps() {

	// We can use range to iterate over maps.

	salaries := map[string]int{"Berkay": 10000, "John": 185000, "Kevin": 7200, "Son": 22000, "Pethy": 15000}

	for name, salary := range salaries {
		fmt.Printf("Salary of %v is %v $\n", name, salary)
	}

}
