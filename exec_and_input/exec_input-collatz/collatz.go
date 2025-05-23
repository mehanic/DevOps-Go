package main

import (
	"fmt"
	"os"
	"strconv"
)

func collatz(number int) int {
	if number%2 == 0 {
		return number / 2
	}
	return 3*number + 1
}

func main() {
	fmt.Print("Enter integer number:    ")
	var input string
	fmt.Scanln(&input)

	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("The integer number is only accepted.")
		os.Exit(1)
	}

	if number == float64(int(number)) {
		fmt.Println(int(number))
		for number != 1 {
			number = float64(collatz(int(number)))
			fmt.Println(int(number))
		}
	} else {
		fmt.Println("The integer number is only accepted.")
		os.Exit(1)
	}
}


// Enter integer number:    67
// 67
// 202
// 101
// 304
// 152
// 76
// 38
// 19
// 58
// 29
// 88
// 44
// 22
// 11
// 34
// 17
// 52
// 26
// 13
// 40
// 20
// 10
// 5
// 16
// 8
// 4
// 2
// 1

