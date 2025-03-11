# **Working with Arrays in Go**

## **Algorithm Explanation**

1. **Array Initialization with Values:**
   ```go
   books := [...]string{"red", "orange", "yellow", "blue", "violet"}
   ```
   - An array `books` is initialized with 5 colors: `"red"`, `"orange"`, `"yellow"`, `"blue"`, and `"violet"`. The size is inferred using `...` based on the number of elements.

2. **Accessing an Array Element:**
   ```go
   fmt.Println(books[3])
   ```
   - The element at index `3` (which is `"blue"`) is accessed and printed.

## **Example Output**
```
[red orange yellow blue violet]
blue
```

---

### **Key Takeaways**
- Arrays in Go are zero-indexed, and their size can be inferred with `...`.
- Elements are accessed using the index, starting from 0.