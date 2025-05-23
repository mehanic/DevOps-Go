This Go program demonstrates the use of the `errors.Is()` function, which is used to check if an error is equal to another error, potentially even if they are of different types. This function is part of the `errors` package introduced in Go 1.13 and is used to simplify error handling, allowing you to check for specific error values within chains of errors.

Let's break down the program step by step and explain the rules being demonstrated:

### 1. **Using `url.InvalidHostError` with `errors.Is()`**

```go
e1 := url.InvalidHostError("www.golang-courses.ru")
e2 := url.InvalidHostError("www.golang-courses.ru")
fmt.Println(errors.Is(e1, e2)) // true
```

- **Explanation**:
  - `url.InvalidHostError` creates a specific error type used for invalid host names in URLs.
  - Both `e1` and `e2` are created using `url.InvalidHostError` with the same string (`"www.golang-courses.ru"`).
  - **`errors.Is()`** is used to check if `e1` is the same as `e2`. The `errors.Is()` function returns `true` if the two errors are **the same error** or if `e2` is an error that wraps `e1`.
  
  - **Why `errors.Is(e1, e2)` returns `true`**:
    - In Go, when you use `url.InvalidHostError` with the same argument (`"www.golang-courses.ru"`), they are considered **equal** because the error type (`InvalidHostError`) and the error message are identical.
    - Therefore, `errors.Is()` will return `true` because both errors are of the same type and have the same message.

### 2. **Using `syscall.Errno` with `errors.Is()`**

```go
e3 := syscall.Errno(0xd)
e4 := syscall.Errno(0xd)
fmt.Println(errors.Is(e3, e4)) // true
```

- **Explanation**:
  - `syscall.Errno` is a type representing an error code in the form of a system error number.
  - Here, `e3` and `e4` are both created with the same error number (`0xd`), which represents the same system error.
  - **`errors.Is()`** is used to check if `e3` is the same as `e4`.

  - **Why `errors.Is(e3, e4)` returns `true`**:
    - When you create two `syscall.Errno` errors with the same numeric value (`0xd`), they are considered equal because they both represent the same system error.
    - Therefore, `errors.Is()` returns `true` because both errors are of the same type (`syscall.Errno`) and have the same value (`0xd`).

### Key Concepts and Rules Demonstrated:

1. **Equality of Errors**:
   - `errors.Is()` checks if two errors are **the same**. This includes checking if they are of the same type and hold the same value.
   - If two errors have the same type and the same underlying value (message or numeric value), `errors.Is()` will return `true`.
   - In the examples above, both pairs of errors (`e1` and `e2`, `e3` and `e4`) have the same type and value, so `errors.Is()` returns `true`.

2. **Errors of Same Type**:
   - If two errors are of the **same type** and contain the same value (e.g., the same string or number), `errors.Is()` will return `true`.
   - This rule is demonstrated by `url.InvalidHostError("www.golang-courses.ru")` and `syscall.Errno(0xd)`, where two errors with identical types and values return `true`.

3. **Errors from Different Types**:
   - `errors.Is()` can be used to compare errors of different types, and it checks if one error wraps another. However, in this example, both comparisons involve errors of the same type.

4. **Use of `errors.Is()`**:
   - `errors.Is()` is typically used to check for a specific error or a class of errors in error chains, allowing you to check if an error is wrapped by another error or if the error matches a certain value.
   - **Note**: In your examples, both errors are simple and not wrapped within other errors, so the comparison directly checks their values.

### Summary of Expected Output:

1. **First Comparison (`url.InvalidHostError`)**:
   - Since `e1` and `e2` are both `InvalidHostError` errors with the same message (`"www.golang-courses.ru"`), `errors.Is(e1, e2)` will return `true`.

2. **Second Comparison (`syscall.Errno`)**:
   - Since `e3` and `e4` are both `syscall.Errno` errors with the same error code (`0xd`), `errors.Is(e3, e4)` will return `true`.

### Final Output:

```
true
true
```

### Key Takeaways:

- **Equality of Errors**: `errors.Is()` can be used to compare errors of the same type and value. In this case, errors with the same string or numeric value will be considered equal.
- **`errors.Is()` Behavior**: `errors.Is()` checks both the type and value of the errors to determine equality. If the types and values match, it returns `true`. If they differ, it returns `false`.