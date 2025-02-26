package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	hello()

	res := adder(3, 5)
	fmt.Println(res)

	tot := proadder(4, 2, 1, 41, 1231, 2)
	fmt.Println(tot)
	arr := initarr()
	fmt.Println(arr)
	min, mi := findmin(arr, len(arr))
	fmt.Println("Minimum element in array is", min, "stored at index", mi)
}

func hello() {
	fmt.Println("Hello from the function")
}

func adder(a int, b int) int {
	return a + b
}

func proadder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func initarr() []int {
	arr := make([]int, 0)
	arr = append(arr, 1211, 211, 123, 244, 453, 126, 172)
	return arr
}

func findmin(arr []int, n int) (val int, ind int) {
	min := arr[0]
	k := 0
	for i := 0; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
			k = i
		}
	}
	return min, k
}
