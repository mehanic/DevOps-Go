package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func error_handler(err error) float64 {
	if err != nil {
		return float64(-1)
	}
	return float64(0)
}

func add_two_num(a, b float64) float64 {
	return a + b
}

func sub_two_num(a, b float64) float64 {
	return a - b
}

func mul_two_num(a, b float64) float64 {
	return a * b
}

func div_two_num(a, b float64) float64 {
	return a / b
}

func main() {
	fmt.Println("Calculator app")
	for {
		fmt.Println("Enter 1 for addition:")
		fmt.Println("Enter 2 for substarction:")
		fmt.Println("Enter 3 for multiplication:")
		fmt.Println("Enter 4 for divison:")
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter the first number:")
		scanner.Scan()
		a, err := strconv.ParseFloat(scanner.Text(), 64)
		error_handler(err)
		fmt.Println("Enter the first number:")
		scanner.Scan()
		b, err := strconv.ParseFloat(scanner.Text(), 64)
		error_handler(err)
		fmt.Println("Enter the choice:")
		scanner.Scan()
		choice, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		switch choice {
		case 1:
			num := add_two_num(a, b)
			fmt.Println(num)
		case 2:
			num := add_two_num(a, b)
			fmt.Println(num)
		case 3:
			num := add_two_num(a, b)
			fmt.Println(num)
		}
	}
}
