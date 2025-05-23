This Go program demonstrates how to use **maps (dictionaries in other languages like Python)**. Let’s break it down step by step:

---

### **1. Initial Print Statement**
```go
fmt.Println("Go tutorial for Maps")
```
- This prints a simple message to indicate the program is about maps.

---

### **2. Creating a Map**
```go
mp := make(map[string]string)
```
- **`make(map[string]string)`** creates an empty map where:
  - The **keys** are of type `string`.
  - The **values** are also of type `string`.
- The variable `mp` stores the map.

---

### **3. Adding Key-Value Pairs**
```go
mp["rohan"] = "patel"
mp["hello"] = "world"
mp["hotel"] = "trivigo"
mp["hi"] = "bye"
```
- Here, we add key-value pairs to the map:
  - `"rohan" → "patel"`
  - `"hello" → "world"`
  - `"hotel" → "trivigo"`
  - `"hi" → "bye"`

---

### **4. Printing the Entire Map**
```go
fmt.Println("Map is :", mp)
```
- This prints all key-value pairs stored in `mp`.

#### Example Output:
```
Map is : map[hello:world hi:bye hotel:trivigo rohan:patel]
```
(*The order of the elements may vary since Go maps do not guarantee order!*)

---

### **5. Accessing a Value by Key**
```go
fmt.Println("last name of rohan is", mp["rohan"])
```
- This retrieves and prints the value associated with the key `"rohan"`.
  
#### Output:
```
last name of rohan is patel
```

---

### **6. Deleting an Entry from the Map**
```go
delete(mp, "hi")
```
- This **removes the key `"hi"` and its associated value `"bye"`** from the map.

#### Map Before Deletion:
```
map[hello:world hi:bye hotel:trivigo rohan:patel]
```
#### Map After Deletion:
```
map[hello:world hotel:trivigo rohan:patel]
```

---

### **7. Iterating Over the Map**
```go
for i, v := range mp {
	fmt.Println(i, ":", v)
}
```
- The `for range` loop iterates over the map:
  - `i` stores the **key**.
  - `v` stores the **value**.
- Each key-value pair is printed.

#### Example Output:
```
hello : world
hotel : trivigo
rohan : patel
```
(*The order may be different each time the program runs!*)

---

### **Key Takeaways**
1. **Maps store key-value pairs.**
2. **Creating a map:** Use `make(map[keyType]valueType)`.
3. **Adding elements:** `map[key] = value`.
4. **Accessing a value:** `map[key]` returns the value.
5. **Deleting elements:** `delete(map, key)`.
6. **Iterating over a map:** Use `for key, value := range map`.

This is a simple introduction to **maps in Go**, which are very efficient for lookups and data storage. 🚀