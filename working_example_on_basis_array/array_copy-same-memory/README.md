## **Algorithm Explanation**

1. **Array Initialization**:
   - An integer slice `arr` is initialized with values 
   
   `[1, 2, 3, 4]`.

2. **Copying the Array**:
   - The slice `arr1` is assigned the value of `arr`. In Go, slices are reference types, meaning `arr1` refers to the same underlying array as `arr`.

3. **Modifying the Array**:
   - The first element of `arr` is updated to `123`. Since `arr1` references the same underlying array, it reflects the same change.

4. **Output**:
   - Both `arr` and `arr1` are printed, showing the modified value in both slices.

### **Example Output**
```
[123 2 3 4]
[123 2 3 4]
```

This demonstrates how slices in Go are reference types, and modifying one slice affects all references to the same underlying array.