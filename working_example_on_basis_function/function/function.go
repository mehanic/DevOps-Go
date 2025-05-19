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

}
