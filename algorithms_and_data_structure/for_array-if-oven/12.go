package main

import "fmt"

func main() {
	//массив состоит из чисел 10 8 1 2 3 45 12 20
	//1) посчитать сумму только четных чисел
	//2) вне цикла сделать сумму первого и последнего числа
	numbers := [...]int{10, 8, 1, 2, 3, 45, 12, 20}
	n := len(numbers)
	sumi := 0
	for i := 0; i < n; i++ {
		//i-индекс элемента 0 1 2 3 4 5 6 7 8 9
		//numbers[i] это значение элемента
		if numbers[i]%2 == 0 {
			sumi = sumi + numbers[i]
		}
	}
	fmt.Println(sumi)
	c := numbers[0] + numbers[n-1]
	fmt.Println(c)
}

// 52
// 30
