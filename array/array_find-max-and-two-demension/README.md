This Go program demonstrates **array handling** through two functions:  
1. `Array3()`: **Finds the largest number in a 1D array**  
2. `Array4()`: **Works with a 2D array (matrix)**  

---

## **1. Understanding `Array3()`: Finding the Maximum Number**
### **Code:**
```go
func Array3() {
    sayilar := [5]int{233, 24, 56, 87, 23}
    fmt.Println(sayilar)
    fmt.Println("-------------------")

    enBüyük := sayilar[0] // Initialize the max value as the first element

    for i := 0; i < len(sayilar); i++ { // Iterate over the array
        fmt.Println(sayilar[i])
        if sayilar[i] > enBüyük { // Compare each element with the max
            enBüyük = sayilar[i] // Update max if a larger value is found
        }
    }

    fmt.Println(enBüyük) // Print the max value
    fmt.Println("-------------------")
}
```

### **Explanation of Rules:**
1. **Array Initialization**
   ```go
   sayilar := [5]int{233, 24, 56, 87, 23}
   ```
   - Creates an array **`sayilar`** of length `5` with predefined values.
   - **Arrays in Go have a fixed length**.

2. **Iterating Over the Array**
   ```go
   for i := 0; i < len(sayilar); i++ {
   ```
   - Uses a `for` loop to iterate over **all elements**.
   - `len(sayilar)` returns the **length** of the array (`5` in this case).

3. **Finding the Maximum Value**
   ```go
   if sayilar[i] > enBüyük {
       enBüyük = sayilar[i]
   }
   ```
   - Compares the current element (`sayilar[i]`) with the **stored max value** (`enBüyük`).
   - Updates `enBüyük` **if the current element is greater**.
   - After the loop, `enBüyük` holds the **largest number in the array**.

### **Output:**
```
[233 24 56 87 23]
-------------------
233
24
56
87
23
233
-------------------
```
- The largest number (`233`) is **correctly identified**.

---

## **2. Understanding `Array4()`: Working with a 2D Array (Matrix)**
### **Code:**
```go
func Array4() {
    var sayilar [2][3]int // Declare a 2D array (2 rows, 3 columns)

    sayilar[0][0] = 12
    sayilar[0][1] = 31
    sayilar[0][2] = 98
    sayilar[1][0] = 23
    sayilar[1][1] = 54
    sayilar[1][2] = 13

    fmt.Println(sayilar) // Print the entire 2D array
}
```

### **Explanation of Rules:**
1. **2D Array Declaration**
   ```go
   var sayilar [2][3]int
   ```
   - Declares a **2D array** with **2 rows** and **3 columns**.
   - By default, all elements are **initialized to 0**.

2. **Assigning Values to the 2D Array**
   ```go
   sayilar[0][0] = 12
   sayilar[0][1] = 31
   sayilar[0][2] = 98
   sayilar[1][0] = 23
   sayilar[1][1] = 54
   sayilar[1][2] = 13
   ```
   - Assigns values to each index explicitly.

3. **Printing a 2D Array**
   ```go
   fmt.Println(sayilar)
   ```
   - Prints the **entire matrix structure**.

### **Output:**
```
[[12 31 98] [23 54 13]]
```
- The matrix is correctly structured with **two rows and three columns**.

---

## **3. Running `main()`**
```go
func main() {
    Array3()
    Array4()
}
```
- Calls `Array3()` → Finds and prints the **largest number** in `sayilar`.
- Calls `Array4()` → Prints a **2D array**.

---

## **Key Takeaways:**
| Concept | Explanation |
|---------|------------|
| **Arrays have fixed size** | Arrays in Go have a predefined size (e.g., `[5]int`). |
| **Looping over an array** | Use `for i := 0; i < len(arr); i++` to access each element. |
| **Finding the maximum number** | Compare each element with a stored max value and update if necessary. |
| **2D Arrays (Matrix)** | Use `[rows][cols]` format, e.g., `[2][3]int`. |
| **Printing 2D arrays** | Directly print the matrix using `fmt.Println(array)`. |

---

Would you like an **optimized version** of the `Array3()` function using a `for range` loop? 🚀