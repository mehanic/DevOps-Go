This Go program demonstrates how to handle errors and determine whether they are temporary by checking error chains. It uses the `errors` package from CockroachDB, which provides advanced functionality for error wrapping, unwrapping, and inspecting error chains. Let's break down the key concepts and behavior of this program.

### **1. The `isTemporary` Interface**

```go
type isTemporary interface {
	Temporary() bool
}
```

- This defines a custom `isTemporary` interface that requires a method `Temporary()` which returns a `bool` indicating whether the error is considered temporary. 
- Any type that implements this interface is considered to be "temporary."
- For example, network errors like `net.DNSError` implement this interface, where `Temporary()` indicates whether the error is a temporary one that might resolve itself later (e.g., DNS lookup failure due to network issues).

### **2. The `IsTemporaryOnce` Function**

```go
// IsTemporaryOnce считает цепочку ошибок временной, если хотя бы одна
// из ошибок в ней была временной.
func IsTemporaryOnce(err error) bool {
	for c := err; c != nil; c = errors.UnwrapOnce(c) {
		e, ok := c.(isTemporary)
		if ok && e.Temporary() {
			return true
		}
	}
	return false
}
```

- **Purpose**: This function checks if **any error in the chain** is temporary. It traverses the error chain, starting from the given `err` and unwraps it once at a time.
  
- **How it works**:
  - `errors.UnwrapOnce(c)` is used to move to the next error in the chain. This function retrieves the next error that was wrapped, but **only once** (not recursively, which is different from `errors.UnwrapAll`).
  - For each error in the chain, it checks if the error implements the `isTemporary` interface and if `Temporary()` returns `true`.
  - If any error in the chain is temporary, the function returns `true`.
  - If no errors in the chain are temporary, it returns `false`.
  
- **Key Insight**: The function only checks **one level of the error chain at a time**.

### **3. The `IsTemporary` Function**

```go
// IsTemporary считает цепочку ошибок временной, только если
// корневая ошибка в ней была временной.
func IsTemporary(err error) bool {
	c := errors.UnwrapAll(err)
	e, ok := c.(isTemporary)
	return ok && e.Temporary()
}
```

- **Purpose**: This function checks if the **root error** (the first error in the chain) is temporary. Unlike `IsTemporaryOnce`, it doesn't check the entire error chain but only inspects the root error.

- **How it works**:
  - `errors.UnwrapAll(err)` is used to unwrap all the errors in the chain and return the innermost (root) error.
  - Then, it checks if this root error implements the `isTemporary` interface and if `Temporary()` returns `true`.
  
- **Key Insight**: This function focuses solely on the **root error**, not any other errors in the chain.

### **4. The `main` Function**

```go
func main() {
	dnsErr := &net.DNSError{IsTemporary: true} // Произошла сетевая ошибка.

	err := fmt.Errorf("second wrap: %w",
		fmt.Errorf("first wrap: %w", dnsErr))

	fmt.Printf("is temporary: %t\n", IsTemporaryOnce(err))     // is temporary: true
	fmt.Printf("is temporary at root: %t\n", IsTemporary(err)) // is temporary at root: true
}
```

- **Creating a DNS Error**:
  - `dnsErr` is created as a `net.DNSError` with `IsTemporary: true`. This simulates a network-related error that is temporary.
  
- **Wrapping the Error**:
  - `err` is a wrapped error using `fmt.Errorf`. The `dnsErr` is wrapped twice in the error chain:
    - The `dnsErr` is wrapped by `"first wrap: %w"`.
    - The result of the first wrap is wrapped again by `"second wrap: %w"`.
  
- **Checking if Any Error in the Chain is Temporary**:
  - `IsTemporaryOnce(err)` checks if **any error in the chain** is temporary. Since `dnsErr` is temporary, the result is `true`.
  - **Output**: `is temporary: true`

- **Checking if the Root Error is Temporary**:
  - `IsTemporary(err)` checks if the **root error** in the chain (`dnsErr`) is temporary. Since `dnsErr` is marked as temporary, the result is also `true`.
  - **Output**: `is temporary at root: true`

### **Key Differences Between `IsTemporaryOnce` and `IsTemporary`**

- **`IsTemporaryOnce`**:
  - Traverses the error chain and returns `true` if **any error in the chain** is temporary.
  - It **unwraps the error chain only once** at each step.
  - Returns `true` if any error in the chain (including the root) is temporary.
  
- **`IsTemporary`**:
  - Focuses only on the **root error** and returns `true` if the root error is temporary.
  - It **unwraps all** the errors in the chain and only checks the root error.

### **Conclusion**

- The program demonstrates how to work with **wrapped errors** in Go and how to check if an error (or any error in the chain) is **temporary**.
- The main difference between the `IsTemporaryOnce` and `IsTemporary` functions lies in whether they check the entire chain or just the root error. `IsTemporaryOnce` will return `true` if any error in the chain is temporary, while `IsTemporary` checks if the root error is temporary.
- **`errors.UnwrapOnce`** unpacks the error chain one level at a time, while **`errors.UnwrapAll`** retrieves the root error after unwrapping all levels.
