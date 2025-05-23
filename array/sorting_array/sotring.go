package main

import "fmt"

func main() {
	arr := [][]int{
		{3, -2, 6, 4},
		{8, 1, 12, 2},
		{5, 4, -8, 0},
	}
	//мы переворачивали массив чтобы колонки были строками а строки колонками
	newArr := [][]int{}
	n := len(arr)
	m := len(arr[0])
	for i := 0; i < m; i++ {
		tempArr := []int{}
		for j := 0; j < n; j++ {
			tempArr = append(tempArr, arr[j][i])
		}
		newArr = append(newArr, tempArr)
	}
	n = len(newArr)
	//менял массивы местами если их первые элементы меньше друг друга
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if newArr[i][0] < newArr[j][0] {
				temp := newArr[i]
				newArr[i] = newArr[j]
				newArr[j] = temp
			}
		}
	}
	//мы переворачивали массив чтобы колонки были строками а строки колонками
	m = len(newArr[0])
	result := [][]int{}
	for i := 0; i < m; i++ {
		tempArr := []int{}
		for j := 0; j < n; j++ {
			tempArr = append(tempArr, newArr[j][i])
		}
		result = append(result, tempArr)
	}
	fmt.Println(result)
}
