This Go program demonstrates **maps** in detail, including **creating, modifying, deleting elements, checking existence, and getting values**. Let's go step by step.

---

## **1. Creating a Map with Values**
```go
var money map[string]int = map[string]int{
	"dollar": 1000,
	"euro":   100,
	"ap":     5,
}
```
- A **map** is created where:
  - The **keys** are `string` (`"dollar"`, `"euro"`, `"ap"`)
  - The **values** are `int` (`1000`, `100`, `5`).
- Prints the whole map:
  ```
  map[ap:5 dollar:1000 euro:100]
  ```
- Prints the value associated with `"dollar"`:
  ```
  1000
  ```

---

## **2. Creating and Modifying a Map**
```go
money2 := map[string]int{
	"dollar": 1000,
	"euro":   100,
	"ap":     5,
}
```
- Creates another map `money2` in the **same way**.
- Prints:
  ```
  map[ap:5 dollar:1000 euro:100]
  ```

### **Updating a Value**
```go
money2["dollar"] = 5000
```
- Updates the value of `"dollar"` to `5000`.
- Prints:
  ```
  map[ap:5 dollar:5000 euro:100]
  ```

### **Deleting a Key**
```go
delete(money2, "ap")
```
- Removes the key `"ap"`.
- Prints:
  ```
  map[dollar:5000 euro:100]
  ```

---

## **3. Checking for Non-Existent Keys**
```go
fmt.Println(money2["gg"])
```
- `"gg"` **does not exist**, so Go returns **default int value** → `0`.
- Prints:
  ```
  0
  ```

### **Checking Key Existence Properly**
```go
val, status := money2["dollar"]
fmt.Println(val, status)
```
- `"dollar"` exists, so:
  - `val = 5000`
  - `status = true`
- Prints:
  ```
  5000 true
  ```

```go
val2, status2 := money2["gg"]
fmt.Println(val2, status2)
```
- `"gg"` does not exist:
  - `val2 = 0`
  - `status2 = false`
- Prints:
  ```
  0 false
  ```

---

## **4. Creating a Map with `make()`**
```go
m1 := make(map[string]int)
```
- Creates an **empty map** where:
  - `string` keys
  - `int` values
- Adds values:
```go
m1["foo"] = 100
m1["baa"] = 200
```

---

## **5. Declaring a Map with Initial Values**
```go
m := map[string]int{"foo": 300, "baa": 400}
```
- Creates a **map with predefined values**.
- Prints:
  ```
  map[baa:400 foo:300]
  ```
- **Checking number of items**:
  ```go
  fmt.Println(len(m))
  ```
  - Prints `2` since it has two key-value pairs.

---

## **6. Deleting a Key**
```go
delete(m, "foo")
fmt.Println(m)
```
- Deletes `"foo"`, leaving only `"baa": 400`.
- Prints:
  ```
  map[baa:400]
  ```

---

## **7. Checking if a Key Exists**
```go
v, ok := m["baa"]
fmt.Println(v)  // prints value
fmt.Println(ok) // prints true/false
```
- `"baa"` exists:
  - `v = 400`
  - `ok = true`
- Prints:
  ```
  400
  true
  ```

---

## **Summary of Key Concepts**
| **Action** | **Code** |
|------------|---------|
| Create a map | `m := make(map[string]int)` |
| Create a map with values | `m := map[string]int{"key": value}` |
| Add a key-value pair | `m["newKey"] = value` |
| Update a value | `m["existingKey"] = newValue` |
| Get a value | `val := m["key"]` |
| Check if key exists | `val, exists := m["key"]` |
| Delete a key | `delete(m, "key")` |
| Get map length | `len(m)` |

This program is a **comprehensive guide to Go maps**, covering **creation, modification, deletion, and key existence checks**! 🚀