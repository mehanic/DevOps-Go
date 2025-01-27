package main

import "fmt"


func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	newArr := GetOnlyEven(arr)
	fmt.Println(newArr)
}

func GetOnlyEven(arr []int) []int {
	evens := []int{}
	for i := 0; i < len(arr); i++ {
		if CheckForEven(arr[i]) {
			evens = append(evens, arr[i])
		}
	}
	return evens
}

func CheckForEven(k int) bool {
	if k%2 == 0 {
		return true
	}
	return false
}

// [2 4 6 8]
