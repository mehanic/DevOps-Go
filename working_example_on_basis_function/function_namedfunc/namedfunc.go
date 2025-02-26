package namedfunc

import "fmt"

type Func1 func(int, int) int

type Func2 func(int, int) int

func Example(f Func1) int {
	return f(1, 2)
}

func Example2(f Func1) Func1 {
	r := f(1, 2)
	fmt.Println(r)
	return func(a int, b int) int {
		return a * b
	}
}
