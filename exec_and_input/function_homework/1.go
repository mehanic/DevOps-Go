package main

import "fmt"


func direction(n,m int)(func()int,func()int){
	return func()int{ return n+m},func()int{ return n-m}
}

func main() {

	plus, minus := direction(20, 30)

	fmt.Println(plus())  // 50
	fmt.Println(minus()) // -10
}

// 50
// -10
