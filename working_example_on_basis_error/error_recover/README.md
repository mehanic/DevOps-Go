This Go program demonstrates how to handle runtime panics using the `panic` and `recover` mechanism. Here's an explanation of the rules and flow of the code:

### Key Concepts:

1. **Panic**:
   - The `panic` function is used in Go to abruptly stop the execution of a function or program. When `panic()` is called, the program begins to unwind, running deferred functions in the reverse order of their declaration. If a panic is not recovered (via `recover`), the program will terminate.

2. **Defer**:
   - The `defer` statement schedules a function to be executed right before the surrounding function returns. It's often used for resource cleanup (e.g., closing files, unlocking mutexes), but in this case, it's used for handling panics.

3. **Recover**:
   - The `recover` function is used to catch a panic that occurs within a function. It must be called inside a deferred function to catch the panic and stop the program from terminating. If there is no panic, `recover` returns `nil`. If there is a panic, `recover` returns the value passed to `panic()` and prevents the program from crashing.
   - Once a panic is recovered, the execution continues normally after the deferred function.

### Code Breakdown:

```go
package main

import "fmt"

func recovery() {
    if recover := recover(); recover != nil {
        fmt.Println("Recovered error:", recover)
    }
}

func someProcess() {
    defer recovery()  // Deferring recovery() to catch any panic
    fmt.Println("some process...")
    panic("Panic!")  // This causes the panic to happen
}

func main() {
    someProcess()  // Calling someProcess
    fmt.Println("After panic")  // This will be printed after panic is recovered
}
```

### Step-by-Step Explanation:

1. **`main()` Function**:
   - The `main()` function calls the `someProcess()` function.
   - After `someProcess()` returns (or when it finishes execution), the program attempts to print "After panic".
   - If `someProcess()` doesn't panic (or if the panic is recovered), "After panic" will be printed.

2. **`someProcess()` Function**:
   - `defer recovery()`: This line schedules the `recovery()` function to be executed when `someProcess()` finishes. `recovery()` will attempt to catch any panic that occurs inside `someProcess()`.
   - `fmt.Println("some process...")`: This line prints "some process...".
   - `panic("Panic!")`: This line causes the panic to happen. When `panic()` is called, the normal flow of execution is interrupted, and the program starts unwinding. However, since `recovery()` is deferred, it will be executed during the unwinding process.
   
3. **`recovery()` Function**:
   - The `recovery()` function contains a check for `recover()`.
   - `recover()` is called inside the deferred `recovery()` function to catch any panic. If a panic occurs, `recover()` will return the value passed to `panic()`. In this case, the panic value is the string `"Panic!"`.
   - If a panic is caught, it prints: `Recovered error: Panic!`. This prevents the program from terminating, and execution continues after `recovery()`.
   - If no panic occurs, `recover()` returns `nil`, and nothing happens in the `recovery()` function.

4. **After `recovery()` is Called**:
   - Once `recovery()` has executed, the program continues with the execution of the next statement after the `defer`. In this case, after the panic is recovered, the `main()` function prints `"After panic"`.
   
### Output Explanation:

The program outputs the following:

```
some process...
Recovered error: Panic!
After panic
```

1. **"some process..."**: This is printed by `fmt.Println` in `someProcess()`, before the panic occurs.
2. **"Recovered error: Panic!"**: This is printed by the `recovery()` function. `recover()` catches the panic, and the program continues executing.
3. **"After panic"**: This is printed by `main()`, which runs after the panic is recovered.

### Important Rules:

1. **Panic stops normal execution**: Once a panic occurs, the normal flow of the program is halted, and deferred functions start executing. In this case, the deferred `recovery()` function is called during the unwinding process.
   
2. **Recover prevents program termination**: The `recover()` function catches the panic and prevents the program from terminating. After the panic is recovered, the program continues executing the subsequent code.

3. **Defer is used to handle panics**: `defer` is essential in this program. It's used to ensure that `recovery()` is executed when the `someProcess()` function finishes, regardless of whether it finishes normally or due to a panic.

4. **Order of execution**: Deferred functions are executed in **LIFO (Last In, First Out)** order. So, in this case, `recovery()` is executed last when `someProcess()` finishes.

### Summary:

This Go program demonstrates the use of `panic`, `defer`, and `recover` to handle runtime errors in a controlled manner. When a panic occurs in `someProcess()`, the deferred `recovery()` function catches it, prints the error message, and allows the program to continue. Without `recover()`, the program would terminate immediately after the panic.