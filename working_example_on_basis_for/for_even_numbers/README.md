This Go program demonstrates how to **create a slice of even numbers from 2 to 10** and print them.

---

## ✅ **Step-by-step explanation:**

### 🎯 1. **Slice Declaration:**
```go
var evenNumbers []int
```
- `var evenNumbers []int` declares a **slice of integers**.
- A **slice** in Go is a dynamic, resizable array that can grow or shrink as needed.

---

### 🎯 2. **For Loop to Generate Even Numbers:**
```go
for i := 2; i < 11; i += 2 {
	evenNumbers = append(evenNumbers, i)
}
```
- The loop **starts from 2**, and `i` is incremented by `2` after each iteration (`i += 2`), which ensures that only **even numbers** are generated.
- The condition `i < 11` ensures the loop stops at 10.
- `append()` is a built-in Go function used to **add elements to a slice**. It dynamically expands the slice to accommodate new elements.

---

### 🎯 3. **Print the Slice:**
```go
fmt.Println(evenNumbers)
```
- This line prints the `evenNumbers` slice to the console.

---

## ✅ **Final Output:**
```
[2 4 6 8 10]
```

---

## 🔥 **Key Concepts Explained:**
| Concept             | Explanation                                         |
|----------------|---------------------------------------------------|
| **Slice Declaration** | Declares an empty slice of integers. |
| **Loop with Step Size** | `i += 2` increments `i` by 2, generating even numbers. |
| **`append()` Function** | Adds new elements to the slice dynamically. |
| **Print Output** | Displays the final slice with even numbers. |

---

Let me know if you'd like more details or need help with other Go programs! 😊