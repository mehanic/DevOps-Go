# **Working with Arrays in Go**

## **Algorithm Explanation**

1. **Declaring an Array with Zero Values**
   ```go
   var numbers [10]int
   ```
   - An array `numbers` of size 10 is declared, and by default, all its elements are initialized to `0` (default value for `int` in Go).
   
2. **Initializing an Array with Specific Values**
   ```go
   elements := [5]int{1, 2, 3, 4, 5}
   ```
   - An array `elements` of size 5 is created and initialized with the values `{1, 2, 3, 4, 5}`.

3. **Using Ellipsis for Array Length Inference**
   ```go
   array := [...]int{1, 2, 4, 5, 6, 7, 8, 9, 2, 3, 4}
   ```
   - An array `array` is created with 11 elements. The length of the array is inferred automatically by Go from the number of elements in the initialization list using the ellipsis (`...`).

## **Example Output**
```
[0 0 0 0 0 0 0 0 0 0]
[1 2 3 4 5]
[1 2 4 5 6 7 8 9 2 3 4]
```

---

## **Key Points**
- Arrays in Go are fixed in size and can be initialized either explicitly or with length inference using `...`.
- Array elements are automatically initialized to default values if not explicitly set.