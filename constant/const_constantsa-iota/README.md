This Go program demonstrates the **rules and behavior of constants in Go**, including **typed and untyped constants**, **constant grouping**, and **the use of `iota` for automatic enumeration**.

---

## **1. Untyped vs Typed Constants**
```go
const a = 10       // untyped constant
const b string = "hello"  // typed constant
```
- **Untyped constants** (`a`) do not have a specific type until they are assigned to a variable.
- **Typed constants** (`b`) have a fixed type (`string`), and they enforce type safety.

### ✅ **Rule**: 
- An **untyped constant** can be assigned to different types of variables (e.g., `int`, `float64`, etc.), while a **typed constant** is restricted to its declared type.

---

## **2. Constants Can't Hold Mutable Data Structures**
```go
// const arr []int = []int{0, 1, 2} // error, array can not be constant
```
- Constants **cannot be arrays, slices, maps, or other mutable types**.
- Go constants are immutable and only support **basic types** like `int`, `string`, `bool`, and `float`.

---

## **3. Grouping Constants**
```go
const (
	c = 10
	d = "oh"
	e = true
)
```
- Grouping allows you to define multiple constants together.
- Each constant in the group is assigned a value explicitly or implicitly (if the previous constant's value is reused).

---

## **4. Using `iota` for Sequential Values**
```go
const (
	_ = iota       // 0 (ignored with `_`)
	f              // 1
	g              // 2
	h              // 3
	i = iota*2 + 1 // 4*2 + 1 = 9
)
```
- `iota` is a **counter that starts at 0** and increments by 1 for each constant declaration.
- The first `iota` is ignored with `_` to skip the 0 value.
- `f = 1`, `g = 2`, `h = 3`.
- `i` is explicitly set to `iota * 2 + 1`, which evaluates to `9`.

---

## **5. Constants Are Immutable**
```go
// a = "else" // error, constant is unchanged value
```
- Attempting to **change the value of a constant** results in a **compilation error**.

---

## **6. Output**
```go
fmt.Println(a)  // 10
fmt.Println(b)  // hello
fmt.Println(c)  // 10
fmt.Println(d)  // oh
fmt.Println(e)  // true
fmt.Println(f)  // 1
fmt.Println(g)  // 2
fmt.Println(h)  // 3
fmt.Println(i)  // 9
```
- The program prints the constant values as expected.

---

## ✅ **Key Takeaways**:
1. **Untyped constants** are more flexible, while **typed constants** enforce type safety.
2. **Constants cannot be mutable data structures** (like arrays or slices).
3. **Grouped constants** can reuse the previous value if the value is not explicitly specified.
4. **`iota` is a powerful tool for auto-incrementing constants**, often used in enumerations.
5. **Constants are immutable**, meaning their values cannot be changed after declaration.

---

Let me know if you'd like more examples or deeper details! 😊