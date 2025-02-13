
package main


		import (
			
			"fmt"
		)

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
		}
		
		type list []printer
		
		func (l list) print() {
			if len(l) == 0 {
				fmt.Println("Sorry. We're waiting for delivery 🚚.")
				return
			}
		
			for _, it := range l {
				// fmt.Printf("(%-10T) --> ", it)
		
				it.print()
		
				// you cannot access to the discount method of the game type.
				// `it` is a printer not a game.
				// it.discount(.5)
			}
		}
		
		// PREVIOUS CODE:
		
		// type list []*game
		
		// func (l list) print() {
		// 	if len(l) == 0 {
		// 		fmt.Println("Sorry. Our store is closed. We're waiting for the delivery 🚚.")
		// 		return
		// 	}
		
		// 	for _, it := range l {
		// 		it.print()
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
		
		// if you remove this method,
		// you can no longer add it to the `store` in `main()`.
		func (p puzzle) print() {
			fmt.Printf("%-15s: %s\n", p.title, p.price.string())
		}

		
func main() {
	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
		rubik     = puzzle{title: "rubik's cube", price: 5}
	)

	// thanks to the printer interface we can add different types of values
	// to the list.
	//
	// only rule: they need to implement the `printer` interface.
	// to do that: each type needs to have a print method.

	var store list
	store = append(store, &minecraft, &tetris, mobydick, rubik)
	store.print()

	// interface values are comparable
	fmt.Println(store[0] == &minecraft)
	fmt.Println(store[3] == rubik)
}


// minecraft      : $20.00
// tetris         : $5.00
// moby dick      : $10.00
// rubik's cube   : $5.00
// true
// true
