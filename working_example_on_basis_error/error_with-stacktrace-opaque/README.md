This Go program demonstrates how to create and work with custom errors that include a stack trace. Let's break down the key elements and explain the rules behind them:

### Key Concepts

1. **Custom Error Type (`WithStacktraceError`):**
   - The program defines a custom error type `WithStacktraceError` that stores both an error message and a stack trace.
   - This custom error type implements the `Error()` method, which is required by the `error` interface in Go, allowing it to be used as an error.
   - Additionally, it defines a `StackTrace()` method to retrieve the stack trace as a string.

   Here's the definition:
   ```go
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
   ```

   - **`message`**: A string containing the error message (e.g., `"something went wrong"`).
   - **`stacktrace`**: A byte slice (`[]byte`) that contains the stack trace (captured by `debug.Stack()`).

2. **Capturing Stack Trace (`debug.Stack()`):**
   - The `debug.Stack()` function from the `runtime/debug` package captures the current goroutine's stack trace.
   - This stack trace will be returned as a byte slice and can be converted into a string for printing.

   ```go
   stacktrace: debug.Stack(),
   ```

3. **`doSomething()` Function:**
   - This function simulates an operation that could fail. In this case, it always returns an error (`WithStacktraceError`), capturing the current stack trace using `debug.Stack()`.
   - The `doSomething()` function returns an instance of `WithStacktraceError`, which includes both a message and the stack trace.

   ```go
   func doSomething() error {
       return &WithStacktraceError{
           message:    "something went wrong",
           stacktrace: debug.Stack(),
       }
   }
   ```

4. **Type Assertion (`stackTracer` interface):**
   - The `main` function checks if the error returned from `doSomething()` satisfies the `stackTracer` interface, which has a `StackTrace()` method.
   - A type assertion is used to check if the error implements the `StackTrace()` method and, if it does, prints the error message followed by the stack trace.

   ```go
   type stackTracer interface {
       StackTrace() string
   }

   if st, ok := err.(stackTracer); ok {
       fmt.Printf("%s\n%s", err, st.StackTrace())
   }
   ```

   - **`err.(stackTracer)`**: This is a type assertion that checks if the error `err` implements the `stackTracer` interface. If it does, `st` will hold the unwrapped error and you can call `st.StackTrace()` to get the stack trace.

5. **Output:**
   - If an error occurs, it will print both the error message and the stack trace.

   **Example Output:**
   ```
   something went wrong
   goroutine 1 [running]:
   runtime/debug.Stack()
   	/home/mehanic/.gvm/gos/go1.23.0/src/runtime/debug/stack.go:26 +0x5e
   main.doSomething(...)
   	/home/mehanic/structure/13.error/error-examples/03-go-errors-concept/error-with-stacktrace-opaque/main.go:24
   main.main()
   	/home/mehanic/structure/13.error/error-examples/03-go-errors-concept/error-with-stacktrace-opaque/main.go:29 +0x18
   ```

   - The first part of the output is the error message: `something went wrong`.
   - The second part is the stack trace, showing the functions that were called leading to the error. The stack trace includes:
     - The call to `runtime/debug.Stack()` in the Go runtime (`stack.go:26`).
     - The call to `main.doSomething()` in the user-defined code (`main.go:24`).
     - The call to `main.main()` in the user-defined code (`main.go:29`).
   - The stack trace also provides the exact file paths and line numbers, helping to debug where the error occurred in the program.

### Detailed Explanation of the Flow:

1. **Error Creation with Stack Trace:**
   - Inside `doSomething()`, a new error is created of type `WithStacktraceError`. The `debug.Stack()` function captures the stack trace when the error is created.
   
2. **Error Handling with Type Assertion:**
   - In the `main` function, the error is checked to see if it implements the `stackTracer` interface. If it does, it prints the error message and the stack trace.

3. **Printing the Stack Trace:**
   - When printing the stack trace, the program uses `fmt.Printf` to format the output. The `%s` format specifier is used to print both the error message and the stack trace.
   - The `StackTrace()` method of the custom error type returns the stack trace, which is printed in the output.

### Key Concepts:

1. **Custom Error Types:** 
   - Go allows you to define your own error types that can contain more information than just the error message. In this case, the `WithStacktraceError` type holds both the error message and a stack trace.
   
2. **Stack Trace in Go:** 
   - You can capture a stack trace using `debug.Stack()` and include it in custom errors to provide more detailed information for debugging.

3. **Type Assertion:** 
   - The program uses type assertion to check if the error implements the `StackTrace()` method, which is part of the `stackTracer` interface. This allows the program to call `StackTrace()` and print the stack trace.

4. **Error Handling with Additional Context:**
   - By wrapping errors with additional information (like a stack trace), it is easier to debug and track down the source of issues, especially in large and complex applications.

### Conclusion:

- This program illustrates how to create a custom error type that includes both an error message and a stack trace, which can be useful for debugging.
- The program shows how to capture a stack trace, wrap it with an error, and then later print it along with the error message using a custom method.
- By using type assertions, you can extract the stack trace from custom error types and print it to the console.