package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	Wrap()
	fmt.Println()
	Unwrap()
	fmt.Println()
	StackTrace()
}

// WrappedError demonstrates error wrapping and
// annotating an error
func WrappedError(e error) error {
	return errors.Wrap(e, "An error occurred in WrappedError")
}

// ErrorTyped is a error we can check against
type ErrorTyped struct {
	error
}

// Wrap shows what happens when we wrap an error
func Wrap() {
	e := errors.New("standard error")

	fmt.Println("Regular Error - ", WrappedError(e))

	fmt.Println("Typed Error - ", WrappedError(ErrorTyped{errors.New("typed error")}))

	fmt.Println("Nil -", WrappedError(nil))

}

//---

// Unwrap will unwrap an error and do
// type assertion to it
func Unwrap() {

	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")

	fmt.Println("wrapped error: ", err)

	// we can handle many error types
	switch errors.Cause(err).(type) {
	case ErrorTyped:
		fmt.Println("a typed error occurred: ", err)
	default:
		fmt.Println("an unknown error occurred")
	}
}

// StackTrace will print all the stack for
// the error
func StackTrace() {
	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")

	fmt.Printf("%+v\n", err)
}


// Regular Error -  An error occurred in WrappedError: standard error
// Typed Error -  An error occurred in WrappedError: typed error
// Nil - <nil>

// wrapped error:  wrapped: an error occurred
// a typed error occurred:  wrapped: an error occurred

// an error occurred
// wrapped
// main.StackTrace
// 	/home/mehanic/structure/13.error/error_wrap/errwrap.go:64
// main.main
// 	/home/mehanic/structure/13.error/error_wrap/errwrap.go:14
// runtime.main
// 	/home/mehanic/.gvm/gos/go1.23.0/src/runtime/proc.go:272
// runtime.goexit
// 	/home/mehanic/.gvm/gos/go1.23.0/src/runtime/asm_amd64.s:1700
