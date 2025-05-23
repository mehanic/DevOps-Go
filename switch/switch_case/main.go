package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in GoLang")

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	diceNumber := rng.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can move 1 spot")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spots")
		fallthrough
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can open or move 6 spots and you get to roll dice again!")
	default:
		fmt.Println("What was that!")
	}
}

// Switch and case in GoLang
// Value of dice is  6
// You can open or move 6 spots and you get to roll dice again!

