package main

import (
	"fmt"
	"runtime/debug"
)

type WithStacktraceError struct {
	message    string
	stacktrace []byte
}

func (w *WithStacktraceError) Error() string {
	return w.message
}

func (w *WithStacktraceError) StackTrace() string {
	return string(w.stacktrace)
}

func doSomething() error {
	return &WithStacktraceError{
		message:    "something went wrong",
		stacktrace: debug.Stack(),
	}
}

func main() {
	if err := doSomething(); err != nil {
		type stackTracer interface {
			StackTrace() string
		}
		if st, ok := err.(stackTracer); ok {
			fmt.Printf("%s\n%s", err, st.StackTrace())
		}
	}
}

// something went wrong
// goroutine 1 [running]:
// runtime/debug.Stack()
// 	/home/mehanic/.gvm/gos/go1.23.0/src/runtime/debug/stack.go:26 +0x5e
// main.doSomething(...)
// 	/home/mehanic/structure/13.error/error-examples/03-go-errors-concept/error-with-stacktrace-opaque/main.go:24
// main.main()
// 	/home/mehanic/structure/13.error/error-examples/03-go-errors-concept/error-with-stacktrace-opaque/main.go:29 +0x18

