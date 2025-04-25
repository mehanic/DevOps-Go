package main

import "fmt"

func addPrefix(origin string) string {
	return "prefix_" + origin
}

func addPrefixWithErr(origin string) (string, error) {
	return "prefix_" + origin, nil
}

func addPrefixWithLen(origin string) (res string, length int) {
	res = "prefix_" + origin
	length = len(res)
	return res, length
}

func main() {
	myString := addPrefix("my_string")
	fmt.Println(myString) // prefix_my_string

	result := addPrefix("world")
	fmt.Println(result) // Выведет: prefix_world}

	myString, err := addPrefixWithErr("error_string")
	fmt.Println(err)      // nil
	fmt.Println(myString) // prefix_error_string

	withLen, length := addPrefixWithLen("go_lang")
	fmt.Printf("Результат: %s, длина: %d\n", withLen, length)

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
