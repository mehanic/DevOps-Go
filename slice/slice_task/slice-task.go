package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"time"
)

// declare-nil
func main() {
	var (
		names     []string  // The names of your friends
		distances []int     // The distances
		data      []byte    // A data buffer
		ratios    []float64 // Currency exchange ratios
		alives    []bool    // Up/Down status of web servers
	)

	fmt.Printf("names    : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data     : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios   : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives   : %T %d %t\n", alives, len(alives), alives == nil)
	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
	main7()
	main8()
}

// empty
func main1() {
	var (
		names     []string  // The names of your friends
		distances []int     // The distances
		data      []byte    // A data buffer
		ratios    []float64 // Currency exchange ratios
		alives    []bool    // Up/Down status of web servers
	)

	names = []string{}
	distances = []int{}
	data = []byte{}
	ratios = []float64{}
	alives = []bool{}

	fmt.Printf("names    : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data     : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios   : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives   : %T %d %t\n", alives, len(alives), alives == nil)
}

// slice-literal
func main2() {
	var (
		names     []string  // The names of your friends
		distances []int     // The distances
		data      []byte    // A data buffer
		ratios    []float64 // Currency exchange ratios
		alives    []bool    // Up/Down status of web servers
	)

	names = []string{"serpil", "ebru", "lina"}
	distances = []int{100, 200, 300, 400, 500}
	data = []byte{'I', 'N', 'A', 'N', 'C'}
	ratios = []float64{3.14, 6.28}
	alives = []bool{true, false, false, true}

	fmt.Printf("names    : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data     : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios   : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives   : %T %d %t\n", alives, len(alives), alives == nil)

	if len(distances) == len(data) {
		fmt.Println("The length of the distances and the data slices are the same.")
	}
}

// declare-array-as-slices
func main3() {
	names := []string{"Einstein", "Tesla", "Shepard"}
	distances := []int{50, 40, 75, 30, 125}
	data := []byte{'H', 'E', 'L', 'L', 'O'}
	ratios := []float64{3.14145}
	alives := []bool{true, false, true, false}
	zero := []byte{}

	fmt.Printf("names    : %[1]T %[1]q\n", names)
	fmt.Printf("distances: %[1]T %[1]d\n", distances)
	fmt.Printf("data     : %[1]T %[1]d\n", data)
	fmt.Printf("ratios   : %[1]T %.2[1]f\n", ratios)
	fmt.Printf("alives   : %[1]T %[1]t\n", alives)
	fmt.Printf("zero     : %[1]T %[1]d\n", zero)
}

// fix-the-problems
func main4() {
	names := []string{"Einstein", "Shepard", "Tesla"}

	books := []string{"Stay Golden", "Fire", "Kafka's Revenge"}
	sort.Strings(books)

	nums := [...]int{5, 1, 7, 3, 8, 2, 6, 9}
	sort.Ints(nums[:])

	fmt.Printf("%q\n", strings.Join(names, " and "))
	fmt.Printf("%q\n", books)
	fmt.Printf("%d\n", nums)
}

// compare-the-slices
func main5() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	namesC := strings.Split(namesA, ", ")

	sort.Strings(namesC)
	sort.Strings(namesB)

	if len(namesC) == len(namesB) {
		for i := range namesC {
			if namesC[i] != namesB[i] {
				return
			}
		}
		fmt.Println("They are equal.")
	}
}

// append
func main6() {
	png, header := []byte{'P', 'N', 'G'}, []byte{}

	header = append(header, png...)

	if bytes.Equal(png, header) {
		fmt.Println("They are equal")
	}
}

// append-2
func main7() {
	var pizza []string
	var departures []time.Time
	var grads []int
	var lights []bool

	pizza = append(pizza, "pepperoni", "onions", "extra cheese")
	now := time.Now()
	departures = append(departures,
		now,
		now.Add(time.Hour*24), // 24 hours after `now`
		now.Add(time.Hour*48)) // 48 hours after `now`

	grads = append(grads, 1998, 2005, 2018)
	lights = append(lights, true, false, true)

	fmt.Printf("pizza       : %s\n", pizza)
	fmt.Printf("\ngraduations : %d\n", grads)
	fmt.Printf("\ndepartures  : %s\n", departures)
	fmt.Printf("\nlights      : %t\n", lights)
}

// append-3
func main8() {
	toppings := []string{"black olives", "green peppers"}

	var pizza []string
	pizza = append(pizza, toppings...)
	pizza = append(pizza, "onions", "extra cheese")

	fmt.Printf("toppings: %s\n", pizza)
}
