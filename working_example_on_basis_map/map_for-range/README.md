### **Explanation of the Code**
This Go program demonstrates how to use a **map** with an `int` key and a `string` value. The key concepts here include **unordered maps** and **iteration over maps**.

---

### **1. Creating a Map**
```go
map1 := map[int]string{
    0: "first",
    1: "second",
    2: "third",
    3: "fourth",
}
```
- A **map** named `map1` is defined.
- The **keys** are integers: `0, 1, 2, 3`.
- The **values** are strings: `"first"`, `"second"`, `"third"`, `"fourth"`.

---

### **2. Iterating Over the Map**
```go
for key, val := range map1 {
    fmt.Println("key", key, "value", val)
}
```
- Uses a **for-range loop** to iterate through the `map1`.
- **Each iteration provides:**
  - `key`: The integer key.
  - `val`: The string value.

#### **Important Note: Maps in Go Are Unordered**
- The elements of a **map in Go are not stored in a specific order**.
- When iterating, the output order may **not** match the insertion order.

Example Output:
```
key 2 value third
key 0 value first
key 1 value second
key 3 value fourth
```
**(The order may change in different runs of the program.)**

---

### **Why Are Maps Unordered in Go?**
- Go **optimizes maps for fast lookup** rather than maintaining insertion order.
- The way Go internally stores maps uses **hash tables**, which do not guarantee any specific order when iterating.

---

### **Key Takeaways**
✅ **Go maps are unordered**—iteration order may vary.  
✅ **Using a for-range loop**, you can iterate over key-value pairs.  
✅ **Fast lookups**—maps are optimized for performance rather than order.  

Would you like an example of how to sort map keys before printing? 🚀