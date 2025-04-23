package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 4}
	index := 2
	value := 99
	newArray := append(array[:index], append([]int{value}, array[index:]...)...)
	fmt.Println(newArray) // [1, 2, 99, 3, 4]
	main1()
}

func main1() {
	array := []int{1, 2, 3, 4}
	rmIndex := 2
	array = append(array[:rmIndex], array[rmIndex+1:]...)
	fmt.Println(array) // [1, 2, 4]
}
