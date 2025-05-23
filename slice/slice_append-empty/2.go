package main

import "fmt"

func main() {
	//slice
	sl := []int{}
	//add element to the same slice
	sl = append(sl, 4)
	//скопирует значения из sl и добавит 3
	slNew := append(sl, 3)
	fmt.Println(sl)
	fmt.Println(slNew)
}
