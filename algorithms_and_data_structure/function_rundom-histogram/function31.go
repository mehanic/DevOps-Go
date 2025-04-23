package main

import (
	"fmt"
	"math/rand"
	"time"
)

// example1 creates a histogram using a basic map.
func example1(source []int) map[int]int {
	histogram := make(map[int]int)
	for _, item := range source {
		if _, exists := histogram[item]; !exists {
			histogram[item] = 0
		}
		histogram[item]++
	}
	return histogram
}

// example2 creates a histogram using map's setdefault functionality.
func example2(source []int) map[int]int {
	histogram := make(map[int]int)
	for _, item := range source {
		if _, exists := histogram[item]; !exists {
			histogram[item] = 0
		}
		histogram[item]++
	}
	return histogram
}

// example3 creates a histogram using a default map with int zero value.
func example3(source []int) map[int]int {
	histogram := make(map[int]int)
	for _, item := range source {
		histogram[item]++
	}
	return histogram
}

// example4 creates a histogram using the Counter from the standard library.
func example4(source []int) map[int]int {
	histogram := make(map[int]int)
	for _, item := range source {
		histogram[item]++
	}
	return histogram
}

// show prints a textual histogram of the data.
func show(histogram map[int]int) {
	maxValue := 0
	for _, v := range histogram {
		if v > maxValue {
			maxValue = v
		}
	}
	for k, v := range histogram {
		bar := ""
		if maxValue > 0 {
			barLength := int(50 * float64(v) / float64(maxValue))
			for i := 0; i < barLength; i++ {
				bar += "*"
			}
		}
		fmt.Printf("%3d | %s\n", k, bar)
	}
}

// samples generates a list of random samples.
func samples(limit int, generator <-chan int) []int {
	result := make([]int, 0, limit)
	for i := 0; i < limit; i++ {
		result = append(result, <-generator)
	}
	return result
}

// arrival1 generates random arrivals based on the given limit.
func arrival1(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			out <- rand.Intn(n)
		}
	}()
	return out
}

// arrival2 generates arrivals based on a random walk with the given limit.
func arrival2(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		p := 0
		for {
			step := rand.Intn(3) - 1 // Randomly choose -1, 0, or +1
			p += step
			out <- abs(p) % n
		}
	}()
	return out
}

// abs returns the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Examples of histograms
	fmt.Println("\narrival1")
	show(example1(samples(1000, arrival1(8))))
	fmt.Println("\narrival1")
	show(example2(samples(1000, arrival1(8))))
	fmt.Println("\narrival1")
	show(example3(samples(1000, arrival1(8))))
	fmt.Println("\narrival1")
	show(example4(samples(1000, arrival1(8))))

	fmt.Println("\narrival2")
	show(example1(samples(1000, arrival2(8))))

	fmt.Println("\nexample4 output:")
	for k, v := range example4(samples(1000, arrival1(8))) {
		fmt.Printf("%d: %d\n", k, v)
	}
}


// arrival1
//   3 | ************************************************
//   0 | **************************************************
//   4 | **********************************************
//   5 | *************************************************
//   1 | *****************************************
//   2 | ************************************************
//   7 | ***********************************************
//   6 | **********************************************

// arrival1
//   7 | *********************************************
//   1 | ****************************************
//   5 | **********************************************
//   2 | ******************************************
//   6 | **************************************************
//   4 | ***********************************************
//   3 | ******************************************
//   0 | ***********************************************

// arrival1
//   4 | ***********************************************
//   3 | *******************************************
//   0 | **************************************************
//   1 | **********************************
//   2 | ****************************************
//   7 | ***********************************
//   6 | *****************************************
//   5 | *********************************************

// arrival1
//   6 | *********************************************
//   5 | ***********************************************
//   1 | ******************************************
//   3 | **********************************************
//   0 | ****************************************
//   7 | *********************************************
//   2 | **************************************************
//   4 | **********************************************

// arrival2
//   5 | **********************************************
//   6 | ***************************************
//   7 | *********************************
//   0 | ***************************************
//   1 | **************************************************
//   2 | *********************************************
//   3 | *******************************************
//   4 | *******************************************

// example4 output:
// 3: 116
// 7: 110
// 4: 136
// 0: 137
// 6: 113
// 5: 117
// 2: 149
// 1: 122

