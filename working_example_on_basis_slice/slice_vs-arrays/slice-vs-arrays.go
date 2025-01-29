package main

import "fmt"

func main() {
	{
		// its length is part of its type
		var nums [5]int
		fmt.Printf("nums array: %#v\n", nums)
	}

	{
		// its length is not part of its type
		var nums []int
		fmt.Printf("nums slice: %#v\n", nums)

		fmt.Printf("len(nums) : %d\n", len(nums))

		// won't work: the slice is nil.
		// fmt.Printf("nums[0]: %d\n", nums[0])
		// fmt.Printf("nums[1]: %d\n", nums[1])
	}

	main1()
	main2()
}

//--

func main1() {
	var array [2]int

	// zero value of an array is zero-valued elements
	fmt.Printf("array       : %#v\n", array)

	// nope: arrays are fixed length
	// array[2] = 0

	var slice []int

	// zero value of a slice is nil
	fmt.Println("slice == nil?", slice == nil)

	// nope: they don't exist:
	// _ = slice[0]
	// _ = slice[1]

	// len function still works though
	fmt.Println("len(slice)  :", len(slice))

	// array's length is part of its type
	fmt.Printf("array's type: %T\n", array)

	// whereas, slice's length isn't part of its type
	fmt.Printf("slice's type: %T\n", slice)
}

//--

func main2() {
	var books [5]string
	books[0] = "dracula"
	books[1] = "1984"
	books[2] = "island"

	newBooks := [5]string{"ulysses", "fire"}
	if books == newBooks {
	}
	books = newBooks

	games := []string{"pokemon", "sims"}
	newGames := []string{"pacman", "doom", "pong"}
	newGames = games
	games = nil
	games = []string{}

	var ok string
	for i, game := range games {
		if game != newGames[i] {
			ok = "not "
			break
		}
	}

	if len(games) != len(newGames) {
		ok = "not "
	}

	fmt.Printf("games and newGames are %sequal\n\n", ok)

	fmt.Printf("books         : %#v\n", books)
	fmt.Printf("new books     : %#v\n", newBooks)
	fmt.Printf("games         : %T\n", games)
	fmt.Printf("new games     : %#v\n", newGames)
	fmt.Printf("games's length: %d\n", len(games))
	fmt.Printf("games's nil   : %t\n", games == nil)
}
