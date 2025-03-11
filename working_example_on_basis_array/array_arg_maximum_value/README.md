# **Find Maximum Element in an Array**

## **Overview**
This program reads `n` integers, stores them in an array, and then finds the maximum element in the array.

---

## **Algorithm Explanation**

### **1. Input Reading**  
```go
fmt.Scanf("%d", &n)
```
- The program first reads an integer `n`, representing the number of elements to input.

### **2. Populate the Array**  
```go
for i := 0; i < n; i++ {
    fmt.Scanf("%d", &a)
    arr[i] = a
}
```
- An array of size 100 is initialized, and the program reads `n` integers and stores them in the array.

### **3. Find Maximum Element**  
```go
maxi := arr[0]
for i := 0; i < n; i++ {
    if arr[i] > maxi {
        maxi = arr[i]
    }
}
```
- The program initializes the variable `maxi` to the first element of the array, and iterates through the array. If a larger element is found, `maxi` is updated to that element.

### **4. Output the Maximum Value**  
```go
fmt.Println(maxi)
```
- The maximum element found in the array is printed.

---

## **Example Output**
```
Enter the number of elements: 5
Enter the elements: 20 30 40 50 60
60
```