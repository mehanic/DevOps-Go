Let's break down the code and explain how the error wrapping works, particularly focusing on the `errors.Is` and `errors.Cause` methods, along with how standard errors (`fmt.Errorf`) and the `github.com/pkg/errors` package are used.

### Overview

In this code, the `wrappedErr` is set to `os.ErrNotExist`, a standard error defined in the Go `os` package. The code demonstrates various ways of wrapping errors, using both the Go standard library's `fmt.Errorf` and the `github.com/pkg/errors` package.

#### 1. **Error Wrapping with `fmt.Errorf` and `%w`**
`fmt.Errorf` with the `%w` specifier is used to **wrap** an error, allowing you to chain errors while preserving the original error (`wrappedErr` in this case). When wrapping errors, we can use the `errors.Is` and `errors.Cause` functions to inspect the error chain.

#### 2. **Error Wrapping with `errors.Wrap` (from `github.com/pkg/errors`)**
The `errors.Wrap` function from the `github.com/pkg/errors` package is another way to wrap errors. It preserves the original error while adding context, and it also provides functionality for stack traces (not used in this code).

### Code Breakdown

#### 1. **Case 1: "only std errors"**

```go
{
    title: "only std errors",
    err:   fmt.Errorf("msg 2: %w", fmt.Errorf("msg 1: %w", wrappedErr)),
}
```

- **Explanation**:
  - The `fmt.Errorf("msg 2: %w", fmt.Errorf("msg 1: %w", wrappedErr))` wraps `wrappedErr` twice:
    - The first wrap: `fmt.Errorf("msg 1: %w", wrappedErr)` wraps `wrappedErr` with a message `"msg 1"`.
    - The second wrap: `fmt.Errorf("msg 2: %w", ...)` wraps the previous error with an additional message `"msg 2"`.
  - The final error message will be: `"msg 2: msg 1: not exists"`.

- **Behavior**:
  - `errors.Is(c.err, wrappedErr)` checks whether the final error (`c.err`) contains `wrappedErr`. This will return `true` because `wrappedErr` is the original error in the chain.
  - `errors.Cause(c.err) == wrappedErr` retrieves the **root cause** of the error (the original error), and it will also return `true`.

#### 2. **Case 2: "only pkg/errors errors"**

```go
{
    title: "only pkg/errors errors",
    err:   errors.Wrap(errors.Wrap(wrappedErr, "msg 1"), " msg 2"),
}
```

- **Explanation**:
  - The `errors.Wrap(wrappedErr, "msg 1")` wraps `wrappedErr` with the message `"msg 1"`, and then `errors.Wrap(..., " msg 2")` wraps that result with the message `"msg 2"`.
  - The final error message will be: `"msg 2: msg 1: not found"`.

- **Behavior**:
  - `errors.Is(c.err, wrappedErr)` checks if `wrappedErr` is part of the error chain. It will return `true` because `wrappedErr` is the underlying error.
  - `errors.Cause(c.err) == wrappedErr` will also return `true` because `errors.Cause` retrieves the original error from the error chain, and `wrappedErr` is the root cause.

#### 3. **Case 3: "combined 1"**

```go
{
    title: "combined 1",
    err:   errors.Wrap(fmt.Errorf("msg 1: %w", wrappedErr), "msg 2"),
}
```

- **Explanation**:
  - The `fmt.Errorf("msg 1: %w", wrappedErr)` wraps `wrappedErr` with the message `"msg 1"`.
  - Then, `errors.Wrap(..., "msg 2")` wraps the result of the previous operation with `"msg 2"`.
  - The final error message will be: `"msg 2: msg 1: not found"`.

- **Behavior**:
  - `errors.Is(c.err, wrappedErr)` checks if `wrappedErr` is part of the error chain. It will return `true`.
  - `errors.Cause(c.err) == wrappedErr` will also return `true` because `wrappedErr` is the root cause.

#### 4. **Case 4: "combined 2"**

```go
{
    title: "combined 2",
    err:   fmt.Errorf("msg 2: %w", errors.Wrap(wrappedErr, "msg 1")),
}
```

- **Explanation**:
  - The `errors.Wrap(wrappedErr, "msg 1")` wraps `wrappedErr` with the message `"msg 1"`.
  - Then, `fmt.Errorf("msg 2: %w", ...)` wraps the result of the previous operation with `"msg 2"`.
  - The final error message will be: `"msg 2: msg 1: not found"`.

- **Behavior**:
  - `errors.Is(c.err, wrappedErr)` checks if `wrappedErr` is part of the error chain. It will return `true`.
  - `errors.Cause(c.err) == wrappedErr` will also return `true` because `wrappedErr` is the root cause.

### Explanation of `errors.Is` and `errors.Cause`

- **`errors.Is(c.err, wrappedErr)`**:
  - This function checks if the `wrappedErr` is present in the error chain, no matter how many times the error has been wrapped.
  - It returns `true` if `wrappedErr` is part of the chain.

- **`errors.Cause(c.err)`**:
  - This function retrieves the **root cause** of the error chain, i.e., the original error (`wrappedErr` in this case).
  - If the error is wrapped multiple times, `errors.Cause` will return the original error before any wrapping occurred.
  - It will return `wrappedErr` when called on any of the cases.

### Summary

- **`fmt.Errorf`** and **`errors.Wrap`** are both used to wrap errors in Go, but `fmt.Errorf` with `%w` is a standard library feature, and `errors.Wrap` comes from the third-party `github.com/pkg/errors` package.
- Both `fmt.Errorf` and `errors.Wrap` preserve the underlying error, which allows us to chain errors and add additional context.
- **`errors.Is`** helps to check if a particular error is part of an error chain, and **`errors.Cause`** retrieves the root cause of the error.

### Output:

The output of the program will be as follows:

```
only std errors
        errors.Is: true
        errors.Cause: true

only pkg/errors errors
        errors.Is: true
        errors.Cause: true

combined 1
        errors.Is: true
        errors.Cause: true

combined 2
        errors.Is: true
        errors.Cause: true
```

In each case, both `errors.Is` and `errors.Cause` return `true`, as they correctly identify the original error (`wrappedErr`) regardless of how the error is wrapped.