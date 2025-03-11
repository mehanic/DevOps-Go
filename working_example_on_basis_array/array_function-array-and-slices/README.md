This Go program demonstrates concepts related to **arrays** and **slices**, showing the differences between them, as well as how they behave when passed into functions. Here's a detailed breakdown of how the program works and the rules it illustrates:

### **1. Arrays in Go:**
- **Arrays are fixed in size**. Once an array's size is defined, it cannot change. Arrays in Go are value types, meaning they are copied when passed to functions or assigned to another variable.

#### Code Explanation:
```go
var threeNillIntegers [3]int
fmt.Println(threeNillIntegers)
```
- `threeNillIntegers` is an array of size 3, initialized with default values (`0` in the case of integers). The output will be `[0 0 0]`.

```go
var threeIntegers = [3]int{1, 2, 3}
fmt.Println(threeIntegers)
```
- `threeIntegers` is an array of size 3, initialized with specific values `{1, 2, 3}`. The output will be `[1 2 3]`.

```go
var integers = [...]int{1, 2, 3, 4, 5}
fmt.Println(fmt.Sprintf(`Array %d. Size %d`, integers, len(integers)))
```
- `integers` is an array of size 5, defined using the ellipsis `...` which allows the compiler to automatically determine the size. The output will be:
  ```
  Array [1 2 3 4 5]. Size 5
  ```

#### Passing Arrays to Functions (By Value):
```go
num := [...]int{5, 6, 7, 8, 8}
fmt.Println("before passing to function ", num)
changeLocal(num) //num is passed by value
fmt.Println("after passing to function ", num)
```
- **Arrays are passed by value** in Go, meaning that when an array is passed to a function, a **copy** of the array is passed. The original array remains unaffected by changes made to the copy.
- In the function `changeLocal`, the first element of the array is set to `55`, but the change does not affect the original `num` array in `main`.

Output:
```
before passing to function  [5 6 7 8 8]
inside function  [55 6 7 8 8]
after passing to function  [5 6 7 8 8]
```

---

### **2. Slices in Go:**
- **Slices are more flexible than arrays**. They are dynamically sized and can grow or shrink as needed.
- **Slices are references to an underlying array**. When a slice is passed into a function, any changes made to the slice will affect the original array.

#### Code Explanation:
```go
arr := [5]int{1, 2, 3, 4, 5}
var sli []int = arr[1:4]
fmt.Println(sli)
```
- Here, `sli` is a slice of the array `arr`, starting from index 1 and ending at index 4 (excluding index 4). The slice `sli` will contain the elements `[2 3 4]`.

```go
slic := []int{6, 7, 8}
fmt.Println(slic)
```
- `slic` is a slice with elements `{6, 7, 8}`. It doesn't reference an array but is created directly as a slice.

```go
darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
dslice := darr[2:5]
fmt.Println("array before", darr)
for i := range dslice {
    dslice[i]++
}
fmt.Println("array after", darr)
```
- Here, `dslice` is a slice of the array `darr` from index 2 to 4. Since slices are references to the original array, modifying `dslice` will modify the underlying `darr`.
- The output shows that the changes made to `dslice` affect `darr`:
  ```
  array before [57 89 90 82 100 78 67 69 59]
  array after [57 89 91 83 101 78 67 69 59]
  ```

#### Appending to Slices:
```go
veggies := []string{"potatoes", "tomatoes", "brinjal"}
fruits := []string{"oranges", "apples"}
food := append(veggies, fruits...)
fmt.Println("food:", food)
```
- The `append` function is used to concatenate two slices. `fruits...` uses the **ellipsis** to unpack the `fruits` slice and append each element to the `veggies` slice.
- The resulting `food` slice will be `["potatoes", "tomatoes", "brinjal", "oranges", "apples"]`.

---

### **3. Slices Are Passed By Reference:**
- Since slices are references to arrays, changes made to a slice inside a function will affect the original array or slice outside the function.

```go
nos := []int{8, 7, 6}
fmt.Println("slice before function call", nos)
subtactOne(nos)                               //function modifies the slice
fmt.Println("slice after function call", nos) //modifications are visible outside
```
- The `subtactOne` function subtracts `2` from each element of the `nos` slice.
- Since slices are passed by reference, the changes made inside the function will be reflected outside.

Output:
```
slice before function call [8 7 6]
slice after function call [6 5 4]
```

---

### **Summary of Key Concepts:**
1. **Arrays**:
   - Fixed size.
   - Passed by **value** (copy), meaning changes in a function do not affect the original array.
   - Declared with a fixed size or inferred size using `[...]`.

2. **Slices**:
   - Dynamic size, can grow and shrink.
   - Passed by **reference** (pointer), meaning changes in a function **affect** the original slice and its underlying array.
   - Created using `[]`, and can be sliced from an array or another slice.

3. **Key Operations**:
   - **Slices** can be appended using `append()`, which returns a new slice.
   - **Array passing** creates a copy, while **slice passing** refers to the original slice.

### **Function Breakdown:**
- `changeLocal(num [5]int)` demonstrates that **arrays are passed by value**.
- `subtactOne(numbers []int)` demonstrates that **slices are passed by reference** and modifying the slice inside the function reflects outside the function.

