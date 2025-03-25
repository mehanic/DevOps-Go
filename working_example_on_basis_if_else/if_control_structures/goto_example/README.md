### **📌 Explanation of the `goto` Statement in Go**

In this Go program, you use a `goto` statement to jump to a specific label inside the function.

---

### **🔹 Code Breakdown**
```go
package main

import "fmt"

func main() {
	goto customLabel  // ⬅️ Jumps directly to 'customLabel'

	// This line is skipped because of 'goto'
	fmt.Println("Hello")

customLabel:  // ⬅️ This is the label where execution jumps
	fmt.Println("World")
}
```

---

### **🔹 Execution Flow**
1️⃣ **`goto customLabel`** makes the program jump to the `customLabel` label.  
2️⃣ The `fmt.Println("Hello")` line is **never executed** because `goto` skips over it.  
3️⃣ The execution continues from `customLabel`, printing `"World"` to the console.

---

### **💡 Expected Output**
```
World
```
(The `"Hello"` line is completely skipped.)

---

### **🔹 Key Takeaways**
✅ **`goto` jumps directly to a labeled statement** inside the same function.  
✅ It **skips** any code between the `goto` statement and the target label.  
✅ **Labels must be in the same function** (you can't jump across functions).  
✅ `goto` is rarely used in modern Go programming—it can make code harder to read.

Would you like an example of a practical use case for `goto` in Go? 😊