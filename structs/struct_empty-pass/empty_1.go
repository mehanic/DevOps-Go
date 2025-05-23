package main

import (
	"fmt"
	"strconv"
)

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Not enough money")
	}
	w.Cash -= amount
	return nil
}

func (w *Wallet) String() string {
	return "Wallet with " + strconv.Itoa(w.Cash) + " money"
}

func main() {
	myWallet := &Wallet{Cash: 100}
	fmt.Printf("Raw payment: %#v\n", myWallet)
	fmt.Printf("Payment method: %s\n", myWallet)
}

// Raw payment: &main.Wallet{Cash:100}
// Payment method: Wallet with 100 money
