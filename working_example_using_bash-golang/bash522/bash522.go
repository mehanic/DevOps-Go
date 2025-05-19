package main

import (
	"fmt"
	"strconv"
)

func decToBin(num int) string {
	// Convert decimal to binary
	return strconv.FormatInt(int64(num), 2)
}

func binToDec(binStr string) (int, error) {
	// Convert binary to decimal
	val, err := strconv.ParseInt(binStr, 2, 0)
	if err != nil {
		return 0, err // Return zero and the error if parsing fails
	}
	return int(val), nil // Convert int64 to int before returning
}

func main() {
	var op int

	// Display the menu
	fmt.Println("Choose from the following operations:")
	fmt.Println("[1] Decimal to Binary Conversion")
	fmt.Println("[2] Binary to Decimal Conversion")
	fmt.Print("Your choice: ")
	fmt.Scan(&op)

	switch op {
	case 1:
		var intNum int
		fmt.Print("Enter integer number: ")
		fmt.Scan(&intNum)
		bin := decToBin(intNum)
		fmt.Printf("The binary equivalent of %d is: %s\n", intNum, bin)

	case 2:
		var binStr string
		fmt.Print("Enter binary number: ")
		fmt.Scan(&binStr)
		dec, err := binToDec(binStr)
		if err != nil {
			fmt.Println("Invalid binary number:", err)
			return
		}
		fmt.Printf("The decimal equivalent of %s is: %d\n", binStr, dec)

	default:
		fmt.Println("Wrong Choice!")
	}
}

