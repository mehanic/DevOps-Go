
package main

import (
	"fmt"
	"strconv"
	"time"
)
func main() {
	store := list{
		&book{product{"moby dick", 10}, 118281600},
		&book{product{"odyssey", 15}, "733622400"},
		&book{product{"hobbit", 25}, nil},
		&puzzle{product{"rubik's cube", 5}},
		&game{product{"minecraft", 20}},
		&game{product{"tetris", 5}},
		&toy{product{"yoda", 150}},
	}

	store.discount(.5)
	store.print()

	// t := &toy{product{"yoda", 150}}
	// fmt.Printf("%#v\n", t)

	// b := &book{product{"moby dick", 10}, 118281600}
	// fmt.Printf("%#v\n", b)
}

type item interface {
	print()
	discount(ratio float64)
}

type list []item

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
	for _, it := range l {
		it.discount(ratio)
	}
}




type book struct {
	// embed the product type into the book type.
	// all the fields and methods of the product will be
	// available in this book type.
	product
	published interface{}
}

// book type's print method takes priority.
//
// + when you call it on a book value, the following method will
//   be executed.
//
// + if it wasn't here, the product type's print method would
//   have been executed.
func (b *book) print() {
	// the book can also call the embedded product's print method
	// if it wants to, as in here:
	b.product.print()

	p := format(b.published)
	fmt.Printf("\t - (%v)\n", p)
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
	product
}

type money float64

func (m money) string() string {
	return fmt.Sprintf("$%.2f", m)
}



type product struct {
	title string
	price money
}

func (p *product) print() {
	fmt.Printf("%-15s: %s\n", p.title, p.price.string())
}

func (p *product) discount(ratio float64) {
	p.price *= money(1 - ratio)
}

type puzzle struct {
	product
}

type toy struct {
	product
}


// moby dick      : $5.00
// 	 - (1973/10)
// odyssey        : $7.50
// 	 - (1993/04)
// hobbit         : $12.50
// 	 - (unknown)
// rubik's cube   : $2.50
// minecraft      : $10.00
// tetris         : $2.50
// yoda           : $75.00

