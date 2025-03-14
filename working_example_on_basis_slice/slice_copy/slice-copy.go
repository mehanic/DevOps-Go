package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	evens := []int{2, 4}
	odds := []int{3, 5, 7}

	s.Show("evens [before]", evens)
	s.Show("odds  [before]", odds)

	N := copy(evens, odds)
	fmt.Printf("%d element(s) are copied.\n", N)

	s.Show("evens [after]", evens)
	s.Show("odds  [after]", odds)
}

func main1() {
	// #1: received the raining probabilities
	data := []float64{10, 25, 30, 50}

	// #2: received new data
	// newData := []float64{80, 90}
	// for i := 0; i < len(newData); i++ {
	// 	data[i] = newData[i]
	// }

	// #3: use copy
	// copy(data, []float64{99, 100})

	// #4: received more data than the original
	// copy(data, []float64{10, 5, 15, 0, 20})

	// #5: returns the # of copied elements
	// n := copy(data, []float64{10, 5, 15, 0, 20})
	// fmt.Printf("%d probabilities copied.\n", n)

	// #6: (sometimes) use append instead of copy
	// data = append(data[:0], []float64{10, 5, 15, 0, 20}...)

	// #7: clone a slice using copy
	// saved := make([]float64, len(data))
	// copy(saved, data)

	// #9: clone a slice using append nil (I prefer this)
	// saved := append([]float64(nil), data...)

	// data[0] = 0 // #8

	// s.Show("Probabilities (saved)", saved) // #7

	// #1: print the probabilities
	s.Show("Probabilities (data)", data)

	fmt.Printf("Is it gonna rain? %.f%% chance.\n",
		(data[0]+data[1]+data[2]+data[3])/
			float64(len(data)))
}
