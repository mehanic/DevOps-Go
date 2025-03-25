This Go program demonstrates **basic conditional (if) statements** using **comparison operators**.

---

## **Code Breakdown**
```go
package main

import (
	"fmt"
)

func main() {
	x := 40

	// Check if x is equal to 40
	if x == 40 {
		fmt.Println("001")
	}

	// Check if x is NOT equal to 40
	if x != 40 {
		fmt.Println("002")
	}

	// Check if x is less than or equal to 50
	if x <= 50 {
		fmt.Println("003")
	}

	// Check if x is greater than or equal to 0
	if x >= 0 {
		fmt.Println("004")
	}
}
```

---

## **Output**
```
001
003
004
```

---

## **Explanation of Rules**
### **1. `if x == 40`**
```go
if x == 40 {
    fmt.Println("001")
}
```
- **`==` (Equal to)**: This checks if `x` is exactly `40`.  
- Since `x` is `40`, this condition is `true`, so `"001"` is printed.

---

### **2. `if x != 40`**
```go
if x != 40 {
    fmt.Println("002")
}
```
- **`!=` (Not Equal to)**: This checks if `x` is **not** `40`.  
- Since `x` is `40`, this condition is `false`, so `"002"` is **not** printed.

---

### **3. `if x <= 50`**
```go
if x <= 50 {
    fmt.Println("003")
}
```
- **`<=` (Less than or equal to)**: This checks if `x` is **less than or equal to** `50`.
- Since `40` is **less than** `50`, this condition is `true`, so `"003"` is printed.

---

### **4. `if x >= 0`**
```go
if x >= 0 {
    fmt.Println("004")
}
```
- **`>=` (Greater than or equal to)**: This checks if `x` is **greater than or equal to** `0`.
- Since `40` is **greater than** `0`, this condition is `true`, so `"004"` is printed.

---

## **Summary of Comparison Operators**
| Operator | Meaning             | Example (`x = 40`) |
|----------|---------------------|--------------------|
| `==`     | Equal to            | `x == 40` → ✅ True  |
| `!=`     | Not equal to        | `x != 40` → ❌ False |
| `>`      | Greater than        | `x > 50` → ❌ False |
| `<`      | Less than           | `x < 50` → ✅ True  |
| `>=`     | Greater than or equal to | `x >= 0` → ✅ True  |
| `<=`     | Less than or equal to | `x <= 50` → ✅ True  |

---

Let me know if you have any questions! 🚀😃