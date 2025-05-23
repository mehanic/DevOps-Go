Let's break down the code in your example and understand the rules, especially regarding how Go handles error checking.

### Code Explanation:

```go
package main

import (
	"fmt"
	"io"
)

func main() {
	// First error check
	if err := work(); nil != err {
		fmt.Println("error")
	}

	// Second error check
	if err := work(); err == nil {
		fmt.Println("no error")
	}
}

func work() error {
	return io.EOF
}
```

### Key Points to Understand:

1. **The `work` function**:
   - The function `work()` always returns `io.EOF`, which is a predefined error in the `io` package that represents the end-of-file error. In this case, the error is not `nil`; it's an actual error (`io.EOF`).

2. **First `if` statement**: 
   ```go
   if err := work(); nil != err {
       fmt.Println("error")
   }
   ```
   - This statement checks if the error returned by the `work()` function is **not** `nil`. Since `work()` returns `io.EOF` (which is not `nil`), the condition `nil != err` is `true`.
   - As a result, `"error"` is printed to the console.

3. **Second `if` statement**:
   ```go
   if err := work(); err == nil {
       fmt.Println("no error")
   }
   ```
   - This statement checks if the error returned by `work()` is `nil`. Since `work()` returns `io.EOF`, the error is **not** `nil`, meaning the condition `err == nil` is `false`.
   - As a result, `"no error"` is **not** printed to the console.

### Detailed Breakdown:

- **Error Checking in Go**:
  - In Go, it's common to check whether an error is `nil` or not after calling a function that returns an error. 
  - If the error is not `nil`, it typically means something went wrong, and you may want to handle it (e.g., print an error message, log it, or return from the function).
  
- **Condition `nil != err`**:
  - The condition `nil != err` is simply checking whether the error is not `nil`. If it’s not `nil` (meaning an error occurred), the code inside the block will execute.
  - In your case, `work()` returns `io.EOF`, so this condition evaluates to `true`, and `"error"` is printed.

- **Condition `err == nil`**:
  - The condition `err == nil` is checking whether the error is `nil` (no error occurred).
  - Since `work()` returns `io.EOF`, which is not `nil`, this condition evaluates to `false`, and the code inside this block doesn't execute (hence `"no error"` is not printed).

### Output:

```
error
```

### Conclusion:
- The first `if` statement evaluates to `true` because `err` is not `nil` (it's `io.EOF`), so it prints `"error"`.
- The second `if` statement evaluates to `false` because `err` is not `nil`, so it does **not** print `"no error"`.
  
This is how error handling is typically done in Go — checking if the returned error is `nil` or not to determine whether the operation was successful or not.