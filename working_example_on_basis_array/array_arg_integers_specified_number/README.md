# **Sum of N Elements in an Array**

## **Overview**  
This program reads `n` integers, stores them in an array, and then computes the sum of the first `n` elements of the array.

---

## **Algorithm Explanation**

### **1. Read the Input**  
```go
fmt.Scanf("%d", &n)
```
- The program reads an integer `n`, which defines how many numbers will be entered.

### **2. Initialize an Array**  
```go
arr := [100]int{}
```
- An array of size 100 is initialized to store integers.

### **3. Populate the Array**  
```go
for i := 0; i < n; i++ {
    fmt.Scanf("%d", &a)
    arr[i] = a
}
```
- The program loops `n` times, reading an integer each time and storing it in the array `arr`.

### **4. Calculate the Sum**  
```go
sumi := 0
for i := 0; i < n; i++ {
    sumi = sumi + arr[i]
}
```
- The sum of the first `n` elements is calculated by iterating over the array and adding each element to `sumi`.

### **5. Print the Results**  
```go
fmt.Println(arr)
fmt.Println(k)
fmt.Println(sumi)
```
- The array, its length, and the sum of its elements are printed.

---

## **Example Output**
```
Enter the number of elements: 3
Enter the elements: 1 2 3
[1 2 3 0 0 ...]
3
6
```