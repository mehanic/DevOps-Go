package main

import (
	"fmt"
	"strings"
)

func main() {
	ws := []WorkWith{
		{"Example", 1},
		{"Example 2", 2},
	}

	fmt.Printf("Initial list: %#v\n", ws)

	// first lower case the list
	ws = Map(ws, LowerCaseData)
	fmt.Printf("After LowerCaseData Map: %#v\n", ws)

	// next increment all versions
	ws = Map(ws, IncrementVersion)
	fmt.Printf("After IncrementVersion Map: %#v\n", ws)

	// lastly remove all versions older than 3
	ws = Filter(ws, OldVersion(3))
	fmt.Printf("After OldVersion Filter: %#v\n", ws)
}

// WorkWith is the struct we'll
// be implementing collections for
type WorkWith struct {
	Data    string
	Version int
}

// Filter is a functional filter. It takes a list of
// WorkWith and a WorkWith Function that returns a bool
// for each "true" element we return it to the resultant
// list
func Filter(ws []WorkWith, f func(w WorkWith) bool) []WorkWith {
	// depending on results, smalles size for result
	// is len == 0
	result := make([]WorkWith, 0)
	for _, w := range ws {
		if f(w) {
			result = append(result, w)
		}
	}
	return result
}

// Map is a functional map. It takes a list of
// WorkWith and a WorkWith Function that takes a WorkWith
// and returns a modified WorkWith. The end result is
// a list of modified WorkWiths
func Map(ws []WorkWith, f func(w WorkWith) WorkWith) []WorkWith {
	// the result should always be the same
	// length
	result := make([]WorkWith, len(ws))

	for pos, w := range ws {
		newW := f(w)
		result[pos] = newW
	}
	return result
}

// ------------------
// LowerCaseData does a ToLower to the
// Data string of a WorkWith
func LowerCaseData(w WorkWith) WorkWith {
	w.Data = strings.ToLower(w.Data)
	return w
}

// IncrementVersion increments a WorkWiths
// Version
func IncrementVersion(w WorkWith) WorkWith {
	w.Version++
	return w
}

// OldVersion returns a closures
// that validates the version is greater than
// the specified amount
func OldVersion(v int) func(w WorkWith) bool {
	return func(w WorkWith) bool {
		return w.Version >= v
	}
}

// Initial list: []main.WorkWith{main.WorkWith{Data:"Example", Version:1}, main.WorkWith{Data:"Example 2", Version:2}}
// After LowerCaseData Map: []main.WorkWith{main.WorkWith{Data:"example", Version:1}, main.WorkWith{Data:"example 2", Version:2}}
// After IncrementVersion Map: []main.WorkWith{main.WorkWith{Data:"example", Version:2}, main.WorkWith{Data:"example 2", Version:3}}
// After OldVersion Filter: []main.WorkWith{main.WorkWith{Data:"example 2", Version:3}}
