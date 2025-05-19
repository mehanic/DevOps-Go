package main

import "fmt"

func singleIn(in int) int {
	return in
}

func multIn(a, b int, c int) int {
	return a + b + c
}

func namedReturn() (out int) {
	out = 2
	return
}

func multipleReturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error happend")
	}
	return in, nil
}

func multipleNamedReturn(ok bool) (rez int, err error) {
	if ok {
		err = fmt.Errorf("some error happend")
		return
	}
	rez = 2
	return
}

func sum(in ...int) (result int) {
	fmt.Printf("in := %#v \n", in)
	for _, val := range in {
		result += val
	}
	return
}

func main() {
	// fmt.Println(multipleNamedReturn(true))
	// return

	nums := []int{1, 2, 3, 4}
	fmt.Println(nums, sum(nums...))
	return
}

// ╰─λ go run functions.go
// in := []int{1, 2, 3, 4}
// [1 2 3 4] 10
