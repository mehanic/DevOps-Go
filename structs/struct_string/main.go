
package main

import "fmt"

func main() {
	l := list{
		{title: "moby dick", price: 10, released: toTimestamp(118281600)},
		{title: "odyssey", price: 15, released: toTimestamp("733622400")},
		{title: "hobbit", price: 25},
	}

	l.discount(.5)

	// The list is a stringer.
	// The `fmt.Print` function can print the `l`
	// by calling `l`'s `String()` method.
	//
	// Underneath, `fmt.Print` uses a type switch to
	// detect whether a type is a Stringer:
	// https://golang.org/src/fmt/print.go#L627
	fmt.Print(l)

	// The money type is a stringer.
	// You don't need to call the String method when printing a value of it.
	// var pocket money = 10
	// fmt.Println("I have", pocket)
}


|
func (p *pp) handleMethods(argument interface{}) (handled bool) {
	// ...

	// Checks whether a given argument is an error or an fmt.Stringer
	switch v := argument.(type) {
	// ...
	// If the argument is a Stringer, calls its String() method
	case Stringer:
		// ...
		//        pocket.String()
		//              ^
		//              |
		p.fmtString(v.String(), verb)
		return
	}

	// ...
}

type list []*product

func (l list) String() string {
	if len(l) == 0 {
		return "Sorry. We're waiting for delivery 🚚.\n"
	}

	var str strings.Builder
	for _, p := range l {
		// fmt.Printf("* %s\n", p)
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}
	return str.String()
}

func (l list) discount(ratio float64) {
	for _, p := range l {
		p.discount(ratio)
	}
}

type money float64

// String() returns a string representation of money.
// money is an fmt.Stringer.
func (m money) String() string {
	return fmt.Sprintf("$%.2f", m)
}

type product struct {
	title    string
	price    money
	released timestamp
}

func (p *product) String() string {
	return fmt.Sprintf("%s: %s (%s)", p.title, p.price, p.released)
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

// String() returns a string representation of timestamp.
// timestamp is an fmt.Stringer.
func (ts timestamp) String() string {
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

