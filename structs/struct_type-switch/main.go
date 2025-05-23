
package main
import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	store := list{
		book{title: "moby dick", price: 10, published: 118281600},
		book{title: "odyssey", price: 15, published: "733622400"},
		book{title: "hobbit", price: 25},
		puzzle{title: "rubik's cube", price: 5},
		&game{title: "minecraft", price: 20},
		&game{title: "tetris", price: 5},
		&toy{title: "yoda", price: 150},
	}

	store.discount(.5)
	store.print()
}



type book struct {
	title     string
	price     money
	published interface{}
}

func (b book) print() {
	p := format(b.published)
	fmt.Printf("%-15s: %s - (%v)\n", b.title, b.price.string(), p)
}

func format(v interface{}) string {
	var t int

	switch v := v.(type) {
	case int:
		// book{title: "moby dick", price: 10, published: 118281600},
		t = v
	case string:
		// book{title: "odyssey", price: 15, published: "733622400"},
		t, _ = strconv.Atoi(v)
	default:
		// book{title: "hobbit", price: 25},
		return "unknown"
	}

	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "2006/01"

	u := time.Unix(int64(t), 0)
	return u.Format(layout)
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
		it.print()
	}
}

func (l list) discount(ratio float64) {
	type discounter interface {
		discount(float64)
	}

	for _, it := range l {
		if it, ok := it.(discounter); ok {
			it.discount(ratio)
		}
	}
}

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


// moby dick      : $10.00 - (1973/10)
// odyssey        : $15.00 - (1993/04)
// hobbit         : $25.00 - (unknown)
// rubik's cube   : $5.00
// minecraft      : $10.00
// tetris         : $2.50
// yoda           : $75.00

