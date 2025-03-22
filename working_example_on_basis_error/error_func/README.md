Let's break down the Go program and understand how it works, focusing on error handling and the specific behavior of the `f1` and `f2` functions, as well as the custom error type.

### Code Explanation:

```go
package main

import (
    "fmt"
    "errors"
)

// f1 function which returns an integer result and an error
func f1(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New("can't work with 42") // Return error when arg is 42
    }
    return arg + 3, nil // Otherwise, return arg + 3 without error
}

// argError struct to represent a custom error type
type argError struct {
    arg int
    problem string
}

// The Error method makes argError implement the error interface
func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.problem) // Return a string representation of the error
}

// f2 function, similar to f1, but returns a custom error (argError)
func f2(arg int) (int, error) {
    if arg == 42 {
        return -1, &argError{arg, "can't work with it"} // Return custom error when arg is 42
    }
    return arg + 3, nil // Otherwise, return arg + 3 without error
}

func main() {
    // Call f1 for different inputs and handle errors
    for _, i := range []int{10, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e) // Print error message when f1 fails
        } else {
            fmt.Println("f1 work:", r) // Print result when f1 succeeds
        }
    }

    // Call f2 for different inputs and handle errors
    for _, i := range []int{10, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e) // Print error message when f2 fails
        } else {
            fmt.Println("f2 work:", r) // Print result when f2 succeeds
        }
    }

    // Call f2 with 42 and type assert the error into *argError
    _, e := f2(42)
    if ae, ok := e.(*argError); ok { // Type assert error to *argError type
        fmt.Println(ae.arg)   // Print the arg field of the custom error
        fmt.Println(ae.problem) // Print the problem field of the custom error
    }
}
```

### Detailed Explanation:

#### 1. **Function `f1`:**

- **Behavior:**
  - Takes an integer `arg` as input.
  - If `arg == 42`, it returns a generic error message (`errors.New("can't work with 42")`).
  - Otherwise, it returns the integer `arg + 3` with no error.

- **Output for inputs 10 and 42:**
  - When `arg = 10`, it returns `10 + 3 = 13` with no error.
  - When `arg = 42`, it returns the error `"can't work with 42"`.

#### 2. **Function `f2`:**

- **Behavior:**
  - Similar to `f1`, but uses a custom error type (`*argError`).
  - If `arg == 42`, it returns an error of type `*argError{arg, "can't work with it"}`.
  - Otherwise, it returns `arg + 3` with no error.

- **Custom error type `argError`:**
  - The `argError` struct holds the problematic argument (`arg`) and a descriptive message (`problem`).
  - The `Error()` method implements the `error` interface for `argError`, formatting it as a string like `"<arg> - <problem>"`.

- **Output for inputs 10 and 42:**
  - When `arg = 10`, it returns `10 + 3 = 13` with no error.
  - When `arg = 42`, it returns an error in the format `"42 - can't work with it"`.

#### 3. **Error Handling:**
- **For `f1`:**
  - The program checks if the error is non-nil. If so, it prints the error message (e.g., `"can't work with 42"`).
  - If no error, it prints the result (e.g., `"f1 work: 13"`).

- **For `f2`:**
  - The program also checks for errors. When `f2` returns a custom error (`*argError`), the error is printed in the format: `"42 - can't work with it"`.
  - For a custom error, we can type assert it (`e.(*argError)`) to access the fields of `argError`. This allows us to print the specific `arg` and `problem` fields of the error.

#### 4. **Type Assertion with `*argError`:**
- After calling `f2(42)`, the error is type-asserted to `*argError`.
  - This type assertion (`ae, ok := e.(*argError)`) checks whether the error is of type `*argError`. If it is, we can safely access the `arg` and `problem` fields.
  - In this case, the output would print:
    ```
    42
    can't work with it
    ```

### Output:

```
f1 work: 13
f1 failed: can't work with 42
f2 work: 13
f2 failed: 42 - can't work with it
42
can't work with it
```

### Key Points:
1. **Error Handling in Go:**
   - Go uses the `error` interface for error handling, and custom error types can implement this interface.
   - The `fmt.Errorf` function is used to format and return errors in a specific format.
   
2. **Custom Error Types:**
   - You can define custom error types in Go, such as the `argError` struct in this example. This allows you to provide more specific error messages and data associated with the error.

3. **Type Assertions:**
   - Go allows you to type-assert an error to a specific error type, enabling access to additional data fields and providing more detailed error information.

