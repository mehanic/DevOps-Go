package main

import (
	"fmt"
)

func main() {
	classOne := [2]string{"John", "Doe"}
	classTwo := [2]string{"Jane", "Doe"}

	fmt.Println(classTwo)

	for index, item := range classOne {
		fmt.Println(index, item)
	}
}

// [Jane Doe]
// 0 John
// 1 Doe
