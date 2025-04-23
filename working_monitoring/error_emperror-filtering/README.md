This Go program demonstrates how to handle errors with the **`emperror`** package, a utility for error handling and filtering in Go. Here's a detailed explanation of the rules, components, and behavior of the program.

### Overview of Key Concepts:

1. **`emperror.ErrorHandlerFunc`**:
   - This is a type that implements the `emperror.ErrorHandler` interface, which defines how errors should be handled when they occur.
   - In the program, the error handler is defined by the `errHandler` variable, which prints out the error when called.

2. **`emperror.ErrorMatcher`**:
   - This type is used to match or filter certain errors based on the provided condition.
   - In the program, the `errMatcher` is responsible for checking if an error should be discarded (i.e., ignored) based on whether it matches one of the errors in the `errsToSkip` list.

3. **`emperror.WithFilter`**:
   - The `WithFilter` function combines the error handler (`errHandler`) with a custom error matcher (`errMatcher`).
   - This allows you to conditionally handle or skip errors based on the matching logic defined in `errMatcher`.

### Code Explanation:

```go
var (
	err1       = errors.New("error 1")  // Define error 1
	err2       = errors.New("error 2")  // Define error 2
	errsToSkip = []error{err1, err2}    // List of errors to skip

	// Define an error handler function that prints the error when called
	errHandler = emperror.ErrorHandlerFunc(func(err error) {
		if err == nil {
			return // Skip if the error is nil
		}
		fmt.Printf("error handler called: %v\n", err) // Print the error
	})

	// Define a custom error matcher that checks if the error should be skipped
	errMatcher = emperror.ErrorMatcher(func(err error) bool {
		for i := range errsToSkip {
			if needDiscard := errors.Is(err, errsToSkip[i]); needDiscard {
				return true // Skip errors that match any in the `errsToSkip` list
			}
		}
		return false // Otherwise, don't skip the error
	})
)
```

### How the Program Works:

1. **Error Handler (`errHandler`)**:
   - This is a simple function that prints out the error when it's called. It's wrapped using `emperror.ErrorHandlerFunc` to make it conform to the `emperror.ErrorHandler` interface.
   - If the error is `nil`, the handler does nothing. Otherwise, it prints the error message.

2. **Error Matcher (`errMatcher`)**:
   - The matcher checks if the provided error is one of the errors listed in `errsToSkip` (i.e., `err1` and `err2`).
   - If the error matches one of the errors to be skipped, the function returns `true`, indicating the error should be discarded by the handler.
   - If it doesn't match any of the errors in the list, it returns `false`, meaning the error should be passed to the handler.

3. **Combining the Handler and Matcher (`WithFilter`)**:
   - The `emperror.WithFilter` function is used to combine the error handler and the matcher. It ensures that the error handler is only called if the matcher returns `false` (i.e., the error is **not** one of the errors to skip).
   - This is essentially a conditional handler setup: the handler is active only for errors that do not match the `errsToSkip` list.

### Main Function:

```go
func main() {
	handler := emperror.WithFilter(errHandler, errMatcher)

	handler.Handle(err1)                        // This will not trigger the handler.
	handler.Handle(errors.New("unknown error")) // This will trigger the handler.
}
```

#### Steps in `main`:
1. **Handling `err1`**:
   - The `err1` is passed to the handler. Since `err1` is in the `errsToSkip` list, the `errMatcher` will return `true`, indicating that `err1` should be skipped.
   - As a result, the error handler will **not** be triggered for `err1`.

2. **Handling `unknown error`**:
   - A new error, `"unknown error"`, is created and passed to the handler.
   - Since this error is not in the `errsToSkip` list, the `errMatcher` will return `false`, indicating that this error should **not** be skipped.
   - Therefore, the error handler will be triggered and print out: `error handler called: unknown error`.

### Expected Output:
```
error handler called: unknown error
```

### Key Concepts Recap:
1. **Error Skipping**:
   - By using a custom error matcher (`errMatcher`), you can filter out certain errors from being handled. In this case, `err1` and `err2` are skipped, and only errors not in `errsToSkip` are passed to the handler.
   
2. **Error Handler**:
   - The error handler is responsible for defining what to do when an error occurs. In this case, the handler simply prints the error.

3. **Combination of Matcher and Handler**:
   - `emperror.WithFilter` combines the matcher and handler into one, so the handler only gets called for errors that don’t match the criteria defined by the matcher.

### Practical Use Case:
This pattern is useful when you want to handle some errors differently or skip certain types of errors in a more controlled manner. For instance:
- You may want to silently handle known, non-critical errors (like temporary network errors).
- You may want to log or alert on other types of errors that require attention (like unexpected or unknown errors).

This program demonstrates a flexible and scalable approach to error handling in Go using the `emperror` package.