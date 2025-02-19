// package main

// import "fmt"

// func main() {
// 	a := 2

// 	b := &a

// 	fmt.Println("a = ", a)
// 	fmt.Println("b = ", b)

// 	fmt.Println("Выполняем: \n *b = 3")
// 	*b = 3
// 	fmt.Println("a = ", a)

// 	// с - новый указатель на a
// 	c = &a

// 	// получение указателя на переменную типа int
// 	// инициализировано значением по умолчанию
// 	d := new(int)

// 	*d = 12
// 	*c = *d // c = 12 -> a = 12
// 	*d = 13 // c и a не изменились

// 	c = d   // теперь с указывает туда же, куда d
// 	*c = 14 // c = 14 -> d = 14, a = 12

// }

package main

import "fmt"

func main() {
	a := 2
	b := &a

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	fmt.Println("Выполняем: \n *b = 3")
	*b = 3
	fmt.Println("a = ", a)

	// c - новый указатель на a
	c := &a

	// получение указателя на переменную типа int
	// инициализировано значением по умолчанию
	d := new(int)

	*d = 12
	*c = *d // c = 12 -> a = 12
	*d = 13 // c и a не изменились

	c = d   // теперь с указывает туда же, куда d
	*c = 14 // c = 14 -> d = 14, a = 12

	fmt.Println("a = ", a)
	fmt.Println("*d = ", *d)
	fmt.Println("*c = ", *c)
}

// ╰─λ go run 5_pointers.go                                                     0 (0.001s) < 22:00:23
// a =  2
// b =  0xc0000120d0
// Выполняем:
//  *b = 3
// a =  3
// a =  12
// *d =  14
// *c =  14
