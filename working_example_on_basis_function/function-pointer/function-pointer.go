package main

import "fmt"

func changeValue(x int) {
	x = x * x
}
func main() {

	d := 5
	fmt.Println("d before:", d) // 5
	changeValue(d)              // изменяем значение
	fmt.Println("d after:", d)  // 5 - значение не изменилось
}
