package main
import "fmt"
func main() {
	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)

	// sends a pointer of minecraft to the discount method
	// same as: (&minecraft).discount(.1)
	minecraft.discount(.1)

	mobydick.print()
	minecraft.print()
	tetris.print()

	// creates a variable that occupies 200mb on memory
	var h huge
	for i := 0; i < 10; i++ {
		h.addr()
	}
}


type book struct {
	title string
	price float64
}

func (b book) print() {
	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
}


type game struct {
	title string
	price float64
}

// be consistent:
// if another method in the same type has a pointer receiver,
// use pointer receivers on the other methods as well.
func (g *game) print() {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}

// + discount gets a copy of `*game`.
// + discount can update the original `game` through the game pointer.
// + it's better to use the same receiver type: `*game` for all methods.
func (g *game) discount(ratio float64) {
	g.price *= (1 - ratio)
}

// PREVIOUS CODE:

// + `g` is a copy: `discount` cannot change the original `g`.
// func (g *game) discount(ratio float64) {
// 	g.price *= (1 - ratio)
// }

type huge struct {
	// about ~200mb
	games [10_000_000]game

	// if this syntax doesn't work on your system, type it as:
	// games [10000000]game
}

// only copies a single pointer.
func (h *huge) addr() {
	fmt.Printf("%p\n", h)
}

// each time it is called,
// the original value will be copied.
// calling addr() x 10 times = ~2 GB.
// func (h huge) addr() {
// 	fmt.Printf("%p\n", &h)
// }

// moby dick      : $10.00
// minecraft      : $18.00
// tetris         : $5.00
// 0xc000180000
// 0xc000180000
// 0xc000180000
// 0xc000180000
// 0xc000180000
// 0xc000180000
// 0xc000180000
// 0xc000180000
// 0xc000180000
// 0xc000180000
