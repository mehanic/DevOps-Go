package main

import (
	"fmt"
	"reflect"
)

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "1: エラーが発生しました。", ErrCode: 1234}
}

func main() {
	err := RaiseError()
	fmt.Println("err =>", err)
	fmt.Println("reflect.TypeOf(err) =>", reflect.TypeOf(err))
	fmt.Println("err.Error() =>", err.Error())
	// fmt.Println("err.Error() =>", err.Message()) // error型にキャストされてて、MyError型じゃなくなってるので、Message フィールドが存在しない。
	// fmt.Println("err.Error() =>", err.ErrCode()) // error型にキャストされてて、MyError型じゃなくなってるので、ErrCode フィールドが存在しない。

	fmt.Println("err.(*MyError) =>", err.(*MyError))
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}
	fmt.Println("\n==================================================================\n")

	err2 := MyError{Message: "2: エラーが発生しました。", ErrCode: 1234}
	fmt.Println("err2 =>", err2)
	fmt.Println("reflect.TypeOf(err2) =>", reflect.TypeOf(err2))
	fmt.Println("err2.Message =>", err2.Message)
	fmt.Println("err2.ErrCode =>", err2.ErrCode)
	fmt.Println("\n==================================================================\n")

	err3 := &MyError{Message: "3: エラーが発生しました。", ErrCode: 1234}
	fmt.Println("err3 =>", err3)
	fmt.Println("reflect.TypeOf(err3) =>", reflect.TypeOf(err3))
	fmt.Println("err3.Message =>", err3.Message)
	fmt.Println("err3.ErrCode =>", err3.ErrCode)
	fmt.Println("\n==================================================================\n")
}

// err => 1: エラーが発生しました。
// reflect.TypeOf(err) => *main.MyError
// err.Error() => 1: エラーが発生しました。
// err.(*MyError) => 1: エラーが発生しました。
// 1234

// ==================================================================

// err2 => {2: エラーが発生しました。 1234}
// reflect.TypeOf(err2) => main.MyError
// err2.Message => 2: エラーが発生しました。
// err2.ErrCode => 1234

// ==================================================================

// err3 => 3: エラーが発生しました。
// reflect.TypeOf(err3) => *main.MyError
// err3.Message => 3: エラーが発生しました。
// err3.ErrCode => 1234

