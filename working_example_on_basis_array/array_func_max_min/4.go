package main

import (
	"fmt"
	"math"
)

func getMax(a []int) (int, int) {
	maxi := a[0]
	mini := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > maxi {
			maxi = a[i]
		}
		if a[i] < mini {
			mini = a[i]
		}
	}
	return maxi, mini
}

// Return Values: It returns two integers: the maximum (maxi) and minimum
// (mini) values in the slice.
// Logic: It initializes both maxi and mini to the first element of the
// array and iterates through the slice, updating these values as necessary.
func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{4, 5, 6, 2, 323, 23, 4, 232}
	arr3 := []int{1, 23, 4, 123, 23213, 43434, 4545, 12321, 1313213, 343443, 54565}
	maxi1, mini1 := getMax(arr1)
	maxi2, mini2 := getMax(arr2)
	maxi3, mini3 := getMax(arr3)
	//Output: It prints the maximum and minimum values for each array.
	fmt.Println(maxi1, maxi2, maxi3)
	fmt.Println(mini1, mini2, mini3)
	//Mathematical Operation: The program calculates 3 in 2 using math.Pow
	k := math.Pow(float64(3), float64(2))
	fmt.Println(k)
}

// 4 323 1313213
// 1 2 1
// 9
