package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	numbers := []int{}
	for i := 0; i < n; i++ {
		var d int
		fmt.Scan(&d)
		numbers = append(numbers, d)
	}
	posSum := 0
	maxi := numbers[0]
	mini := numbers[0]
	maxiIndex := 0
	miniIndex := 0
	proz := 1
	start := 0
	end := 0
	for i, v := range numbers {
		if v > 0 {
			posSum = posSum + v
		}
		if v > maxi {
			maxi = v
			maxiIndex = i
		}
		if v < mini {
			mini = v
			miniIndex = i
		}
	}
	if maxiIndex > miniIndex {
		start = miniIndex
		end = maxiIndex
	} else {
		start = maxiIndex
		end = miniIndex
	}
	for i := start + 1; i < end; i++ {
		proz = proz * numbers[i]
	}
	fmt.Println(posSum, proz)
}
