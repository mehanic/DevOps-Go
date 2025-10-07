package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 10}
	fmt.Println(len(sl))
	fmt.Println(cap(sl))
	lk := make([]int, 10)
	fmt.Println(len(lk))
	fmt.Println(cap(lk))
	lk = append(lk, 3)

	// 1. Літерал слайсу
	sl1 := []int{1, 2, 3, 4, 5, 6, 10}
	fmt.Println("sl:", sl1, "len:", len(sl), "cap:", cap(sl))

	// 2. Слайс через make
	lk1 := make([]int, 10)
	fmt.Println("lk:", lk1, "len:", len(lk), "cap:", cap(lk))

	// 3. append додає елемент
	lk = append(lk, 3)
	fmt.Println("lk після append:", lk, "len:", len(lk), "cap:", cap(lk))
}

// 7
// 7
// 10
// 10
