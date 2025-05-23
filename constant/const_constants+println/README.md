This Go program demonstrates the use of **constants** with different types, operations, and behaviors. Let's go through the rules and concepts step-by-step:

---

## 1. **Basic Constant Arithmetic**
```go
const (
	minsPerDay = 60 * 24
	weekDays   = 7
)
```
- Constants can perform arithmetic operations at compile time.
- `minsPerDay` calculates the number of minutes in a day (60 minutes × 24 hours).
- The expression `minsPerDay * weekDays * 2` calculates the total minutes in two weeks.

### ✅ **Rule**: Constants allow **compile-time expressions**, which makes them efficient and avoids runtime calculation overhead.

---

## 2. **Constant Expressions with Multiplication**
```go
const (
	hoursInDay   = 24
	daysInWeek   = 7
	hoursPerWeek = hoursInDay * daysInWeek
)
```
- `hoursPerWeek` is calculated as a constant expression (`24 * 7`).
- Then it multiplies by `5` to calculate the hours in 5 weeks.

---

## 3. **`len()` Function with Constants**
```go
const (
	home   = "Milky Way Galaxy"
	length = len(home)
)
```
- The `len()` function can be used with string literals in constant expressions.
- `len(home)` calculates the number of characters at **compile time**, not runtime.

---

## 4. **Floating-Point Constants**
```go
const (
	pi  = 3.14159265358979323846264
	tau = pi * 2
)
```
- Floating-point constants are supported.
- `tau` is calculated as `pi * 2`.
- `%g` in `Printf` is used to display a concise floating-point number.

---

## 5. **Typeless Constants**
```go
const (
	width  = 25
	height = width * 2
)
```
- When constants are declared without a specific type, they are considered **typeless constants**.
- The expression `width * height` works because the constants are implicitly treated as `int`.

---

## 6. **Using Constants with `time.Duration`**
```go
const later = 10
hours, _ := time.ParseDuration("1h")
fmt.Printf("%s later...\n", hours*later)
```
- The `time.ParseDuration("1h")` function returns a `time.Duration`.
- The constant `later` is typeless, so it can be multiplied with `hours`, which is of type `time.Duration`.

---

## 7. **Type Inference with Constants**
```go
const min int32 = 1

max := 5 + min
```
- Since `min` is explicitly typed as `int32`, the result of `5 + min` is also inferred as `int32`.
- Without explicit conversion, Go will not allow operations between `int32` and `int` due to type mismatch.

---

## ✅ **Key Takeaways**:
1. **Constants are evaluated at compile time**, not runtime.
2. **Typeless constants** are more flexible and can adapt to different types when used in expressions.
3. **Typed constants enforce type safety**, avoiding accidental type mismatches.
4. Go allows **built-in functions like `len()`** in constant expressions.
5. Go’s `time.Duration` can multiply with constants when the types are compatible.

---

Let me know if you'd like more examples or deeper insights! 😊