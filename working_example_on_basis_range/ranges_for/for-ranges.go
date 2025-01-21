package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// i := 1
	// fmt.Printf("%q\n", os.Args[1])
	// fmt.Printf("%q\n", os.Args[i])

	// i++
	// fmt.Printf("%q\n", os.Args[2])
	// fmt.Printf("%q\n", os.Args[i])

	// i++
	// fmt.Printf("%q\n", os.Args[3])
	// fmt.Printf("%q\n", os.Args[i])

	// --------------------------------
	// #1st way:
	// --------------------------------

	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Printf("%q\n", os.Args[i])
	// }

	// --------------------------------
	// #2nd way:
	// --------------------------------

	// for i, v := range os.Args {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	fmt.Printf("%q\n", v)
	// }

	// --------------------------------
	// #3rd way (best):
	// --------------------------------

	for _, v := range os.Args[1:] {
		fmt.Printf("%q\n", v)
	}
}

func main1() {
	// this will split the following string by spaces
	// and then it will put it inside words variable
	// as a slice of strings
	words := strings.Fields(
		"lazy cat jumps again and again and again",
	)
	//    1    2    3    4    5     6    7    8

	// --------------------------------
	// #1st way:
	// --------------------------------
	// for j := 0; j < len(words); j++ {
	// 	fmt.Printf("#%-2d: %q\n", j+1, words[j])
	// }

	// --------------------------------
	// #2nd way (best):
	// --------------------------------
	for i, v := range words {
		fmt.Printf("#%-2d: %q\n", i+1, v)
	}

	// --------------------------------
	// #3rd way (reuse mechanism):
	// --------------------------------
	// var (
	// 	i int
	// 	v string
	// )

	// for i, v = range words {
	// 	fmt.Printf("#%-2d: %q\n", i+1, v)
	// }

	// fmt.Printf("Last value of i and v %d %q\n", i, v)
}

func main2() {
	// rand.Seed(10)
	// rand.Seed(100)

	// t := time.Now()
	// rand.Seed(t.UnixNano())

	// ^-- same:

	rand.Seed(time.Now().UnixNano())

	guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

//--------------

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?
`
)

func main3() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 1 {
		// fmt.Println("Pick a number.")
		fmt.Printf(usage, maxTurns)
		return
	}
}

// ----
const corpus1 = "lazy cat jumps again and again and again"

func main4() {
	words := strings.Fields(corpus1)
	query := os.Args[1:]

	// after the inner loop breaks
	// this parent loop will look for the next queried word
	for _, q := range query {

		// "break" will terminate this loop
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				// find the first word then break
				// the nested loop
				break
			}
		}

	}
}

// --
const corpus2 = "lazy cat jumps again and again and again"

func main5() {
	words := strings.Fields(corpus2)
	query := os.Args[1:]

	// labels and other names do not share the same scope
	// var queries string
	// _ = queries

	// this label labels the parent loop below
	// label's scope is the whole function body
	// not only it's block
queries:
	for _, q := range query {
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				// find the first word then quit
				break queries
			}
		}
	}
}

// ----
const corpus3 = "lazy cat jumps again and again and again"

func main6() {
	words := strings.Fields(corpus3)
	query := os.Args[1:]

queries:
	for _, q := range query {
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				// skip duplicate words
				continue queries
			}
		}
	}
}

//----

const corpus4 = "lazy cat jumps again and again and again"

func main7() {
	words := strings.Fields(corpus4)
	query := os.Args[1:]

	// labels and other names do not share the same scope
	// this will work even though `queries` label exists
	//
	// var queries string
	// _ = queries

	// this label labels the parent loop below.
	// label's scope is the whole func main()
queries:
	for _, q := range query {

	search:
		for i, w := range words {
			switch q {
			case "and", "or", "the":
				break search
			}

			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				// find the first word then quit
				continue queries
			}
		}
	}
}

// ---
func main8() {
	// cannot step over any variable declarations
	// ERROR: "i" variable is declared after the jump
	//
	// goto loop

	var i int

loop:
	if i < 3 {
		fmt.Println("looping")
		i++
		goto loop
	}
	fmt.Println("done")
}
