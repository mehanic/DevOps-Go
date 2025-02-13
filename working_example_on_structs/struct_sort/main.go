package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type timestamp struct {
	time.Time
}

func (ts timestamp) String() string {
	if ts.IsZero() {
		return "unknown"
	}
	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "2006/01"
	return ts.Format(layout)
}

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

type money float64

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

// list is a custom type for sorting products.
type list []*product

func (l list) String() string {
	if len(l) == 0 {
		return "Sorry. We're waiting for delivery 🚚.\n"
	}

	var str strings.Builder
	for _, p := range l {
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

// byRelease sorts by product release dates.
type byRelease struct {
	list
}

// Implementing Len() method for byRelease
func (br byRelease) Len() int {
	return len(br.list)
}

// Implementing Swap() method for byRelease
func (br byRelease) Swap(i, j int) {
	br.list[i], br.list[j] = br.list[j], br.list[i]
}

// Less takes priority over the Less method of the list.
func (br byRelease) Less(i, j int) bool {
	return br.list[i].released.Before(br.list[j].released.Time)
}

func byReleaseDate(l list) sort.Interface {
	return &byRelease{l}
}

func main() {
	// Initializing a list of products
	l := list{
		{title: "moby dick", price: 10, released: toTimestamp(118281600)},
		{title: "odyssey", price: 15, released: toTimestamp("733622400")},
		{title: "hobbit", price: 25},
	}

	// Sorting by release date in descending order
	sort.Sort(sort.Reverse(byReleaseDate(l)))

	// Printing sorted list
	fmt.Print(l)
}


// * odyssey: $15.00 (1993/04)
// * moby dick: $10.00 (1973/10)
// * hobbit: $25.00 (unknown)

