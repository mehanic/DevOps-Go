package main

import "fmt"

func compareTriplets(a [3]int, b [3]int) [2]int {
	var alicePoints, bobPoints int

	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			alicePoints++ 
		} else if a[i] < b[i] {
			bobPoints++ 
		}
	}

	return [2]int{alicePoints, bobPoints}
}

func main() {
	var a, b [3]int
	fmt.Scan(&a[0], &a[1], &a[2]) 
	fmt.Scan(&b[0], &b[1], &b[2]) 

	result := compareTriplets(a, b)

	fmt.Println(result[0], result[1])
}
