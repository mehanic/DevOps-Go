package main

import (
	"fmt"
	"strconv"
)

// --------------

type Wallet struct {
	Cash int
}

// Method to implement the Pay method of the Payer interface
func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Not enough cash")
	}
	w.Cash -= amount
	return nil
}

// Method to implement the String method
func (w *Wallet) String() string {
	return "Wallet with " + strconv.Itoa(w.Cash) + " money"
}

// --------------

type Payer interface {
	Pay(int) error
}

// --------------

func Buy(in interface{}) {
	var p Payer
	var ok bool
	// Type assertion to check if the input implements the Payer interface
	if p, ok = in.(Payer); !ok {
		// Print message if the input doesn't implement the Payer interface
		fmt.Printf("%T isn't a means of payment\n\n", in)
		return
	}

	// Attempt to make a payment
	err := p.Pay(10) // Hardcoded payment of 10
	if err != nil {
		// Handle any error in payment
		fmt.Printf("Payment error %T: %v\n\n", p, err)
		return
	}

	// Thank you message after a successful purchase
	fmt.Printf("Thank you for purchasing via %T\n\n", p)
}

// --------------

func main() {
	// Create a Wallet with 100 cash
	myWallet := &Wallet{Cash: 100}
	// Call Buy with myWallet as input
	Buy(myWallet)

	// Call Buy with a non-Payer input (slice of ints)
	Buy([]int{1, 2, 3})

	// Call Buy with a non-Payer input (float64)
	Buy(3.14)
}


// Thank you for purchasing via *main.Wallet

// []int isn't a means of payment

// float64 isn't a means of payment

