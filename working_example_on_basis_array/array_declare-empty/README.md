This Go program demonstrates **array declarations, types, zero-values, and formatting rules**. Let's break down the key rules governing this behavior.

---

## **1. Array Declarations & Zero-Value Initialization**
In Go, arrays have:
- A **fixed size** (`[N]T`).
- A **specific type** (`T` is the element type).
- An **automatic zero-value initialization** for all elements.

### **Declared Arrays and Their Zero Values**
```go
var (
	names     [3]string  // Array of 3 strings (empty strings by default)
	distances [5]int     // Array of 5 integers (initialized to 0)
	data      [5]byte    // Array of 5 bytes (0x0 default values)
	ratios    [1]float64 // Array of 1 float64 (0.00 default value)
	alives    [4]bool    // Array of 4 booleans (false default values)
	zero      [0]byte    // Empty array (no elements, takes no memory)
)
```
### **Key Rules**
1. **Arrays are always initialized** to their **zero values**:
   - `string` → `""` (empty string)
   - `int` → `0`
   - `byte` → `0x0`
   - `float64` → `0.0`
   - `bool` → `false`
   - `[0]byte` (zero-length array) → `[]` (empty)

---

## **2. Printing Arrays with `%#v`**
```go
fmt.Printf("names    : %#v\n", names)
fmt.Printf("distances: %#v\n", distances)
fmt.Printf("data     : %#v\n", data)
fmt.Printf("ratios   : %#v\n", ratios)
fmt.Printf("alives   : %#v\n", alives)
fmt.Printf("zero     : %#v\n", zero)
```
- `%#v` prints **Go syntax representation**.
- Example output:
  ```go
  names    : [3]string{"", "", ""}
  distances: [5]int{0, 0, 0, 0, 0}
  data     : [5]uint8{0x0, 0x0, 0x0, 0x0, 0x0}
  ratios   : [1]float64{0}
  alives   : [4]bool{false, false, false, false}
  zero     : [0]uint8{}
  ```
- **Observations**:
  - Strings print as `""` because they are **empty**.
  - Integers, bytes, and floats print as `0`, `0x0`, `0.0` respectively.
  - Booleans print as `false`.
  - `[0]byte{}` → Zero-length array appears as `{}` (empty).

---

## **3. Printing Array Types with `%T`**
```go
fmt.Printf("names    : %T\n", names)
fmt.Printf("distances: %T\n", distances)
fmt.Printf("data     : %T\n", data)
fmt.Printf("ratios   : %T\n", ratios)
fmt.Printf("alives   : %T\n", alives)
fmt.Printf("zero     : %T\n", zero)
```
- `%T` prints the **exact type** of the variable.
- Example output:
  ```
  names    : [3]string
  distances: [5]int
  data     : [5]uint8
  ratios   : [1]float64
  alives   : [4]bool
  zero     : [0]uint8
  ```
- **Observations**:
  - Each array **includes its length in the type** (`[N]T`).
  - `[3]string` and `[5]int` are **not interchangeable**.
  - `[0]byte` is a valid type but has no elements.

---

## **4. Printing Array Elements with Formatting**
```go
fmt.Printf("names    : %q\n", names)
fmt.Printf("distances: %d\n", distances)
fmt.Printf("data     : %d\n", data)
fmt.Printf("ratios   : %.2f\n", ratios)
fmt.Printf("alives   : %t\n", alives)
fmt.Printf("zero     : %d\n", zero)
```
- `%q` → **Quoted strings** (for string arrays).
- `%d` → **Decimal integers** (for `int` and `byte` arrays).
- `%.2f` → **Fixed decimal places for floats**.
- `%t` → **Boolean format (`true/false`)**.

### **Example Output**
```
names    : ["" "" ""]
distances: [0 0 0 0 0]
data     : [0 0 0 0 0]
ratios   : [0.00]
alives   : [false false false false]
zero     : []
```
### **Observations**
1. **Strings**: `"[]"` for empty values.
2. **Numbers**: `[0 0 0 0 0]`, matching default zero values.
3. **Booleans**: `[false false false false]`.
4. **Zero-length array (`[0]byte`)**: Prints as `[]` (empty).

---

## **5. Special Case: Zero-Length Arrays**
```go
var zero [0]byte
fmt.Printf("zero     : %#v\n", zero) // Prints: zero : [0]uint8{}
fmt.Printf("zero     : %T\n", zero)  // Prints: zero : [0]uint8
```
### **Zero-Length Arrays (`[0]T`)**
- **Valid type in Go** (`[0]T`).
- **Takes no memory space**.
- **Still a valid array** (can be assigned but has no elements).

---

## **Key Takeaways**
### **Arrays:**
1. **Fixed-size & type-specific** (`[N]T`).
2. **Cannot change size** after declaration.
3. **Zero-initialized** (empty strings, `0` for numbers, `false` for bools).
4. **Zero-length arrays** (`[0]T`) take **no memory** but are **valid types**.

### **Printing Rules:**
| **Format** | **Usage** | **Example Output** |
|------------|----------|------------------|
| `%#v` | Go-syntax format | `[3]string{"", "", ""}` |
| `%T` | Type of variable | `[5]int` |
| `%q` | Quoted string format | `["" "" ""]` |
| `%d` | Decimal integers | `[0 0 0 0 0]` |
| `%t` | Boolean format | `[false false false false]` |
| `%.2f` | Floating-point format | `[0.00]` |

Would you like a deep dive into **how arrays differ from slices** in Go? 🚀