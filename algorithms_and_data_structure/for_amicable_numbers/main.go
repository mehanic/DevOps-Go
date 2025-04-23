package main

import (
	"fmt"
	"math"
)

func main() {
	AmicableNumberCheck()
}

func AmicableNumberCheck() {
	fmt.Println("We will check is numbers you give are amicable or not!")

	var number1, number2 int

	fmt.Println("Please enter the first number")

	fmt.Scanln(&number1)

	fmt.Println("Please enter the second number")

	fmt.Scanln(&number2)

	number1_dividors_sum := 0
	number2_dividors_sum := 0

	bigger := int(math.Max(float64(number1), float64(number2)))

	for i := 1; i < bigger; i++ {
		if (number1 > i) && (number1%i == 0) {
			number1_dividors_sum = number1_dividors_sum + i
		}

		if (number2 > i) && (number2%i == 0) {
			number2_dividors_sum = number2_dividors_sum + i
		}
	}

	if (number2 == number1_dividors_sum) && (number1 == number2_dividors_sum) {
		fmt.Printf("Hurrrrayy! Numbers %v and %v are amicable", number1, number2)
	} else {
		fmt.Printf("Sorry. Sum of divisors in %v is %v while in %v is %v\n",
			number1, number1_dividors_sum, number2, number2_dividors_sum)
	}
}
