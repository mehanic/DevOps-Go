package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	newArr := GetOnlyEven(arr)
	fmt.Println(newArr)
	main1()
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

func main1() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	evens := []int{}

	for _, val := range arr {
		if val%2 == 0 {
			evens = append(evens, val)
		}
	}

	fmt.Println(evens)
}
