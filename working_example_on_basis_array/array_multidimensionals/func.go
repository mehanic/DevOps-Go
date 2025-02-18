package arrays

import "fmt"

func CreateArray() {

	var numbers [5]int // An integer array contains 5 numbers
	// or numbers:= [5]int{10,20,30,40,50}

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println(numbers)

	// Now write each number while looping

	for ind := 0; ind < len(numbers); ind++ {
		fmt.Printf("Index %v is %v\n", ind, numbers[ind])
	}

}
