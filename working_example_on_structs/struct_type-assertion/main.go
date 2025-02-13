
package main

import "fmt"
func main() {
	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
		rubik     = puzzle{title: "rubik's cube", price: 5}
		yoda      = toy{title: "yoda", price: 150}
	)

	var store list
	store = append(store, &minecraft, &tetris, mobydick, rubik, &yoda)

	// #2
	store.discount(.5)
	store.print()

	// #1
	// var p printer
	// p = &tetris
	// tetris.discount(.5)
	// p.print()
}


type book struct {
	title string
	price money
}

func (b book) print() {
	fmt.Printf("%-15s: %s\n", b.title, b.price.string())
}

type game struct {
	title string
	price money
}

func (g *game) print() {
	fmt.Printf("%-15s: %s\n", g.title, g.price.string())
}

func (g *game) discount(ratio float64) {
	g.price *= money(1 - ratio)
}


type printer interface {
	print()

	// use type assertion when you cannot change the interface.
	// discount(ratio float64)
}

type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery 🚚.")
		return
	}

	for _, it := range l {
		it.print()
	}
}

// type assertion can extract the wrapped value.
// or: it can put the value into another interface.
func (l list) discount(ratio float64) {
	// you can declare an interface in a function/method as well.
	// interface is just a type.
	type discounter interface {
		discount(float64)
	}

	for _, it := range l {
		// you can assert to an interface.
		// and extract another interface.
		if it, ok := it.(discounter); ok {
			it.discount(ratio)
		}
	}
}

// ----

// func (l list) discount(ratio float64) {
// 	for _, it := range l {
// 		// you can inline-assert interfaces
// 		// interface is just a type.
// 		g, ok := it.(interface{ discount(float64) })
// 		if !ok {
// 			continue
// 		}

// 		g.discount(ratio)
// 	}
// }

// ----

// // type assertion can pull out the real value behind an interface value.
// // it can also check whether the value convertable to a given type.
// func (l list) discount(ratio float64) {
// 	for _, it := range l {
// 		g, ok := it.(*game)
// 		if !ok {
// 			continue
// 		}

// 		g.discount(ratio)
// 	}
// }

type money float64

func (m money) string() string {
	return fmt.Sprintf("$%.2f", m)
}

type puzzle struct {
	title string
	price money
}

func (p puzzle) print() {
	fmt.Printf("%-15s: %s\n", p.title, p.price.string())
}

type toy struct {
	title string
	price money
}

func (t *toy) print() {
	fmt.Printf("%-15s: %s\n", t.title, t.price.string())
}

func (t *toy) discount(ratio float64) {
	t.price *= money(1 - ratio)
}


// minecraft      : $10.00
// tetris         : $2.50
// moby dick      : $10.00
// rubik's cube   : $5.00
// yoda           : $75.00

