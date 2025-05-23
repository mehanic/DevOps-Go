package main

import (
	"fmt"
)

/*
	型パラメータを持つ関数f[T Stringer]
*/
// func main() {
// 	fmt.Println(f([]MyInt{1,2,3,4}))
// }

// // T==StringerはMyInt依存のたメソッドを扱える
// func f[T Stringer](xs []T) []string {
// 	var result []string
// 	for _, x := range xs {
// 		result = append(result, x.String())
// 	}
// 	return result
// }

// // MyIntが持つメソッドを満たすインターフェイス
// type Stringer interface {
// 	String() string
// }

// type MyInt int

// // MyIntが持つメソッド
// func(i MyInt) String() string {
// 	return strconv.Itoa(int(i))
// }

/*
	型パラメータを持つ型Stack[T any]
*/
type Stack[T any] []T

func New[T any]() *Stack[T] {
	v := make(Stack[T], 0)
	return &v
}

func (s *Stack[T]) Push(x T) {
	(*s) = append((*s), x)
}

func (s *Stack[T]) Pop() T {
	v := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return v
}

func main() {
	s := New[string]() // 型引数：型パラメータに具体的な型を渡すための構文
	s.Push("hello")
	s.Push("world")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

// world
// hello
