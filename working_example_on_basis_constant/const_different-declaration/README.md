This Go program demonstrates **constant rules, immutability, and the behavior of compile-time and runtime expressions**. Let's break it down:

---

## ✅ **1. Magic Values vs. Named Constants**
```go
cm := 100
m := cm / 100 // 100 is a magic value
```
- **Magic values** are hardcoded numbers with no clear meaning, which makes code less readable and harder to maintain.
- The improved version uses a **named constant** for clarity:
```go
const meters int = 100
m := cm / meters
```
- This improves code readability and prevents accidental mistakes.

---

## ✅ **2. Compile-Time vs. Runtime Errors**
### **Compile-Time Error (Detected by Go Compiler)**
```go
// const n int = 1
// const m int = 0
// fmt.Println(n / m) // Division by zero, detected at compile-time
```
- Go **catches division by zero at compile time** when constants are involved.

### **Runtime Error (Not Detected by Compiler)**
```go
n, m := 1, 0
fmt.Println(n / m) // Runtime panic: integer divide by zero
```
- When using variables, the error is detected at **runtime**, leading to a **panic**.

---

## ✅ **3. Immutability of Constants**
```go
const max int = 100
// max = 200 // ERROR: cannot assign to max
```
- **Constants are immutable**, meaning their value **cannot be changed** after declaration.

---

## ✅ **4. Constants Cannot Use Runtime Constructs**
```go
// const max int = math.Pow10(2) // ERROR: math.Pow is not a constant
```
- `math.Pow10` is a **function call**, which is a **runtime construct**, not allowed in constant expressions.

---

## ✅ **5. Constants Cannot Use Non-Constant Variables**
```go
n := 2
// const max int = n // ERROR: n is not a constant
```
- A constant **cannot be assigned a value from a variable**, as variables are evaluated at runtime.

---

## ✅ **6. `len()` Works in Constant Expressions**
```go
const max int = len("Hello")
```
- The `len()` function can be used in a constant expression **only when the argument is a constant**, like a string literal.

---

## 🎯 **Key Takeaways:**
1. **Magic values** should be avoided; use **named constants** instead for clarity.
2. **Compile-time errors** help catch division by zero when using constants.
3. **Constants are immutable** and cannot be reassigned.
4. Constants **cannot use runtime functions or variables**, only compile-time expressions.
5. The `len()` function is an exception when used with constant arguments.

---

Let me know if you'd like more examples or need further clarifications! 😊