package main

import "fmt"

type error interface {
	Error() string
}

type appError struct {
	Err    error
	Custom string
	Field  int
}

type AppError interface {
	Error() string
}

// func (e *AppError) Error() string {

// 	return e.Err.Error()
// }

// func m() error {
// 	return &AppError{
// 		Err:    fmt.Errorf("my error"),
// 		Custom: "va",
// 		Field:  89,
// 	}
// }

func (e *appError) Error() string {
	return e.Error()
}

func main() {
	err := method1()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")

}

func method1() error {
	return method2()

}

func method2() error {
	return method3()
}

func method3() error {

	return fmt.Errorf("some error")
}

// some error
