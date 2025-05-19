package main

import "fmt"

func main() {

	var f, f2 func(s string) int
	f = func(s string) int {
		return len(s)
	}
	f2 = func(s string) int {
		return 2
	}
	// f2 = func() int { return 1 } // ошибка

	_, _ = f(""), f2("")

	result1 := f("hello")  // длина строки "hello" = 5
	result2 := f2("world") // всегда 2, независимо от входа

	fmt.Println("f(\"hello\") =", result1)  // → f("hello") = 5
	fmt.Println("f2(\"world\") =", result2) // → f2("world") = 2
}
