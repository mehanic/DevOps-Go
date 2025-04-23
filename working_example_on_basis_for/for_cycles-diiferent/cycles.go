package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("-----------------------------")
	for i := 1; i <= 100; i++ {
		if i%7 != 0 {
			continue
		}

		fmt.Println(i)
	}
	fmt.Println("-----------------------------")

	for i := 1; i <= 100; i++ {
		if i == 10 {
			break
		}

		fmt.Println(i)
	}
	fmt.Println("After loop")
	fmt.Println("-----------------------------")

	nums := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for index, element := range nums {
		fmt.Printf("Index: %d Element: %d\n", index, element)
	}

	for _, element := range nums {
		fmt.Printf("Element: %d\n", element)
	}
	fmt.Println("-----------------------------")

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for _, row := range matrix {
		for _, column := range row {
			fmt.Printf("%d ", column)
		}
		fmt.Println()
	}
}
