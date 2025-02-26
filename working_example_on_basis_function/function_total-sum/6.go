package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getNewArray(n int) []int {
	arr := []int{}
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < n; i++ {
		k := rand.Intn(100)
		arr = append(arr, k)
	}
	return arr
}

func getSumArr(arr []int) int {
	sumi := 0
	for i := 0; i < len(arr); i++ {
		sumi += arr[i]
	}
	return sumi
}

//функция для посчета суммы элементов в массиве и возвращать эту сумму
func main() {
	a := getNewArray(10)
	b := getNewArray(20)
	c := getNewArray(30)
	d := getNewArray(40)
	allArr := [][]int{a, b, c, d}
	maxi := 0
	maxiIndex := 0
	for i := 0; i < len(allArr); i++ {
		sumi := getSumArr(allArr[i])
		if sumi > maxi {
			maxi = sumi
			maxiIndex = i
		}
	}
	fmt.Println(maxiIndex)
}
