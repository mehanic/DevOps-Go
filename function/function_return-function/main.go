package main

import "fmt"

func test1(x int) int {
	return x + 1
}


func test4(p_func func(int) int){
	fmt.Println(p_func(40))
}


func returnFunc(x string) func(){
	return func() { fmt.Println(x) }
}

func calculate(x int) (result int) {
    result = x + 2
    return result
}

func main() {
	fmt.Println(test1(10))


	test2 := func(x int) int {
		return x + 2
	}
	fmt.Println(test2(20))


	test3 := func(x int) int {
		return x + 3
	}(30)

	fmt.Println(test3)


	test4(test2)

	returnFunc("hello")()
	returnFunc("world")()

	y := returnFunc("goodbye")
	y()

	//fmt.Println("x=", x, " y=", y, "z=", z)

}

// 11
// 22
// 33
// 42
// hello
// world
// goodbye
