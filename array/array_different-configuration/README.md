This Go program demonstrates **array initialization, indexing, length, and multi-dimensional arrays**. Let's go through the key rules governing this behavior.

---

## **1. Declaring and Initializing an Array**
```go
var a [5]int
fmt.Println("emp:", a)
```
- This declares an **array** `a` of **size 5** with **default zero-values**.
- Since `int` has a zero-value of `0`, all elements are initialized to `0`.
- Output:
  ```
  emp: [0 0 0 0 0]
  ```

### **Key Rule:**
- **Arrays in Go are fixed-size and zero-initialized.**
- `[N]T` means the array has `N` elements of type `T`.

---

## **2. Setting and Accessing Elements**
```go
a[4] = 100
fmt.Println("set a[4]=100:", a)
fmt.Println("a[4]:", a[4])
```
- This assigns `100` to **index `4`** in the array.
- Arrays in Go are **zero-indexed**, meaning:
  - `a[0]` is the first element.
  - `a[4]` is the fifth (last) element.
- Output:
  ```
  set a[4]=100: [0 0 0 0 100]
  a[4]: 100
  ```

### **Key Rule:**
- **Arrays are mutable**: You can change elements using `array[index] = value`.

---

## **3. Getting the Length of an Array**
```go
fmt.Println("len:", len(a))
```
- `len(a)` returns the **size** of the array.
- Since `a` was declared as `[5]int`, its length is **always 5**.
- Output:
  ```
  len: 5
  ```

### **Key Rule:**
- **Array length is fixed and retrieved with `len(array)`.**
- This is **not dynamic** like slices (which can grow or shrink).

---

## **4. Declaring and Initializing an Array with Values**
```go
b := [5]int{1, 2, 3, 4, 5}
fmt.Println("dcl:", b)
```
- This initializes an array **with predefined values**.
- The type `[5]int` ensures that the array has exactly **5 elements**.
- Output:
  ```
  dcl: [1 2 3 4 5]
  ```

### **Key Rule:**
- **You can initialize an array directly using `{values}`.**
- If the number of elements matches the size, Go automatically assigns them.

---

## **5. Multi-Dimensional Arrays**
```go
var twoD [2][3]int
for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
        twoD[i][j] = i + j
    }
}
fmt.Println("2d: ", twoD)
```
- Declares a **2x3 integer array** (`[2][3]int`), meaning:
  - **2 rows**
  - **3 columns per row**
- Loops through each index and sets `twoD[i][j] = i + j`.
- Output:
  ```
  2d:  [[0 1 2] [1 2 3]]
  ```

### **Breakdown of Array Values**
| `i` (Row) | `j` (Column) | `i + j` (Value) |
|-----------|-------------|---------------|
| 0         | 0           | 0 + 0 = 0     |
| 0         | 1           | 0 + 1 = 1     |
| 0         | 2           | 0 + 2 = 2     |
| 1         | 0           | 1 + 0 = 1     |
| 1         | 1           | 1 + 1 = 2     |
| 1         | 2           | 1 + 2 = 3     |

### **Key Rule:**
- **Multi-dimensional arrays are fixed in size (`[rows][cols]T`).**
- **Elements are accessed using `array[row][col]`.**

---

## **Key Takeaways**
| **Concept**           | **Code Example**                | **Explanation** |
|----------------------|-----------------------------|---------------|
| **Declare an array** | `var a [5]int`              | Creates an array of 5 integers (all initialized to `0`). |
| **Set an element**   | `a[4] = 100`                | Modifies index `4` to store `100`. |
| **Access an element** | `fmt.Println(a[4])`        | Retrieves value at index `4`. |
| **Get length**       | `len(a)`                    | Returns **fixed** size of array. |
| **Initialize with values** | `b := [5]int{1,2,3,4,5}` | Creates a **pre-filled** array. |
| **Multi-dimensional arrays** | `var twoD [2][3]int` | Creates a **2-row, 3-column** array. |

---

Would you like a **comparison between arrays and slices** in Go? 🚀