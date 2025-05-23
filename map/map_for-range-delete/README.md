### **Explanation of the Code**
This Go program demonstrates the use of **maps** (key-value pairs), including how to **create**, **add**, **delete**, and **iterate over** a map.

---

### **1. Creating a Map**
```go
languages := make(map[string]string)
```
- The `make(map[string]string)` function creates an **empty map**.
- The **key** type is `string`, and the **value** type is also `string`.

---

### **2. Adding Key-Value Pairs**
```go
languages["JS"] = "Javascript"
languages["RB"] = "Ruby"
languages["PY"] = "Python"
```
- Three key-value pairs are added:
  - `"JS"` → `"Javascript"`
  - `"RB"` → `"Ruby"`
  - `"PY"` → `"Python"`

After this, the map looks like:
```go
map[JS:Javascript RB:Ruby PY:Python]
```

---

### **3. Printing the Map**
```go
fmt.Println("List of all languages: ", languages)
```
**Output:**
```
List of all languages:  map[JS:Javascript RB:Ruby PY:Python]
```

---

### **4. Accessing a Specific Value**
```go
fmt.Println("JS is shorts for: ", languages["JS"])
```
**Output:**
```
JS is shorts for:  Javascript
```

---

### **5. Deleting a Key-Value Pair**
```go
delete(languages, "RB")
```
- **Deletes the key `"RB"` (Ruby) from the map.**
- The updated map:
```go
map[JS:Javascript PY:Python]
```

---

### **6. Iterating Over the Map**
```go
for key, value := range languages {
	fmt.Printf("For key %v, value is %v\n", key, value)
}
```
- The **for-range loop** iterates over all key-value pairs.
- Example output (order may vary because **maps are unordered** in Go):
```
For key JS, value is Javascript
For key PY, value is Python
```

---

### **Key Takeaways**
✅ **Maps store key-value pairs** and allow fast lookups.  
✅ **You can add, retrieve, and delete values using keys.**  
✅ **Maps in Go are unordered**, so iteration order is unpredictable.  
✅ **`delete(map, key)` removes a key-value pair from the map.**  
✅ **You can use a `for-range` loop** to iterate over key-value pairs.  

Would you like an example where we check if a key exists before deleting it? 🚀