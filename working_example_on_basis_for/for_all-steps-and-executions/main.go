package main

import (
	"fmt"
)

const constWidth int = 320

func main() {
	// Константы
	fmt.Println("Width =", constWidth)
	// Функций
	fmt.Println(test())
	fmt.Println(test2())
	s, s2 := test()
	fmt.Println(s, s2)
	fmt.Println(test3())
	fmt.Println(test4())
	// Циклы
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum =", sum)

	sum2 := 0
	for sum2 < 100 { // аналог while
		sum2 += 10
	}
	fmt.Println("sum2 =", sum2)

	sum3 := 0
	for sum3 < 1000 { // аналог while
		sum3 += 10
	}
	fmt.Println("sum3 =", sum3)

	// Ветвления
	a := 0
	for a < 100 {
		if a == 10 {
			fmt.Println("a is 10")
		} else {
			fmt.Println(fmt.Sprintf("a is not 10. a=%d", a))
		}
		a++
	}

	a2 := 0
	for a2 < 100 {
		i := isTest(a2)
		if i == 1 {
			fmt.Println("a2 = 2")
		} else if i == 2 {
			fmt.Println("a2 = 3")
		} else {
			fmt.Println(fmt.Sprintf("unknown a2. a2=%d", a2))
		}
		a2++
	}

	a3 := 0
	for a3 < 100 {
		if i := isTest(a3); i == 1 {
			fmt.Println("a3 = 2")
		} else if i == 2 {
			fmt.Println("a3 = 3")
		} else {
			fmt.Println(fmt.Sprintf("unknown a3. a3=%d", a3))
		}
		a3++
	}

	a4 := 0
	for a4 < 100 {
		j := isTest(a4)
		switch j {
		case 1:
			fmt.Println("a4 = 2")
		case 2:
			fmt.Println("a4 = 3")
		default:
			fmt.Println(fmt.Sprintf("unknown a4. a4=%d", a4))
		}
		a4++
	}

	a5 := 0
	for a5 < 100 {
		switch j := isTest(a5); j {
		case 1:
			fmt.Println("a5 = 2")
		case 2:
			fmt.Println("a5 = 3")
		default:
			fmt.Println(fmt.Sprintf("unknown a5. a5=%d", a5))
		}
		a5++
	}
	// Отложенное выполнения операций
	// stack, Last In First Out
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("!!! Отложенное выполнения операций !!!")

}

// Функций
func test() (string, string) {
	a := "hello"
	b := "world!"
	return a, b
}

func test2() string {
	return "hello world!!"
}

func test3() (a, b string) {
	a = "hello"
	b = "world!!!"
	return a, b
}

func test4() (a, b int) {
	a = 1
	b = 3
	return a, b
}

// Циклы
func isTest(a int) int {
	if a == 2 {
		return 1
	} else if a == 3 {
		return 2
	}
	return 3
}
