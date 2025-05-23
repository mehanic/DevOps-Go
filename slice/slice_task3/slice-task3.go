package main

import (
	"fmt"
	"strings"

	s "github.com/inancgumus/prettyslice"
)

// observe the cap growth
func main() {
	var (
		nums   []int
		oldCap float64
	)

	// loop 10 million times
	for len(nums) < 1e7 {
		// get the capacity
		c := float64(cap(nums))

		// only print when the capacity changes
		if c == 0 || c != oldCap {
			// print also the growth ratio: c/oldCap
			fmt.Printf("len:%-15d cap:%-15g growth:%-15.2f\n",
				len(nums), c, c/oldCap)
		}

		// keep track of the previous capacity
		oldCap = c

		// append an arbitrary element to the slice
		nums = append(nums, 1)
	}
}

//correct the lyric

func main1() {
	// --- Correct Lyric ---
	// yesterday all my troubles seemed so far away
	// now it looks as though they are here to stay
	// oh I believe in yesterday
	lyric := strings.Fields(`all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay`)

	// ------------------------------------------------------------------------
	// #1: Prepend "yesterday" to `lyric`
	// ------------------------------------------------------------------------

	//
	// --- BEFORE ---
	// all my troubles seemed so far away oh I believe in yesterday
	//
	// --- AFTER ---
	// yesterday all my troubles seemed so far away oh I believe in yesterday
	//
	// (warning: allocates a new backing array)
	//
	lyric = append([]string{"yesterday"}, lyric...)

	// ------------------------------------------------------------------------
	// #2: Put the words to the correct positions in the `lyric` slice.
	// ------------------------------------------------------------------------

	//
	// yesterday all my troubles seemed so far away oh I believe in yesterday
	//                                              |               |
	//                                              v               v
	//                                       index: 8          pos: 13
	//
	const N, M = 8, 13

	// --- BEFORE ---
	//
	// yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay
	//
	// --- AFTER ---
	// yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay oh I believe in yesterday
	//                                              ^
	//
	//                                              |
	//                                              +--- duplicate
	//
	// (warning: allocates a new backing array)
	lyric = append(lyric, lyric[N:M]...)

	//
	// --- BEFORE ---
	// yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay oh I believe in yesterday
	//
	// --- AFTER ---
	// yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday
	//
	// (does not allocate a new backing array because cap(lyric[:N]) > M)
	//
	lyric = append(lyric[:N], lyric[M:]...)

	// ------------------------------------------------------------------------
	// #3: Print
	// ------------------------------------------------------------------------
	fmt.Printf("%s\n", lyric)
}

//adv-ops-practice

func main2() {
	// ########################################################
	//
	// #1: Create a string slice: `names` with a length and
	//     capacity of 5, and print it.
	//
	names := make([]string, 5)
	s.Show("1st step", names)

	// ########################################################
	//
	// #2: Append the following names to the names slice:
	//
	//     "einstein", "tesla", "aristotle"
	//
	//     Print the names slice.
	//
	//     Observe how the slice and its backing array change.
	//
	//     Expected output:
	//     ["" "" "" "" "" "einstein", "tesla", "aristotle"]
	//
	names = append(names, "einstein", "tesla", "aristotle")
	s.Show("2nd step", names)

	// ########################################################
	//
	// #3: Overwrite the name slice by creating a new slice
	//     using make().
	//
	//     Adjust the make() function so that it creates a
	//     slice with capacity of 5, and puts the slice pointer
	//     to the first index.
	//
	//     Then append the following names to the slice:
	//
	//     "einstein", "tesla", "aristotle"
	//
	//     Expected output:
	//     ["einstein", "tesla", "aristotle" "" ""]
	//
	//     So: Overwrite and print the names slice.
	//
	names = make([]string, 0, 5)
	names = append(names, "einstein", "tesla", "aristotle")
	s.Show("3rd step", names)

	// ########################################################
	//
	// #4: Copy only the first two elements of the following
	//     array to the last two elements of the `names` slice.
	//
	//     Print the names slice, you should see 5 elements.
	//     So, do not forget extending the slice.
	//
	//     Observe how its backing array stays the same.
	//
	moreNames := [...]string{"plato", "khayyam", "ptolemy"}
	copy(names[3:5], moreNames[:2])
	names = names[:cap(names)]

	s.Show("4th step", names)

	// ########################################################
	//
	// #5:  Only copy the last 3 elements of the `names` slice
	//      to a new slice: `clone`.
	//
	//     Append the first two elements of the `names` to the
	//    `clone`.
	//
	//     Ensure that after appending no new backing array
	//     allocations occur for the `clone` slice.
	//
	//     Print the clone slice before and after the append.
	//
	clone := make([]string, 3, 5)
	copy(clone, names[len(names)-3:])
	s.Show("5th step (before append)", clone)

	clone = append(clone, names[:2]...)
	s.Show("5th step (after append)", clone)

	// ########################################################
	//
	// #6: Slice the `clone` slice between 2nd and 4th (inclusive)
	//     elements into a new slice: `sliced`.
	//
	//     Append "hypatia" to the `sliced`.
	//
	//     Ensure that new backing array allocation "occurs".
	//
	//       Change the 3rd element of the `clone` slice
	//       to "elder".
	//
	//       Doing so should not change any elements of
	//       the `sliced` slice.
	//
	//     Print the `clone` and `sliced` slices.
	//
	//
	sliced := clone[1:4:4]
	sliced = append(sliced, "hypatia")
	clone[2] = "elder"

	s.Show("6th step", clone, sliced)
}

// Don't mind about this function.
//
// For printing the slices: You can either use the
// prettyslice package or `fmt.Printf`.
func init() {
	s.PrintBacking = true // prints the backing array
	s.MaxPerLine = 10     // prints 10 slice elements per line
	s.Width = 60          // prints 60 character per line
}

//add lines

func main6() {
	lyric := strings.Fields(`yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday`)

	// `+3` because we're going to add 3 newline characters to the fix slice.
	fix := make([]string, len(lyric)+3)

	//
	// USE A SLICE TO STORE WHERE EACH SENTENCE ENDS
	//
	// + The first sentence has 8 words so its cutting index is 8.
	//
	//   yesterday all my troubles seemed so far away now it looks as though they are here to stay
	//                                               |
	//                                               v
	//                                cutting index: 8
	//
	//
	// + The second sentence has 10 words so its cutting index is 10.
	//
	//   now it looks as though they are here to stay oh I believe in yesterday
	//                                               |
	//                                               v
	//                                cutting index: 10
	//
	//
	// + The last sentence has 5 words so its cutting index is 5.
	//
	//   oh I believe in yesterday
	//                            |
	//                            v
	//             cutting index: 5
	//
	cutpoints := []int{8, 10, 5}

	//
	// `n` tracks how much we've moved inside the `lyric` slice.
	//
	// `i` tracks the sentence that we're on.
	//
	for i, n := 0, 0; n < len(lyric); i++ {
		//
		// copy to `fix` from the `lyric`
		//
		//   destination:
		//     fix[n+i] because we don't want to delete the previous copy.
		//     it moves sentence by sentence, using the cutpoints.
		//
		//   source:
		//     lyric[n:n+cutpoints[i]] because we want copy the next sentence
		//     beginning from the number of the last copied elements to the
		//     n+next cutpoint (the next sentence).
		//
		n += copy(fix[n+i:], lyric[n:n+cutpoints[i]])

		//
		// add a newline after each sentence.
		//
		// notice that the '\n' position slides as we move over.
		// that's why it's: `n+i`.
		//
		fix[n+i] = "\n"

		// uncomment to see how the fix slice changes.
		// s.Show("fix slice", fix)
	}

	s.Show("fix slice", fix)

	// print the sentences
	for _, w := range fix {
		fmt.Print(w)
		if w != "\n" {
			fmt.Print(" ")
		}
	}
}

func init() {
	//
	// YOU DON'T NEED TO TOUCH THIS
	//
	// This initializes some options for the prettyslice package.
	// You can change the options if you want.
	//
	// This code runs before the main function above.
	//
	// s.Colors(false)     // if your editor is light background color then enable this
	s.PrintBacking = true  // prints the backing arrays
	s.MaxPerLine = 5       // prints max 15 elements per line
	s.SpaceCharacter = '⏎' // print this instead of printing a newline (for debugging)
}

//print daily riquests

func main7() {
	// There are 3 request totals per day
	const N = 3

	// DAILY REQUESTS DATA (PER 8-HOUR)
	reqs := []int{
		500, 600, 250, // 1st day: 1350 requests
		200, 400, 50, // 2nd day: 650 requests
		900, 800, 600, // 3rd day: 2300 requests
		750, 250, 100, // 4th day: 1100 requests
		// grand total: 5400 requests
	}

	// ALSO TRY IT WITH THIS DATA:
	// reqs = []int{
	// 	500, 600, 250,
	// 	200, 400, 50,
	// 	900, 800, 600,
	// 	750, 250, 100,
	// 	150, 654, 235,
	// 	320, 534, 765,
	// 	121, 876, 285,
	// 	543, 642,
	// }

	// ================================================
	// Allocate a slice efficiently with the exact size needed.
	//
	// Find the `size` automatically using the `reqs` slice.
	// Do not type it manually.
	//
	l := len(reqs)
	size := l / N
	if l%N != 0 {
		size++
	}
	daily := make([][]int, 0, size)
	//
	// ================================================

	// ================================================
	// Group the `reqs` per day into the slice: `daily`.
	//
	// So the daily will be:
	// [
	//  [500, 600, 250]
	//  [200, 400, 50]
	//  [900, 800, 600]
	//  [750, 250, 100]
	// ]
	//
	for N <= len(reqs) {
		daily = append(daily, reqs[:N]) // append the daily requests
		reqs = reqs[N:]                 // move the slice pointer for the next day
	}

	// add the residual data
	if len(reqs) > 0 {
		daily = append(daily, reqs)
	}

	// ================================================

	// ================================================
	// Print the header
	fmt.Printf("%-10s%-10s\n", "Day", "Requests")
	fmt.Println(strings.Repeat("=", 20))

	// Print the data per day along with the totals
	var grand int

	for i, day := range daily {
		var sum int

		for _, req := range day {
			sum += req
			fmt.Printf("%-10d%-10d\n", i+1, req)
		}

		fmt.Printf("%9s %-10d\n\n", "TOTAL:", sum)

		grand += sum
	}

	fmt.Printf("%9s %-10d\n", "GRAND:", grand)
	// ================================================
}
