This Go program demonstrates the use of **maps** (similar to dictionaries in Python). Let's go through the code step by step.

---

### **1. Creating a Map**
```go
m := make(map[string]int)
```
- `make(map[string]int)` creates an **empty map** where:
  - **Keys** are of type `string`
  - **Values** are of type `int`
- The variable `m` stores this map.

---

### **2. Adding Key-Value Pairs**
```go
m["k1"] = 7
m["k2"] = 13
```
- Adds two key-value pairs to `m`:
  - `"k1"` → `7`
  - `"k2"` → `13`

---

### **3. Printing the Entire Map**
```go
fmt.Println("map:", m)
```
- Prints the map, which now contains:
  ```
  map: map[k1:7 k2:13]
  ```
  (*The order may vary because Go maps do not guarantee order!*)

---

### **4. Accessing a Value by Key**
```go
fmt.Println("map [k1]:", m["k1"])
```
- Retrieves the value associated with `"k1"`, which is `7`.

#### Output:
```
map [k1]: 7
```

---

### **5. Checking the Length of the Map**
```go
fmt.Println("len:", len(m))
```
- Prints the **number of key-value pairs** in the map.
  
#### Output:
```
len: 2
```
(*There are two entries: `"k1":7` and `"k2":13`*)

---

### **6. Deleting an Entry from the Map**
```go
delete(m, "k2")
fmt.Println("len:", len(m))
```
- **Deletes** the key `"k2"`.
- Prints the new length of the map.

#### Output:
```
len: 1
```
(*Only `"k1"` remains in the map!*)

---

### **7. Checking If a Key Exists**
```go
key, value := m["k3"]
fmt.Println("key:", key, "value", value)
```
- **Attempts to retrieve the value** associated with `"k3"`, which doesn't exist.
- In Go, when accessing a non-existent key:
  - The **default value** for the value type is returned (`0` for `int`).
  - The second return value (`value`) is **false** (indicating the key is missing).

#### Output:
```
key: 0 value false
```

---

### **8. Declaring and Initializing a Map Inline**
```go
n := map[string]int{"foo": 1, "bar": 2}
fmt.Println("map:", n)
```
- **Creates a map with values already initialized**:
  ```
  "foo" → 1
  "bar" → 2
  ```
- Prints the map.

#### Output:
```
map: map[bar:2 foo:1]
```
(*The order may vary!*)

---

## **Summary of Key Concepts**
✅ **Creating a Map:**  
```go
m := make(map[string]int)
```
✅ **Adding Key-Value Pairs:**  
```go
m["key"] = value
```
✅ **Retrieving a Value:**  
```go
val := m["key"]
```
✅ **Checking if a Key Exists:**  
```go
val, exists := m["key"]
```
✅ **Deleting a Key:**  
```go
delete(m, "key")
```
✅ **Initializing a Map with Values:**  
```go
n := map[string]int{"key1": value1, "key2": value2}
```

This program effectively showcases **how maps work in Go**, including **adding, retrieving, deleting, and checking for keys**. 🚀