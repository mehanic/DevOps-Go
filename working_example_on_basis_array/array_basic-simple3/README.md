## **Array in Go: Accessing Elements**

### **Algorithm Explanation**

1. **Array Initialization:**
   ```go
   numbers := [...]int{100, 23, 412, 23}
   ```
   - An array `numbers` is initialized with four integer values: `100`, `23`, `412`, and `23`. The size is inferred automatically using `...`.

2. **Accessing an Array Element:**
   ```go
   n := numbers[1]
   fmt.Println(n)
   ```
   - The element at index `1` (which is `23`) is accessed and stored in variable `n`, then printed.

### **Example Output**
```
[100 23 412 23]
23
```

---

### **Key Points**
- Go arrays are zero-indexed.
- You can access array elements using their indices.