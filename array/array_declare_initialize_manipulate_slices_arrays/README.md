This Go program demonstrates the **rules and behavior of slices and arrays**, including their **initialization, comparison, and capacity management**. Let’s break it down into key concepts and rules.

---

## **1. Slices vs. Arrays in Go**
In Go, **slices** and **arrays** are distinct types:

- **Slices (`[]T`)**: Dynamically sized, references an underlying array.
- **Arrays (`[N]T`)**: Fixed-size, cannot grow dynamically.

### **Key Differences**
| Feature  | Array (`[N]T`) | Slice (`[]T`) |
|----------|---------------|---------------|
| Size | Fixed at compile time | Dynamic, can grow/shrink |
| Memory | Stored inline in struct | Points to an underlying array |
| Comparison | Can be compared if same size/type | Cannot be compared directly |
| Nil Check | Always initialized with zero values | Can be `nil` |

---

## **2. Slice Initialization and Behavior**

### **Rule 1: Declaring a slice without initialization results in a `nil` slice.**
```go
var a []int
log.Printf("a type=%T, a length=%d, a capacity=%d, a=%v, nil? %t\n", a, len(a), cap(a), a, a == nil)
```
- **`var a []int`** creates a `nil` slice with `len=0` and `cap=0`.
- **Comparison `a == nil`** returns **true**.

---

### **Rule 2: `new([]T)` returns a pointer to a nil slice, not an empty slice.**
```go
b := new([]int)
log.Printf("b type=%T, b length=%d, b capacity=%d, b=%v, nil? %t\n", b, len(*b), cap(*b), b, b == nil)
```
- `b` is a **pointer to a slice** (`*[]int`).
- `b` itself is **not `nil`**, but `*b` (dereferenced slice) **is a nil slice** (`len=0, cap=0`).

---

### **Rule 3: `make([]T, length, capacity)` creates a slice with an underlying array.**
```go
c := make([]int, 5)
log.Printf("c type=%T, c length=%d, c capacity=%d, c=%v, nil? %t\n", c, len(c), cap(c), c, c == nil)
```
- `make([]int, 5)` creates a slice with:
  - `len=5`, `cap=5`
  - **Not nil** (has an underlying array)

```go
d := make([]int, 5, 100)
log.Printf("d type=%T, d length=%d, d capacity=%d, d=%v, nil? %t\n", d, len(d), cap(d), d, d == nil)
```
- `make([]int, 5, 100)` creates a slice with:
  - `len=5`, `cap=100` (extra capacity for appending)

---

### **Rule 4: Slice literals initialize a slice directly.**
```go
g := []int{1, 3, 6}
log.Printf("g type=%T, g length=%d, g capacity=%d, g=%v, nil? %t\n", g, len(g), cap(g), g, g == nil)
```
- `g := []int{1, 3, 6}` creates a slice with:
  - `len=3`, `cap=3`
  - **Not nil**, since it has an underlying array.

---

### **Rule 5: An empty slice literal (`[]T{}`) is not `nil`.**
```go
h := []int{}
log.Printf("h type=%T, h length=%d, h capacity=%d, h=%v, nil? %t\n", h, len(h), cap(h), h, h == nil)
```
- `h := []int{}` initializes an **empty but non-nil slice**.
- `len=0`, `cap=0`, but `h != nil`.

---

## **3. Arrays in Go**
### **Rule 6: Arrays have a fixed length and are initialized with zero values.**
```go
var e [3]int
log.Printf("e type=%T, e length=%d, e capacity=%d, e=%v\n", e, len(e), cap(e), e)
```
- `var e [3]int` creates an array of size **3**, initialized with `[0, 0, 0]`.
- Unlike slices, **arrays cannot be `nil`**.

---

### **Rule 7: `new([N]T)` returns a pointer to an array.**
```go
f := new([3]int)
log.Printf("f type=%T, f length=%d, f capacity=%d, f=%v, nil? %t\n", f, len(f), cap(f), f, f == nil)
```
- `new([3]int)` creates a **pointer to an array** (`*[3]int`).
- The array itself is **not `nil`**.

---

### **Rule 8: Array literals must specify their length explicitly or use `[...]`.**
```go
j := [3]int{1, 3, 5}
log.Printf("j type=%T, j length=%d, j capacity=%d, j=%v\n", j, len(j), cap(j), j)
```
- `[3]int{1, 3, 5}` creates an array of **fixed size 3**.

```go
k := [...]int{1, 3, 5, 7}
log.Printf("k type=%T, k length=%d, k capacity=%d, k=%v\n", k, len(k), cap(k), k)
```
- `[...]int{1, 3, 5, 7}` lets Go determine the length automatically (**size 4**).

---

## **4. Special Case: Slice of Structs**
```go
i := []champion{
	{
		Name:    "Evelynn",
		Classes: []string{"Assassin"},
		Origins: []string{"Demon"},
		Cost:    3,
	},
	{
		Name:    "Vi",
		Classes: []string{"Brawler"},
		Origins: []string{"Hextech"},
		Cost:    3,
	},
}
log.Printf("i type=%T, i length=%d, i capacity=%d, i=%v, nil? %t\n", i, len(i), cap(i), i, i == nil)
```
- `i` is a **slice of structs**, initialized with two elements.
- **Valid slice operations apply**, like `append(i, champion{})`.

---

## **Key Takeaways**
1. **Slices vs. Arrays**:
   - Arrays `[N]T` have a fixed size, while slices `[]T` are dynamic.
   - Arrays can be compared; slices cannot (only `nil` checks).
  
2. **Slice Initialization**:
   - `var a []T` → **`nil` slice** (`len=0, cap=0`).
   - `new([]T)` → **Pointer to a nil slice**.
   - `make([]T, len, cap)` → **Allocates memory** (not nil).
   - `[]T{}` → **Empty but not nil**.
  
3. **Array Initialization**:
   - `var e [N]T` → **Zero-filled array**.
   - `new([N]T)` → **Pointer to an array**.
   - `[N]T{}` → **Explicit length**.
   - `[...]T{}` → **Let Go infer the length**.

4. **Capacity vs. Length**:
   - `len(slice)`: Number of elements.
   - `cap(slice)`: Underlying array size (can grow with `append`).

Would you like an example demonstrating how `append` works with capacity growth? 🚀