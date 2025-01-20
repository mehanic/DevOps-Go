package main

import "fmt"

func main() {

	i := 1

	for i <= 3 {
		fmt.Println(i)
		i++ // Equivalent i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {

		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}

	var fruits = [3]string{"apple", "pear", "baban"}
	for index, value := range fruits {
		fmt.Println(index, value)
	}

	for _, value := range fruits {
		fmt.Println(value)
	}

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}
}

// go run /home/mehanic/structure/for/for_different-variants/for.go
// 1
// 2
// 3
// 7
// 8
// 9
// loop
// 1
// 3
// 5
// 0 apple
// 1 pear
// 2 baban
// apple
// pear
// baban
// apple
// pear
// baban
