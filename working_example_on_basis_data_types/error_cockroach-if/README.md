This code demonstrates how to handle errors in Go using interfaces, custom error wrapping, and the CockroachDB `errors` package for conditional error inspection. Let's break down the rules and concepts in this code:

### **1. `isTemporary` Interface**

```go
type isTemporary interface {
	Temporary() bool
}
```

- **`isTemporary`** is an interface that defines a method `Temporary() bool`.
  - This interface is meant to identify errors that represent temporary issues, such as network errors or temporary file system errors, that may be retried later.
  - The method `Temporary()` should return `true` for errors that are temporary and `false` for others.

### **2. `IsTemporary` Function**

```go
func IsTemporary(err error) (any, bool) {
	e, ok := err.(isTemporary)
	return e, ok && e.Temporary() // Возвращаем не только флаг, но и ошибку.
}
```

- **`IsTemporary`** is a function that checks if an error implements the `isTemporary` interface and whether it’s considered temporary.
  - It first tries to type assert `err` into the `isTemporary` interface using `err.(isTemporary)`. If successful, it checks if the error is temporary using `e.Temporary()`.
  - The function returns two values:
    1. The error itself (`e` of type `isTemporary`), if it implements the interface.
    2. A boolean (`ok && e.Temporary()`) that indicates whether the error is temporary.
  - The function signature suggests it’s designed to return a more flexible type (returning `any`), which can be used to inspect the error further.

### **3. Creating an Example `dnsErr`**

```go
dnsErr := &net.DNSError{IsTemporary: true} // Произошла сетевая ошибка.
```

- **`dnsErr := &net.DNSError{IsTemporary: true}`** creates a simulated DNS error.
  - `&net.DNSError` is a pointer to the `net.DNSError` type, which is a predefined error type in Go used for DNS-related errors.
  - The `IsTemporary: true` field sets this error as a temporary one, indicating it’s a network issue that might be retried later.

### **4. Wrapping the Error**

```go
err := fmt.Errorf("second wrap: %w",
	fmt.Errorf("first wrap: %w", dnsErr))
```

- **Error Wrapping with `%w`:**
  - `fmt.Errorf("first wrap: %w", dnsErr)` creates a wrapped error with a message (`first wrap: ...`) and wraps the `dnsErr` error.
  - `fmt.Errorf("second wrap: %w", ...)` further wraps the first error with another message (`second wrap: ...`), effectively creating a chain of errors.
  - The `%w` format specifier is used in Go 1.13 and later for **error wrapping**. This allows the new error to include the original error, maintaining the context while still being able to unwrap it later.

### **5. Using `errors.If` from CockroachDB**

```go
e, ok := errors.If(err, IsTemporary)
```

- **`errors.If(err, IsTemporary)`**:
  - This function from the CockroachDB `errors` package is used to apply a condition (`IsTemporary`) to an error chain (`err`).
  - The `errors.If` function checks if the error chain contains an error that matches the `IsTemporary` condition.
  - It applies the `IsTemporary` function (which checks if an error is temporary) and returns:
    1. The first error (`e`) in the chain that matches the condition (in this case, the `dnsErr` error).
    2. A boolean (`ok`) that indicates if the condition was met. If `ok` is `true`, it means that the condition (`Temporary() == true`) was satisfied.

### **6. Outputting the Result**

```go
fmt.Printf("%T is temporary: %t", e, ok) // *net.DNSError is temporary: true
```

- **`fmt.Printf("%T is temporary: %t", e, ok)`**:
  - This prints the type of the error (`*net.DNSError`) and whether it's considered temporary (`true`).
  - The expected output is: `*net.DNSError is temporary: true`.

---

### **Summary of Key Concepts:**

1. **Custom Error Interfaces (`isTemporary`)**:
   - The `isTemporary` interface is used to define errors that have a `Temporary()` method. This is a way to categorize errors that can be retried due to their temporary nature.

2. **Error Wrapping (`%w`)**:
   - Go 1.13 and later support error wrapping using the `%w` format specifier in `fmt.Errorf()`. This allows chaining errors while preserving the original error context.

3. **Using `errors.If` from CockroachDB**:
   - The `errors.If` function is used to conditionally extract and check specific types of errors (e.g., those implementing the `isTemporary` interface). It simplifies the process of checking for specific error conditions within a chain of errors.

4. **Error Handling with Multiple Layers**:
   - The code wraps errors in multiple layers, making it possible to trace and handle errors at various levels while still providing useful context and preserving the original error. This is useful for handling complex error scenarios in Go.

### **How the Program Works:**

1. **Create a Temporary Error** (`dnsErr`).
2. **Wrap the Error** with two layers of messages (`first wrap` and `second wrap`).
3. **Check the Error Chain** for a temporary error using `errors.If` with the `IsTemporary` condition.
4. **Output the Results**, printing the error type and whether it is considered temporary.

This approach allows for flexible, layered error handling while retaining error context and making it easy to check specific error conditions.