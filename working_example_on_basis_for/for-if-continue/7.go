package main

import "fmt"

func main() {
	fmt.Println("=== Цикл от 0 до 10 (без числа 5) ===")
	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
