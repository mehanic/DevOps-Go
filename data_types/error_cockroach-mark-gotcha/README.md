This Go program demonstrates how the `errors` package from CockroachDB works in terms of error wrapping, marking, and matching. Specifically, the program demonstrates how `errors.Is` works when checking for error equality, including cases where errors are marked or wrapped. Let's go through the program step by step and explain the rules and behavior:

### **1. Defining a Custom Error Type**

```go
type NotFoundError struct {
	message string
}

func (e *NotFoundError) Error() string {
	return e.message
}
```

- The custom error type `NotFoundError` is defined, which contains a `message` field and implements the `Error()` method, which returns the error message. This makes `NotFoundError` a valid error type that can be used in Go.

### **2. Creating the Errors**

```go
err1 := &NotFoundError{"object not found"}
err2 := &NotFoundError{"object not found"}
```

- Two instances of `NotFoundError` are created: `err1` and `err2`. Both have the same message: `"object not found"`.
- Even though `err1` and `err2` have the same message, they are **two different instances** of the `NotFoundError` struct, which means they are not the same error object in memory.

### **3. Marking and Wrapping Errors**

```go
err := io.ErrUnexpectedEOF
err = errors.Mark(err, err1) // Помечаем ошибку как err1.
err = errors.Wrap(err, "something other happened")
```

- The error `err` is initially set to `io.ErrUnexpectedEOF`, a predefined error from the `io` package, which indicates an unexpected end-of-file error.
- The `errors.Mark` function is used to **mark** `err` with `err1`, meaning that `err` is now associated with `err1` as additional context.
- The `errors.Wrap` function wraps `err` with the message `"something other happened"`, adding more context to the error. This creates a chain of errors:
  - `err` now consists of:
    - `io.ErrUnexpectedEOF` as the original error,
    - marked with `err1` (of type `NotFoundError`),
    - and wrapped with the message `"something other happened"`.

### **4. Using `errors.Is` to Check Error Matching**

#### **Checking if `err` is `err1`**

```go
if errors.Is(err, err1) { // Ожидаемо true.
    fmt.Println("err is err1")
}
```

- **`errors.Is`** is used to check if `err` (which is now a wrapped and marked error) matches `err1`.
- Since `err` was marked with `err1` using `errors.Mark`, `errors.Is` will return `true` because it checks if the error is the same as `err1` or if it's wrapped by `err1`.
- The program prints `"err is err1"`.

#### **Checking if `err` is `err2`**

```go
if errors.Is(err, err2) { // Не очень ожидаемо true.
    fmt.Println("err is err2")
}
```

- Here, `errors.Is` checks if `err` matches `err2`.
- **Unexpected Behavior**: Despite `err2` being a different instance of `NotFoundError` (with the same message), `errors.Is` returns `true` because `errors.Is` is not checking for strict error equality. It only checks if the errors are of the same type or if one error is wrapped by another, even if they are different instances of the same type.
- This behavior is due to how `errors.Is` works with wrapped and marked errors. It compares the underlying error type and checks for equality across the error chain. Since `err2` is the same type as `err1` and was marked with `err1`, `errors.Is` considers `err2` a match.
- The program prints `"err is err2"`.

#### **Checking if `err1` is `err2`**

```go
if errors.Is(err1, err2) { // Ещё менее ожидаемо true.
    fmt.Println("err1 is err2")
}
```

- This checks if `err1` is considered equal to `err2` using `errors.Is`.
- **Unexpected Behavior**: The program prints `"err1 is err2"` because `errors.Is` compares the error type and considers `err1` and `err2` the same type (both are `*NotFoundError`), even though they are different instances.
- Since `errors.Is` checks whether one error is of the same type as another error, it returns `true` because both `err1` and `err2` are of type `*NotFoundError`.

#### **Checking if `err1` is `err2` using the standard library's `stderrors.Is`**

```go
if stderrors.Is(err1, err2) { // Ожидаемо false.
    fmt.Println("err1 is err2")
}
```

- Here, the standard library's `stderrors.Is` (from the `errors` package) is used to check if `err1` is `err2`.
- **Expected Behavior**: The standard `errors.Is` function only checks for **exact error equality**. It will return `false` when comparing different instances of an error, even if the error message is the same.
- Since `err1` and `err2` are different instances of the same type, `stderrors.Is` returns `false`, and the program does **not** print anything.

### **Key Rules and Concepts**

1. **Error Marking (`errors.Mark`)**:
   - Marking an error associates it with another error, allowing you to track that error type within the wrapped error chain.
   - `errors.Is` will match errors that are marked with a specific error type, even if the instances are different, as long as they share the same error type.

2. **Error Wrapping (`errors.Wrap`)**:
   - Wrapping an error adds context to the error while preserving the original error. `errors.Is` will still check if the wrapped error is of the same type or is one of the errors in the chain.

3. **`errors.Is` Behavior**:
   - `errors.Is` does not compare errors by their exact memory references but instead checks if the error is of the same type or if one error is wrapped by another.
   - This allows `errors.Is` to match errors of the same type, even if they are different instances, as long as one is marked or wrapped with the other.

4. **Standard Library `errors.Is`**:
   - The standard `errors.Is` function only checks for exact equality between errors and will return `false` if they are different instances, even if their messages are the same.

### **Conclusion**

- **Marking** and **wrapping** errors allow you to add context to errors and make it possible to match errors using `errors.Is` even when they are different instances of the same error type.
- `errors.Is` compares error types and checks if one error is wrapped by another, which leads to behavior where errors with the same message but different instances can be considered equal.
- The standard library's `errors.Is` function is stricter, checking only for exact error matches, which is why it returns `false` when comparing `err1` and `err2` in this case.