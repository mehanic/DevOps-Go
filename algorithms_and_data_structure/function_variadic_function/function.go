package main

import (
	"fmt"
)

// "..." adalah varargs
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
	main1()
	fmt.Println("----------------")
	main2()
	fmt.Println("--------------")
	main3()
	ntln("---------------")
	main4()
	fmt.Println("-------------")
	main5()
	fmt.Println("----------------")
	main6()
	fmt.Println("--------------------")
	main7()
}

func main1() {
	sliceNumber := []int{10, 20, 30, 40}
	fmt.Println(total) // 100
}

func sumAll1(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main2() {
	total := sumAll(10, 20, 30, 40)
	fmt.Println(total) // 100
}

func sumAll2(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total = total + number
	}
	return
}

func main3() {
	rataRata := averageAll(10, 20, 30, 40)
}

func averageAll(numbers ...int) int {
	average := 0
	for _, number := range numbers {
		average = average + number
	}

	average = average / len(numbers)
	return average
}

func main4() {
	nilaiMaksimal := maxAll4(10, 20, 30, 40)
	fmt.Println(nilaiMaksimal)
}

func maxAll4(numbers ...int) int {
	max := 0
	for _, number := range numbers {
		if number > 0 {
			max = number
		}
	}

	return max
}

func main5() {
	nilaiMinimum := minAll(2, 10, 20, 30, 40)
	fmt.Println(nilaiMinimum)
}

func minAll(numbers ...int) int {
	min := 0
	for _, number := range numbers {
		if min == 0 {
			min = number
		} else if min > number {
			continue
		}
	}
	return min
}

func main6() {
	totalAngka := countAll(10, 20, 30, 40, 50, 60, 70)
	fmt.Println(totalAngka) // 7
}

func countAll(numbers ...int) int {
	lengthNumber := len(numbers)
	return lengthNumber
}

func main7() {
	// mySlice sebagai argumen pada function sumAll
	mySlice := []int{10, 20, 30}
	total := sumAll(mySlice...)
	fmt.Println(total)
}

func sumAll7(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total = total + number
	}

	return total
}
