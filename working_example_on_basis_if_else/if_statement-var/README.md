### **Code Explanation**
This Go program demonstrates **variable declaration inside an `if` statement**.

---

### **Code Breakdown**
```go
package main

import (
	"fmt"
)

func main() {

	if b := 22; b > 20 {
		fmt.Println(b)
	}

	// We can put an assignment statement in an if condition.
	// The scope of the variable `b` is only within the if block.

}
```

---

### **Output**
```
22
```

---

### **Explanation of Rules**
#### **1. Declaring a Variable in `if` Statement**
```go
if b := 22; b > 20 {
    fmt.Println(b)
}
```
- The **assignment statement `b := 22`** creates a new variable `b` inside the `if` condition.
- **Scope Rule:** `b` only exists inside the `if` block.
- The condition `b > 20` is `true` (since `22 > 20`).
- Because the condition is **true**, the `fmt.Println(b)` statement executes, printing `22`.

---

#### **2. Variable Scope in `if` Statement**
- The variable `b` **only exists inside the `if` block**.
- If you try to use `b` outside the `if`, **it will cause an error**.

❌ **Invalid Example (outside scope):**
```go
if b := 22; b > 20 {
    fmt.Println(b)
}
fmt.Println(b)  // ERROR: undefined variable b
```
✔ **Valid Example (using `b` inside `if`):**
```go
if b := 22; b > 20 {
    fmt.Println(b)  // Works fine
}
```

---

### **Key Takeaways**
1. ✅ You can **declare a variable inside an `if` statement** using `:=` before the semicolon (`;`).
2. ✅ The declared variable **only exists inside the `if` block**.
3. ✅ This helps in **reducing unnecessary variable declarations** in the main function.

---

### **Example with `else` Block**
The **scope of `b` is still limited** to the `if-else` block:
```go
if b := 18; b > 20 {
    fmt.Println("b is greater than 20")
} else {
    fmt.Println("b is 20 or less:", b)  // Works fine
}

fmt.Println(b)  // ❌ ERROR: b is undefined here
```

---

Let me know if you need more details! 🚀😊