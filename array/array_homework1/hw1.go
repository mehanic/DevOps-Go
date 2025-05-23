package main

import "fmt"

func main() {
	for {
		var operation string
		fmt.Scanf("%s", &operation)
		if operation == "0" {
			break
		} else if operation == "+" || operation == "-" || operation == "*" || operation == "/" {
			var a, b int
			fmt.Scanf("%d", &a)
			fmt.Scanf("%d", &b)
			if operation == "+" {
				fmt.Println(a + b)
			}
			if operation == "*" {
				fmt.Println(a * b)
			}
		} else {
			fmt.Println("there is no this operation")
		}
	}
}
// -
// 78
// 70
// +
// 89
// 89
// 178
// *
// 7
// 8
// 56
