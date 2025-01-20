package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{4, 5, 6, 7}
	arr3 := []int{7, 8, 9, 10}
	arr := [][]int{arr1, arr2, arr3}
	arr4 := []int{111, 2, 3233, 341}
	arr = append(arr, arr4)
	n := len(arr)
	m := len(arr1)
	for i := 0; i < m; i++ {
		sumiColumn := 0
		for j := 0; j < n; j++ {
			sumiColumn += arr[j][i]
		}
		fmt.Println(sumiColumn)
	}
	////выводить все значения по строками
	//arr[0][0] arr[0][1] arr[0][2] arr[0][3]
	////выводить все значения по столбцам
	//arr[0][0] arr[1][0] arr[2][0] arr[3][0]
}
