### **Explanation of the Code**

This Go program demonstrates how to **check if a key exists in a map** and retrieve its value safely.

---

### **1. Initializing a Map**
```go
var people = map[string]int{
    "Tom":   1,
    "Bob":   2,
    "Sam":   4,
    "Alice": 8,
}
```
- A map `people` is created with **string keys** (`names`) and **int values** (`associated numbers`).
- Example entries:
  ```
  "Tom" → 1
  "Bob" → 2
  "Sam" → 4
  "Alice" → 8
  ```

---

### **2. Checking If a Key Exists in the Map**
```go
if val, ok := people["Tom"]; ok {
    fmt.Println(val) // Output: 1
}
```
- **`people["Tom"]`** tries to retrieve the value associated with `"Tom"`.
- The **comma-ok idiom (`val, ok := people["Tom"]`)** is used:
  - `val` stores the value from the map (`1` if `"Tom"` exists).
  - `ok` is a **boolean** indicating whether `"Tom"` exists in `people`.
- **If `"Tom"` exists (`ok == true`), the value is printed.**
- Output:
  ```
  1
  ```

---

### **3. What Happens If the Key Doesn't Exist?**
- If you check for a non-existent key (e.g., `"John"`):
  ```go
  if val, ok := people["John"]; ok {
      fmt.Println(val)
  }
  ```
  - Since `"John"` **is not in the map**, `ok` will be `false`, and the `if` block **won't execute**.
  - No output will be printed.

---

## **Key Takeaways**
✅ **Checking Key Existence:** Use `val, ok := map[key]` to check if a key exists.  
✅ **Avoiding Zero-Values:** If a key **does not exist**, `val` gets the zero-value (`0` for integers).  
✅ **Safe Key Lookup:** Using `if ok {}` ensures you only access the value if the key exists.  

Would you like an example where we modify the map dynamically? 🚀