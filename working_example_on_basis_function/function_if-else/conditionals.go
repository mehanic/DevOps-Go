package main

import (
	"fmt"
)

func Conditionals() {
	var balance float64 = 10000
	var amountToWithdraw float64 = 4000

	if amountToWithdraw > balance {
		fmt.Println("Invalid amount.")
	}

	if amountToWithdraw <= balance {
		fmt.Println("Withdrawal complete.")
		balance = balance - amountToWithdraw
		fmt.Println("Remaining balance: " + fmt.Sprintf("%f", balance))
		fmt.Printf("Remaining balance: %v\n", balance)
	}
}

func main() {
	Conditionals()
}
