This Go program demonstrates basic operations in **string concatenation**, **arithmetic operations**, and **logical operations**.

---

## ✅ **1. String Concatenation:**
```go
fmt.Println("go" + "lang")
```
- The `+` operator concatenates (joins) two strings.
- `"go"` and `"lang"` are combined to produce `"golang"`.

---

## ✅ **2. Arithmetic Operations:**
```go
fmt.Println("1+1 =", 1+1)
```
- Basic integer addition.
- `1 + 1` results in `2`.

---

```go
fmt.Println("7.0/3.0 =", 7.0/3.0)
```
- Floating-point division.
- `7.0 / 3.0` results in `2.3333333333333335`, because Go performs **floating-point division** when both numbers are decimals.

---

## ✅ **3. Logical Operations:**
### 🎯 AND Operator (`&&`):
```go
fmt.Println(true && false)
```
- Returns `true` **only if both conditions are true**.
- `true && false` is `false`.

---

### 🎯 OR Operator (`||`):
```go
fmt.Println(true || false)
```
- Returns `true` **if at least one condition is true**.
- `true || false` is `true`.

---

### 🎯 NOT Operator (`!`):
```go
fmt.Println(!true)
```
- Reverses the boolean value.
- `!true` is `false`.

---

## ✅ **Output:**
```
golang
1+1 = 2
7.0/3.0 = 2.3333333333333335
false
true
false
```

---

## 🎯 **Key Concepts:**
| Concept             | Explanation                            |
|-----------------|----------------------------------|
| **String Concatenation** | Combines strings with `+`. |
| **Arithmetic Operations** | Supports basic math operations (`+`, `-`, `*`, `/`, `%`). |
| **Boolean Logic** | Uses `&&` (AND), `||` (OR), and `!` (NOT). |

