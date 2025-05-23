package main

import "fmt"

func WithdrawExample() {
	var balance float64 = 10000
	var amountToWithdraw float64 = 10000

	if amountToWithdraw < balance {
		fmt.Println("Withdrawal complete.")
		balance = balance - amountToWithdraw
		fmt.Println("Remaining balance: " + fmt.Sprintf("%f", balance))
		fmt.Printf("Remaining balance: %v\n", balance)
	} else if amountToWithdraw == balance {
		fmt.Println("You are withdrawing all your money.")
		balance = 0
		fmt.Printf("Remaining balance: %v\n", balance)
	} else {
		fmt.Println("Insufficient funds.")
	}
}

func main() {
	WithdrawExample()
}
