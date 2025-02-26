package main

import "fmt"

// когда происходит вызов внещней функции происхожит создание внутр функции
func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	orderPrice := totalPrice(10)
	fmt.Println(orderPrice(1))
	fmt.Println(orderPrice(1))
}

// ╰─λ go run closures.go                                                                                                                                               0 (0.001s) < 12:51:10
// 11
// 12
