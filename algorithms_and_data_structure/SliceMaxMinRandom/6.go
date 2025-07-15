package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func CreateH(rows, columns, randomRange int) [][]int {
	//rand.Seed(time.Now().UnixNano())
	arr := [][]int{}
	for i := 0; i < rows; i++ {
		a := []int{}
		for j := 0; j < columns; j++ {
			k := rand.Intn(randomRange)
			a = append(a, k)
		}
		arr = append(arr, a)
	}
	return arr
}

func getSumi(a []int) int {
	sumi := 0
	for i := 0; i < len(a); i++ {
		sumi += a[i]
	}
	return sumi
}

func getMaxIndexAndSum(a []int) (int, int) {
	maxi := a[0]
	index := 0
	for i := 0; i < len(a); i++ {
		if a[i] > maxi {
			maxi = a[i]
			index = i
		}
	}
	return maxi, index
}

func FuncT(a [][]int) {
	var sumi []int
	for _, v := range a {
		fmt.Println(v)
		mySum := getSumi(v)
		sumi = append(sumi, mySum)
	}
	maxSumi, maxIndex := getMaxIndexAndSum(sumi)
	fmt.Println("Сумма элементов:= ", sumi)
	fmt.Println("Максимальная сумма массива = ", maxSumi)
	fmt.Println("Строка максимального массива = ", maxIndex)
}

func main() {
	arr := CreateH(10, 6, 10)
	FuncT(arr)
}

// [2 5 9 6 1 8]
// [2 6 9 3 4 0]
// [2 4 3 1 8 8]
// [0 8 5 7 3 2]
// [4 6 5 6 3 5]
// [2 8 1 5 6 9]
// [3 0 3 5 9 3]
// [1 3 7 1 5 2]
// [0 7 2 1 8 7]
// [3 9 6 9 4 4]
// Сумма элементов:=  [31 24 26 25 29 31 23 19 25 35]
// Максимальная сумма массива =  35
// Строка максимального массива =  9
