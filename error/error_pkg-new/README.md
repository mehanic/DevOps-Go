Let's break down the code and explain the rules step by step.

### Code Explanation:

```go
package main

import (
	stderrors "errors"
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := errors.New("error happened")
	fmt.Printf("%+v", err)

	sErr := stderrors.New("error happened")
	fmt.Printf("\n---\n%+v", sErr)
}
```

### Key Concepts:

1. **`errors.New` from the `github.com/pkg/errors` package:**
   - This is the function from the `pkg/errors` library that creates a new error with the provided message. It is a more flexible version of the `errors.New` function in the Go standard library. It allows you to add stack traces and additional context to errors if needed.
   - `fmt.Printf("%+v", err)` prints the error, including the stack trace if the error is wrapped with a stack (but since no stack is added here, it will print just the message).

2. **`errors.New` from the `errors` package (Go's standard library):**
   - This is the standard error creation function in Go. It creates a simple error without any additional context or stack traces.
   - `stderrors.New("error happened")` creates a new error with the message `"error happened"`, and `fmt.Printf("%+v", sErr)` will only print the message of this error without any stack trace or additional context.

### Code Breakdown:

#### **1. `errors.New` from `github.com/pkg/errors` (Enhanced Error with Stack Tracing)**

```go
err := errors.New("error happened")
fmt.Printf("%+v", err)
```

- `errors.New("error happened")` is the method from the `github.com/pkg/errors` package that creates an error with the message `"error happened"`.
  
- When you print this error with `fmt.Printf("%+v", err)`, the `%+v` format specifier is used. This specifier in the `github.com/pkg/errors` package prints the **error message along with the stack trace** (if a stack trace has been added to the error). However, in this case, since no stack trace was added to the error, it will just print the error message, which is `"error happened"`, but in a more extended form (if there were a stack trace, it would show).

#### **2. `stderrors.New` from the `errors` package (Standard Go Error)**

```go
sErr := stderrors.New("error happened")
fmt.Printf("\n---\n%+v", sErr)
```

- `stderrors.New("error happened")` is the method from Go's standard `errors` package. It creates a simple error with the message `"error happened"` but no stack trace or additional information.
  
- When you print this error with `fmt.Printf("%+v", sErr)`, the `%+v` specifier will only print the error message since the standard `errors.New` doesn't include any stack trace or extra information. Thus, the output will be just the message `"error happened"` without any additional details.

### Output:

```
error happened
---
error happened
```

### Key Points:

1. **Difference in Error Creation**:
   - `errors.New` from the `github.com/pkg/errors` package is more powerful because it supports adding stack traces and more context, which can be printed when the error is formatted with `%+v`.
   - The standard `errors.New` function in the Go standard library only creates basic errors with no additional context or stack trace.

2. **Stack Trace**:
   - When you use the `%+v` format specifier with errors created by `github.com/pkg/errors`, it prints the error message along with the stack trace. However, since you didn't add a stack trace here using functions like `errors.WithStack`, it will just print the error message.
   - The `%+v` specifier has no effect on errors created by `errors.New` from the standard library. It will just print the error message itself without any stack trace or additional information.

3. **Usage**:
   - **`github.com/pkg/errors`** is used when you need to wrap errors, attach additional context, or maintain stack traces for debugging.
   - **`errors.New`** from the standard library is simple and lightweight, suitable when you don't need to preserve stack traces or additional context.

### In Conclusion:

- The code compares the two error creation methods: one from the `github.com/pkg/errors` package and one from Go's standard `errors` package.
- Both errors are created with the same message `"error happened"`, but the way they are handled and displayed differs.
  - The error from `github.com/pkg/errors` can include stack traces, while the error from `errors.New` is more basic.