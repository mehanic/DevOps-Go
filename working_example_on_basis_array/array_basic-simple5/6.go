package main

import "fmt"

func main() {
	//в массиве который состоит из 5 чисел найти сумму всех элементов
	numbers := [...]int{1, 2, 3, 4, 5}
	sumi := 0
	sumi = sumi + numbers[0]
	sumi = sumi + numbers[1]
	sumi = sumi + numbers[2]
	sumi = sumi + numbers[3]
	sumi = sumi + numbers[4]
	fmt.Println(sumi)
}
