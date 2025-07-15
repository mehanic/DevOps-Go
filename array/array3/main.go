package main

import "fmt"

func main() {
	// одномерный массив
	arr1 := []int{1, 2, 3, 4}
	fmt.Println(arr1[1])
	// двумерный массив arr[номер_строки][номер_колонки]
	arr2 := [][]int{
		{1, 2, 3, 4},
		{4, 5, 6, 7},
	}
	//	0 1 2 3
	// 0	1 2 3 4
	// 1	4 5 6 7
	fmt.Println(arr2[0][1])
	main1()
}

func main1() {
	arr := [][]int{
		{1, 2, 3, 4},
		{4, 5, 6, 7},
	}

	// добавляем новую строку
	arr = append(arr, []int{8, 9, 10, 11})
	fmt.Println(arr)
	arr[1][2] = 99 // меняем число 6 на 99
	fmt.Println(arr)
}
