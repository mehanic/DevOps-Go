## **Algorithm Explanation**

### **Step-by-Step Process**

1. **Array Initialization**:
   - An array `originalArray` of size 5 is initialized with the elements `[1, 2, 3, 4, 5]`.
   
2. **Length of Array**:
   - The length of `originalArray` is printed using the `len()` function, which outputs `5`.

3. **Copying the Array**:
   - The array is copied using `var copiedArray = originalArray`. This creates a new array with the same elements, and is printed.

4. **Modifying an Array**:
   - A function `modifyArray()` is defined that takes the array, an index, and a new value as parameters. 
   - This function modifies the value at the specified index in the array and returns the modified array.
   - The third element (index `2`) of `originalArray` is modified to `10`.

5. **Array Comparison**:
   - The original array (`originalArray`) is compared with the modified array (`modifiedArray`). Since the modification changes an element, the arrays are not equal, and the result is `false`.
   - The `originalArray` is compared with the copied array (`copiedArray`). Since no changes were made to the copied array, the comparison results in `true`.

### **Example Output**
```
Length of the array is: 5
Copied Array: [1 2 3 4 5]
Modified Array: [1 2 10 4 5]
Are original and modified arrays equal? false
Are original and copied arrays equal? true
```

This program demonstrates array manipulation in Go, including copying arrays, modifying specific elements, and comparing arrays for equality.