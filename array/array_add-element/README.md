# **Pass by Reference in Slices – Go**

## **Overview**  
This program demonstrates how **slices in Go are passed by reference**, meaning that modifying a slice inside a function changes the original slice.

---

## **Algorithm Explanation**  

### **1. Define the `addElement` Function**  
The function **modifies the first element** of the slice:  
```go
func addElement(arr []int) {
    arr[0] = 123
}
```
- **Slices in Go are references to an underlying array.**  
- Modifying `arr[0]` inside `addElement()` changes the original slice.

### **2. Declare and Print the Slice in `main`**  
```go
arr := []int{1, 2, 3, 4}
fmt.Println(arr)  // Output: [1, 2, 3, 4]
```
- A slice `arr` is initialized with values `[1, 2, 3, 4]`.  

### **3. Call `addElement(arr)`**  
```go
addElement(arr)
```
- Since slices are references, `arr[0]` in `main` is **updated to 123**.

### **4. Print the Modified Slice**  
```go
fmt.Println(arr)  // Output: [123, 2, 3, 4]
```
- The change is reflected in `main` due to **pass-by-reference** behavior.

---

## **Example Output**  
```
[1, 2, 3, 4]
[123, 2, 3, 4]
```