Let's break down the code and explain what's going on with the error handling, including the use of `fmt.Errorf` and `errors.Errorf` in the context of error wrapping.

### Code Breakdown

#### 1. **First Block: `fmt.Errorf`**
```go
{
    err := fmt.Errorf("index.html: %v", ErrNotFound)
    fmt.Println(err, "|", errors.Is(err, ErrNotFound)) // index.html: not found | false
}
```

- **`fmt.Errorf("index.html: %v", ErrNotFound)`**:
  - The `fmt.Errorf` function is used to create a formatted error. Here, the `ErrNotFound` error is formatted as part of the message string `"index.html: not found"`.
  - The `%v` format specifier is used to format `ErrNotFound`. The default string representation of `ErrNotFound` is `"not found"`, so the error message becomes `"index.html: not found"`.
  - **Important:** `fmt.Errorf` does not wrap the original `ErrNotFound` error here. It only formats it into the string but does not preserve the original error as part of a chain. As a result, the original error is not available for detection via `errors.Is`.

- **`errors.Is(err, ErrNotFound)`**:
  - The `errors.Is` function is used to check if the `ErrNotFound` error is part of the error chain.
  - Since the `ErrNotFound` error is not wrapped inside the `err` error created by `fmt.Errorf`, `errors.Is` will return `false`, as the original `ErrNotFound` error is not part of the error chain.

- **Output**:
  - `index.html: not found | false`: The error message is `"index.html: not found"`, but `errors.Is` returns `false` because `ErrNotFound` is not part of the error chain.

#### 2. **Second Block: `errors.Errorf` from `github.com/pkg/errors`**
```go
{
    err := errors.Errorf("index.html: %v", ErrNotFound)
    fmt.Println(err, "|", errors.Is(err, ErrNotFound)) // index.html: not found | false
}
```

- **`errors.Errorf("index.html: %v", ErrNotFound)`**:
  - This is similar to `fmt.Errorf`, but it comes from the `github.com/pkg/errors` package. The difference between `fmt.Errorf` and `errors.Errorf` is that `errors.Errorf` can be used for advanced error handling features, like stack tracing and error wrapping.
  - In this case, `errors.Errorf` is used to create an error with the message `"index.html: not found"`, but it also doesn't wrap `ErrNotFound` inside the new error. It just formats the error message and does not add the original `ErrNotFound` error to the error chain.
  - **Important:** While `errors.Errorf` is capable of error wrapping and stack trace generation, in this case, it only formats the error and does not wrap `ErrNotFound`.

- **`errors.Is(err, ErrNotFound)`**:
  - Just like the first block, `errors.Is` is used to check if `ErrNotFound` is part of the error chain.
  - Since `ErrNotFound` is not wrapped by `errors.Errorf`, `errors.Is` will return `false`.

- **Output**:
  - `index.html: not found | false`: The error message is `"index.html: not found"`, but `errors.Is` returns `false` because `ErrNotFound` is not part of the error chain.

### Why `errors.Is` Returns `false`

The key issue here is how the errors are created and the **lack of error wrapping**:

1. **With `fmt.Errorf` and `errors.Errorf`:**
   - Both of these functions use a format specifier (`%v`) to format the error into a string. However, neither function uses the `%w` format specifier, which is specifically for wrapping errors in Go.
   - The `%v` specifier only formats the error message, and the original `ErrNotFound` error is not wrapped into the new error. This is why `errors.Is` cannot find `ErrNotFound` inside the error chain and returns `false`.
   
2. **Correct Way to Wrap Errors:**
   - To wrap an error properly so that `errors.Is` can detect it, you need to use the `%w` format specifier, which is designed specifically for wrapping an error while preserving the original error as part of the new error.
   
   For example:
   ```go
   err := fmt.Errorf("index.html: %w", ErrNotFound)
   ```

   This would create an error that correctly wraps `ErrNotFound`, and `errors.Is` would return `true`.

### Corrected Code with Error Wrapping

Here’s how you can modify the code to properly wrap `ErrNotFound` so that `errors.Is` can detect it:

```go
{
    err := fmt.Errorf("index.html: %w", ErrNotFound)
    fmt.Println(err, "|", errors.Is(err, ErrNotFound)) // index.html: not found | true
}

{
    err := errors.Errorf("index.html: %w", ErrNotFound)
    fmt.Println(err, "|", errors.Is(err, ErrNotFound)) // index.html: not found | true
}
```

- Now, the `%w` specifier is used to **wrap** `ErrNotFound` within the new error.
- **Output**:
  - `index.html: not found | true`: Now, `errors.Is` will correctly detect that `ErrNotFound` is part of the error chain because the error is wrapped properly.

### Summary

- **`fmt.Errorf`** and **`errors.Errorf`** are used to create errors with formatted messages.
- Both use the `%v` format specifier to format the error message, but **neither wraps the error** using the `%w` specifier, which is necessary for error wrapping.
- To wrap errors properly and make them detectable with `errors.Is`, you should use the `%w` specifier.
- Without wrapping (`%w`), `errors.Is` will return `false` because the original error is not part of the error chain.