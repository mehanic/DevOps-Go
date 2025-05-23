This Go program illustrates various rules and behaviors regarding constants and variables, as well as some common mistakes when using constants. Here's a breakdown of the key concepts explained in the comments:

---

### 1. **Constants in Go (Compile-Time Calculation)**
```go
// const ---> compile time
// var ---> run time
```
- **Constants** are **evaluated at compile time**. They must be known during compilation, so they cannot be set from runtime values or expressions that require runtime evaluation.
- **Variables** are assigned values during **runtime**.

---

### 2. **Re-Assignment of Constants**
```go
const x = 2
// x = 5 // Error: cannot assign to constant
```
- **Constants** cannot be reassigned or changed once they are defined. If you try to assign a new value to `x`, it will cause an error, as constants are immutable.

---

### 3. **Constants and Expressions**
```go
// const x2 = math.Pow(3, 4) // Error: math.Pow is a runtime function
```
- Constants must be initialized with values known at **compile time**.
- Functions like `math.Pow()` are **runtime functions** because they require calculations during program execution, so they **cannot be used** to initialize constants.

---

### 4. **Combining Constants and Variables**
```go
// y := 3
// const x = y // Error: y is a variable, and variables are runtime values
```
- **You cannot assign a variable to a constant** directly. Constants must be initialized with **constant expressions** (those that are evaluated at compile time). In this case, `y` is a variable, so you cannot use it to initialize a constant.

---

### 5. **Declaring Multiple Constants**
```go
const x, y = 3, 5
fmt.Printf("%T, %v\n", x, x)  // int, 3
fmt.Printf("%T, %v\n", y, y)  // int, 5
```
- You can declare multiple constants in a single line. 
- The **types** of constants are automatically inferred from the values.
  - `x` and `y` are both `int` constants because `3` and `5` are integers.
  - `const x, y = 3, 5` implicitly declares `x` and `y` as constants with type `int` because integer literals are inferred as `int`.

---

### 6. **`const` vs. `var` Usage**
```go
// var x1 = math.Pow(3, 4) // Valid: `var` can be used with runtime calculations
// const x2 = math.Pow(3, 4) // Error: `const` cannot use runtime calculations
```
- **Variables (`var`)** can use **runtime values** or functions (like `math.Pow()`) to initialize them because variables are assigned values at runtime.
- **Constants** must have values that are **known at compile time**.

---

### 7. **Constant Initialization**
```go
// const x int
// x = 3 // Error: Cannot assign value to uninitialized constant
```
- When you declare a constant, you **must assign a value to it immediately**. You cannot declare it first and assign the value later, because constants are evaluated at **compile time**.

---

### 8. **Printing Type and Value of Constants**
```go
fmt.Printf("%T, %v\n", x, x)  // Output: int, 3
fmt.Printf("%T, %v\n", y, y)  // Output: int, 5
```
- You can print both the **type** (`%T`) and the **value** (`%v`) of constants using `fmt.Printf`.

---

### **Summary of Rules:**
1. **Constants are evaluated at compile time** and must be initialized with values that can be known before the program runs (no runtime functions or variables).
2. **Constants cannot be changed** once declared.
3. **Variables (`var`)** can be initialized with values known at runtime, including the result of functions like `math.Pow`.
4. **Multiple constants** can be declared in one statement, and their types are inferred from the values.
5. Constants cannot be **assigned from variables** directly, as variables are evaluated at runtime.

Let me know if you need more clarification or examples! 😊