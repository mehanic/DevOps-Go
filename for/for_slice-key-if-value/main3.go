package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{4, 5, 6, 7}
	arr3 := []int{7, 8, 9, 10, 123}
	arr := [][]int{arr1, arr2, arr3}
	arr4 := []int{111, 2, 3233, 341}
	arr = append(arr, arr4)
	counter := 0
	//вывод всех элементов массива
	for _, v := range arr {
		for _, k := range v {
			if k%2 == 0 {
				counter++
			}
		}
	}
	fmt.Println(counter)
}
