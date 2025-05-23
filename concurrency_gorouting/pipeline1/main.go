package main

import "fmt"

func main() {
	sqrResults := sqr(gen(2, 3, 4, 5, 6))
	arr := []int{}
	for sqr := range sqrResults {
		arr = append(arr, sqr)
	}
	fmt.Println(arr)
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func sqr(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range c {
			out <- num * num
		}
		close(out)
	}()
	return out
}

// [4 9 16 25 36]
