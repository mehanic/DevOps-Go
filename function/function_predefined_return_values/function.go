package main

import (
	"errors"
	"fmt"
)

func main() {
	luas, keliling := calculate(10, 4)
	fmt.Println("Luas : ")
	fmt.Println(luas)
	fmt.Println("Keliling : ")
	fmt.Println(keliling)

}

func calculate(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)

	return
}




func main1() {
	scores := []int{10, 5, 9, 8, 7}

	total := sum(scores)
	fmt.Println("Total score in the array is: ")
	fmt.Println(total)
	
	addition, substraction, multiple, division := calculation1(3, 2)
	fmt.Println("Total addition: ")
	fmt.Println(addition)
	fmt.Println("Total subtraction: ")
	fmt.Println(substraction)
	fmt.Println("Total multiplication: ")
	fmt.Println(multiple)
	fmt.Println("Total division: ")
	fmt.Println(division)
				   
}

func sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total = total + number
	}

	return total
}

func calculation1(number1 int, number2 int) (addition int, substraction int, multiple int, division int) {
	addition = number1 + number2
	substraction = number1 - number2
	multiple = number1 * number2
	division = number1 / number2

	return
}

func main2() {
	result, err := calculate2(5, 4, "*")
	if err != nil {
		fmt.Println("Terjadi Kesalahan")
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}

func calculate2(number int, numberTwo int, operation string) (int, error) {
	var result int
	var errorResult error

	switch operation {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "*":
		result = number * numberTwo
	case "/":
		result = number / numberTwo
	default:
		errorResult = errors.New("unknown operation")
	}

	return result, errorResult
}

