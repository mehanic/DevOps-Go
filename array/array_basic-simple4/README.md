## **Array Manipulation in Go**

### **Algorithm Explanation**

1. **Array Initialization:**
   ```go
   numbers := [...]int{100, 12, 3, 53}
   ```
   - An array `numbers` is initialized with four integers: `100`, `12`, `3`, and `53`. The size of the array is inferred by the compiler using `...`.

2. **Modifying an Array Element:**
   ```go
   numbers[2] = 15
   ```
   - The element at index `2` (which was initially `3`) is updated to `15`.

### **Example Output**
```
[100 12 3 53]
[100 12 15 53]
```

---

### **Key Points**
- Arrays in Go are mutable, and their elements can be modified after initialization.