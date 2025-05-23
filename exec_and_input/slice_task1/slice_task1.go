package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"unsafe"
)

// append-sort-nums
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("provide a few numbers")
		return
	}

	var nums []int
	for _, s := range args {
		n, err := strconv.Atoi(s)
		if err != nil {
			continue
		}

		nums = append(nums, n)
	}

	sort.Ints(nums)
	fmt.Println(nums)
	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
	main7()
	main8()
	main9()
}

//housing prise

func main1() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		locs                       []string
		sizes, beds, baths, prices []int
	)

	rows := strings.Split(data, "\n")

	for _, row := range rows {
		cols := strings.Split(row, separator)

		locs = append(locs, cols[0])

		n, _ := strconv.Atoi(cols[1])
		sizes = append(sizes, n)

		n, _ = strconv.Atoi(cols[2])
		beds = append(beds, n)

		n, _ = strconv.Atoi(cols[3])
		baths = append(baths, n)

		n, _ = strconv.Atoi(cols[4])
		prices = append(prices, n)
	}

	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-15s", h)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	for i := range rows {
		fmt.Printf("%-15s", locs[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", prices[i])
		fmt.Println()
	}
}

//housing prise avarage

func main2() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		locs                       []string
		sizes, beds, baths, prices []int
	)

	rows := strings.Split(data, "\n")

	for _, row := range rows {
		cols := strings.Split(row, separator)

		locs = append(locs, cols[0])

		n, _ := strconv.Atoi(cols[1])
		sizes = append(sizes, n)

		n, _ = strconv.Atoi(cols[2])
		beds = append(beds, n)

		n, _ = strconv.Atoi(cols[3])
		baths = append(baths, n)

		n, _ = strconv.Atoi(cols[4])
		prices = append(prices, n)
	}

	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-15s", h)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	for i := range rows {
		fmt.Printf("%-15s", locs[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", prices[i])
		fmt.Println()
	}

	// -------------------------------------------------------------------------
	// print the averages
	//
	// evil note:
	// later on, you will see how easy it would be to solve this
	// using the functions instead. there are a lot of repetitive code here.
	// -------------------------------------------------------------------------
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	// jump over the location column
	fmt.Printf("%-15s", "")

	var sum int

	for _, n := range sizes {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(sizes)))

	sum = 0
	for _, n := range beds {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(beds)))

	sum = 0
	for _, n := range baths {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(baths)))

	sum = 0
	for _, n := range prices {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(prices)))

	fmt.Println()
}

//slicing basics

func main3() {
	data := "2 4 6 1 3 5"
	splitted := strings.Fields(data)

	var nums []int
	for _, s := range splitted {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}

	// rest of the code for this exercise
	fmt.Println("nums        :", nums)

	evens, odds := nums[:3], nums[3:]

	fmt.Println("evens       :", evens)
	fmt.Println("odds        :", odds)

	fmt.Println("middle      :", nums[2:4])
	fmt.Println("first 2     :", nums[:2])
	fmt.Println("last 2      :", nums[len(nums)-2:])

	fmt.Println("evens last 1:", evens[len(evens)-1:])
	fmt.Println("odds last 2 :", odds[len(odds)-2:])
}

//slicing by args

func main4() {
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}

	fmt.Printf("%q\n\n", ships)

	from, to := 0, len(ships)

	switch poss := os.Args[1:]; len(poss) {
	default:
		fallthrough
	case 0:
		fmt.Println("Provide only the [starting] and [stopping] positions")
		return
	case 2:
		to, _ = strconv.Atoi(poss[1])
		fallthrough
	case 1:
		from, _ = strconv.Atoi(poss[0])
	}

	if l := len(ships); from < 0 || from > l || to > l || from > to {
		fmt.Println("Wrong positions")
		return
	}

	fmt.Println(ships[from:to])
}

//slicing housing prices

func main5() {
	const (
		data = `Location,Size,Beds,Baths,Price
New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	// parse the data
	rows := strings.Split(data, "\n")
	cols := strings.Split(rows[0], separator)

	// default case: slice for all the columns
	from, to := 0, len(cols)

	// find the from and to positions depending on the command-line arguments
	args := os.Args[1:]
	for i, v := range cols {
		l := len(args)

		if l >= 1 && v == args[0] {
			from = i
		}

		if l == 2 && v == args[1] {
			to = i + 1 // +1 because the stopping index is a position
		}
	}

	// "from" cannot be greater than or equal to "to": reset invalid arg to 0
	if from >= to {
		from = 0
	}

	for i, row := range rows {
		cols := strings.Split(row, separator)

		// print only the requested columns
		for _, h := range cols[from:to] {
			fmt.Printf("%-15s", h)
		}
		fmt.Println()

		// print extra new line for the header
		if i == 0 {
			fmt.Println()
		}
	}
}

//internals-backing-array-fix

func main6() {
	nums := []int{56, 89, 15, 25, 30, 50}

	// ----------------------------------------
	// breaks the connection:
	// mine and nums now have different backing arrays

	// verbose solution:
	// var mine []int
	// mine = append(mine, nums[:3]...)

	// better solution (almost the same thing):
	mine := append([]int(nil), nums[:3]...)
	// ----------------------------------------

	mine[0], mine[1], mine[2] = -50, -100, -150
	fmt.Println("Mine         :", mine)
	fmt.Println("Original nums:", nums[:3])
}

//internals-backing-array-sort

func main7() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)

	mid := len(items) / 2
	smid := items[mid-1 : mid+2]

	// sorting the smid will affect the items
	// as well. their backing array is the same.
	sort.Strings(smid)

	fmt.Println()
	fmt.Println("Sorted  :", items)
}

//internals-slice-header

const size = 1e7

func main8() {
	debug.SetGCPercent(-1)

	report("initial memory usage")

	var array [size]int
	report("after declaring an array")

	array2 := array
	report("after copying the array")

	passArray(array)

	slice1 := array[:]
	slice2 := array[:1e3]
	slice3 := array[1e3:1e4]
	report("after slicings")

	passSlice(slice3)

	fmt.Println()
	fmt.Printf("Array's size : %d bytes.\n", unsafe.Sizeof(array))
	fmt.Printf("Array2's size: %d bytes.\n", unsafe.Sizeof(array2))
	fmt.Printf("Slice1's size: %d bytes.\n", unsafe.Sizeof(slice1))
	fmt.Printf("Slice2's size: %d bytes.\n", unsafe.Sizeof(slice2))
	fmt.Printf("Slice3's size: %d bytes.\n", unsafe.Sizeof(slice3))
}

func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

func passSlice(items []int) {
	report("inside passSlice")
}

func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}

//observe-len-cap

func main9() {
	// --- #1 ---
	// var games []string
	// fmt.Printf("games's len    : %d cap: %d\n", len(games), cap(games))

	// games := []string{}
	// fmt.Printf("games's len    : %d cap: %d\n", len(games), cap(games))

	// games = append(games, "pacman", "mario", "tetris", "doom")
	// fmt.Printf("games's len    : %d cap: %d\n", len(games), cap(games))

	games := []string{"pacman", "mario", "tetris", "doom"}
	fmt.Printf("games's len    : %d cap: %d\n", len(games), cap(games))

	// --- #2 ---
	fmt.Println()

	for i := 0; i <= len(games); i++ {
		s := games[:i]
		fmt.Printf("games[:%d]'s len: %d cap: %d\n", i, len(s), cap(s))
	}

	// --- #3 ---
	fmt.Println()

	zero := games[:0]
	fmt.Printf("games's     len: %d cap: %d\n", len(games), cap(games))
	fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))

	for _, v := range []string{"ultima", "dagger", "pong", "coldspot", "zetra"} {
		zero = append(zero, v)
		fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))
	}

	// --- #4 ---
	fmt.Println()

	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", n, len(s), cap(s))
	}

	// --- #5 ---
	fmt.Println()

	zero = zero[:cap(zero)]
	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", n, len(s), cap(s), s)
	}

	// --- #6 ---
	fmt.Println()

	zero[0] = "command & conquer"
	games[0] = "red alert"
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", games)

	// --- #7 ---
	// uncomment and see the error.
	// _ = games[:cap(games)+1]
	// or:
	// _ = games[:5]
}
