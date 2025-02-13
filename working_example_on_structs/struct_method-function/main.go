package main

import "fmt"

type book struct {
	title string
	price float64
}

func (b book) print() {
	// b is a copy of the original `book` value here.
	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
}


type game struct {
	title string
	price float64
}

func (g game) print() {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}


func main() {
	mobydick := book{
		title: "moby dick",
		price: 10,
	}

	minecraft := game{
		title: "minecraft",
		price: 20,
	}

	tetris := game{
		title: "tetris",
		price: 5,
	}

	// #3
	mobydick.print()  // sends `mobydick` value to `book.print`
	minecraft.print() // sends `minecraft` value to `game.print`
	tetris.print()    // sends `tetris` value to `game.print`

	// #2
	// mobydick.printBook()
	// minecraft.printGame()

	// #1
	// printBook(mobydick)
	// printGame(minecraft)
}


// moby dick      : $10.00
// minecraft      : $20.00
// tetris         : $5.00
