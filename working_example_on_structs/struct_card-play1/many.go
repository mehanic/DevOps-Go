package main

import (
	"fmt"
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

type Card struct {
	Balance    int
	ValidUntil string
	Cardholder string
	CVV        string
	Number     string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Not enough money")
	}
	c.Balance -= amount
	return nil
}

type ApplePay struct {
	Money   int
	AppleID string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("Not enough money")
	}
	a.Money -= amount
	return nil
}

type Payer interface {
	Pay(int) error
}

func Buy(p Payer) {
	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Payment error %T: %v\n\n", p, err)
		return
	}
	fmt.Printf("Thank you for purchasing via %T\n\n", p)
}

func main() {

	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	var MyMoney Payer
	MyMoney = &Card{Balance: 100, Cardholder: "Hoasker"}
	Buy(MyMoney)

	MyMoney = &ApplePay{Money: 9}
	Buy(MyMoney)
}

// Thank you for purchasing via *main.Wallet

// Thank you for purchasing via *main.Card

// Payment error *main.ApplePay: Not enough money
