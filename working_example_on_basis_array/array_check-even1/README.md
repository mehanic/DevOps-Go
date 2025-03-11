## **Algorithm Explanation**

This Go program filters and returns even numbers from a given array using two functions:

1. **`CheckForEven(k int)`**:
   - This function checks if a number `k` is even. It returns `true` if `k` is divisible by 2 and `false` otherwise.

2. **`GetOnlyEven(arr []int)`**:
   - This function iterates through the input array `arr` and appends only the even numbers to a new slice `evens`.

3. **Main Function**:
   - The `main` function initializes an array `arr` and uses `GetOnlyEven` to filter out even numbers. It then prints the filtered array.

### **Example Output**
```
[2 4 6 8]
```

### **Key Points**
- The program filters even numbers using two functions, `CheckForEven` and `GetOnlyEven`. It showcases how to manipulate arrays and slices in Go efficiently.