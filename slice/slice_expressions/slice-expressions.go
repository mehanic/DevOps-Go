package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	msg := []byte{'h', 'e', 'l', 'l', 'o'}
	s.Show("msg", msg)

	s.Show("msg[0:1]", msg[0:1])
	s.Show("msg[0:2]", msg[0:2])
	s.Show("msg[0:3]", msg[0:3])
	s.Show("msg[0:4]", msg[0:4])
	s.Show("msg[0:5]", msg[0:5])

	// default indexes
	s.Show("msg[0:]", msg[0:])
	s.Show("msg[:5]", msg[:5])
	s.Show("msg[:]", msg[:])

	// error: beyond
	// s.Show("msg", msg)[:6]

	s.Show("msg[1:4]", msg[1:4])

	s.Show("msg[1:5]", msg[1:5])
	s.Show("msg[1:]", msg[1:])

	s.Show("append(msg)", append(msg[:4], '!'))
}

func main1() {
	// think of this as search results of a search engine.
	// it could have been fetched from a database
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	s.MaxPerLine = 4
	s.Show("All items", items)

	top3 := items[:3]
	s.Show("Top 3 items", top3)

	l := len(items)

	// you can use variables in a slice expression
	last4 := items[l-4:]
	s.Show("Last 4 items", last4)

	// reslicing: slicing another sliced slice
	mid := last4[1:3]
	s.Show("Last4[1:3]", mid)

	// the same elements can be in different indexes
	// fmt.Println(items[9], last4[0])

	// slicing returns a slice with the same type of the sliced slice
	fmt.Printf("slicing : %T %[1]q\n", items[2:3])

	// indexing returns a single element with the type of the indexed slice's element type
	fmt.Printf("indexing: %T %[1]q\n", items[2])
}
