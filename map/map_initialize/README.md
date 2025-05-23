## **Explanation of the Code**

This Go program demonstrates:
- **Different ways to initialize maps**
- **Checking for existence of a key in a map**
- **Deleting keys from a map**
- **Difference between reference types (maps) and value types (arrays)**

---

### **1. Initializing a Map (Method 1)**
```go
var map1 = make(map[string]int)
```
- **`make(map[string]int)`** initializes an empty map with string keys and int values.

```go
map1["key1"] = 22
map1["key2"] = 24
fmt.Println("map1", map1)
```
- Adds `"key1": 22` and `"key2": 24` to the map.
- Output:
  ```
  map1 map[key1:22 key2:24]
  ```

---

### **2. Deleting a Key from a Map**
```go
delete(map1, "key2")
fmt.Println("map1", map1)
```
- Removes `"key2"` from `map1`.
- Output:
  ```
  map1 map[key1:22]
  ```

---

### **3. Checking If a Key Exists**
```go
v, ok := map1["key1"]
fmt.Println("ok?:", ok, "value:", v)
```
- **If `key1` exists**, `ok` is `true`, and `v` is the value (`22`).
- Output:
  ```
  ok?: true value: 22
  ```

```go
v, ok = map1["key2"]
fmt.Println("ok?:", ok, "value:", v)
```
- **If `key2` does not exist**, `ok` is `false`, and `v` gets the default int value (`0`).
- Output:
  ```
  ok?: false value: 0
  ```

---

### **4. Initializing a Map (Method 2)**
```go
map2 := map[string]int{
    "key1": 22,
    "key2": 24,
}
fmt.Println("map2", map2)
```
- **Alternate way to declare and initialize a map** using a map literal.
- Output:
  ```
  map2 map[key1:22 key2:24]
  ```

---

### **5. Checking if a Map is Nil**
```go
var map3 map[int]int
fmt.Println("map3 is nil?", map3 == nil)
```
- `map3` is declared **without initialization**, so it's `nil`.
- **A nil map cannot store values** (causes an error if accessed).
- Output:
  ```
  map3 is nil? true
  ```
- **Fix:** Initialize it using `make(map[int]int)` before assigning values.

---

### **6. Arrays vs Maps**
```go
var arr [3]int
fmt.Println("arr", arr)
arr[1] = 2
fmt.Println("arr", arr)
```
- **Arrays in Go are value types**, unlike maps (which are reference types).
- If an array is not initialized, **it still has default values (zero-values).**
- Default int values in an array are `0`.
- Output:
  ```
  arr [0 0 0]
  arr [0 2 0]
  ```

---

## **Key Takeaways**
✅ **Maps are reference types** – they are `nil` if uninitialized.  
✅ **Use `make(map[keyType]valueType)`** to initialize a map before using it.  
✅ **Check if a key exists** using `value, ok := map[key]`.  
✅ **Deleting a key** from a map removes it completely.  
✅ **Arrays are value types** – they always have a default value, unlike maps.  

Would you like an example of iterating over a map or using structs inside maps? 🚀