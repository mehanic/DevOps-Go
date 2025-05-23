package main

import (
	"fmt"
)

func main() {
	var sum int

	// sum += 1
	// sum += 2
	// sum += 3
	// sum += 4
	// sum += 5

	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println(sum)
	main1()
	main2()
	main3()
}

func main1() {
	var (
		sum int
		i   = 1
	)

	for {
		if i > 5 {
			break
		}
		sum += i
		fmt.Println(i, "-->", sum)
		i++
	}
	fmt.Println(sum)
}

func main2() {
	var (
		sum int
		i   = 3
	)
	for {
		if i > 10 {
			break
		}
		if i%2 != 0 {
			// this continue creates an endless loop
			// since the code never increases the i
			continue
		}
		sum += i
		fmt.Println(i, "-->", sum)
		i++
	}
	fmt.Println(sum)
}

//not working

func main3() {
	var (
		sum int
		i   = 1
	)
	for {
		if i > 10 {
			break
		}
		if i%2 != 0 {
			// just by putting this here we solve the problem
			i++
			continue
		}
		sum += i
		fmt.Println(i, "-->", sum)
		i++
	}
	fmt.Println(sum)
}

//---

// 500500
// 1 --> 1
// 2 --> 3
// 3 --> 6
// 4 --> 10
// 5 --> 15
// 15
// 4
// 7
