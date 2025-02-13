
package main
import (
	"strconv"
	"time"
	"fmt"
)

type list []*product

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery 🚚.")
		return
	}

	for _, p := range l {
		p.print()
	}
}

func (l list) discount(ratio float64) {
	for _, p := range l {
		p.discount(ratio)
	}
}



type money float64

func (m money) string() string {
	return fmt.Sprintf("$%.2f", m)
}

type product struct {
	title    string
	price    money
	released timestamp
}

func (p *product) print() {
	fmt.Printf("%s: %s (%s)\n", p.title, p.price.string(), p.released.string())
}

func (p *product) discount(ratio float64) {
	p.price *= money(1 - ratio)
}



// timestamp stores, formats and automatically prints a timestamp.
type timestamp struct {
	// timestamp anonymously embeds a time.
	// no need to convert a time value to a timestamp value to use the methods of the time type.
	time.Time
}

func (ts timestamp) string() string {
	if ts.IsZero() { // same as: ts.Time.IsZero()
		return "unknown"
	}

	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "2006/01"
	return ts.Format(layout) // same as: ts.Time.Format(layout)
}

// toTimestamp returns a timestamp value depending on the type of `v`.
func toTimestamp(v interface{}) (ts timestamp) {
	var t int

	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	}

	ts.Time = time.Unix(int64(t), 0)
	return ts
}


func main() {
	l := list{
		{title: "moby dick", price: 10, released: toTimestamp(118281600)},
		{title: "odyssey", price: 15, released: toTimestamp("733622400")},
		{title: "hobbit", price: 25},
	}

	l.discount(.5)
	l.print()
}

// moby dick: $5.00 (1973/10)
// odyssey: $7.50 (1993/04)
// hobbit: $12.50 (unknown)
