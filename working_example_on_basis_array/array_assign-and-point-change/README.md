# **Working with Arrays in Go**

## **Overview**
This program demonstrates how to initialize an array, assign a value to an array element, and use the `len()` function to get the length of the array.

---

## **Algorithm Explanation**

### **1. Array Initialization**
```go
var x [5]int
fmt.Println(x)
```
- The array `x` is initialized with 5 elements, all set to `0` by default.

### **2. Modify Array Element**
```go
x[3] = 100
fmt.Println(x)
```
- The element at index `3` is set to `100`.

### **3. Length of Array**
```go
fmt.Println(len(x))
```
- The program prints the length of the array, which is `5`.

---

## **Example Output**
```
[0 0 0 0 0]
[0 0 0 100 0]
5
```