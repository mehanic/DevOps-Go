### **📌 Explanation of `for range` Loop in Go**
This Go program demonstrates different ways to iterate over **slices** and **maps** using the `for range` loop.

---

### **🔹 Code Breakdown**
```go
package main

import "fmt"

func main() {
	// Define a slice of integers
	intSlice := []int{2, 4, 6, 8}

	// Iterate over the slice
	for key, value := range intSlice {
		fmt.Println(key, value)
	}
```
✅ `intSlice := []int{2, 4, 6, 8}` → A slice (dynamic array) is created.  
✅ `for key, value := range intSlice` → Iterates over the slice.  
✅ `key` holds the **index**, and `value` holds the **element** at that index.  
✅ **Expected Output for this loop:**
```
0 2
1 4
2 6
3 8
```

---

### **🔹 Iterating Over a Map**
```go
	myMap := map[string]string{
		"d": "Donut",
		"o": "Operator",
	}

	// Iterate over a map
	for key, value := range myMap {
		fmt.Println(key, value)
	}
```
✅ `myMap := map[string]string{"d": "Donut", "o": "Operator"}` → A map (dictionary) is created.  
✅ `for key, value := range myMap` → Iterates over the map, printing both keys and values.  
✅ **Expected Output (unordered, because maps are unordered):**
```
d Donut
o Operator
```
or
```
o Operator
d Donut
```
(Map iteration order is not guaranteed.)

---

### **🔹 Iterating Over a Map (Keys Only)**
```go
	// Iterate but only utilize keys
	for key := range myMap {
		fmt.Println(key)
	}
```
✅ `for key := range myMap` → Iterates over the **keys only**.  
✅ **Expected Output (unordered):**
```
d
o
```
or  
```
o
d
```

---

### **🔹 Iterating Over a Map (Values Only)**
```go
	// Use underscore to ignore keys
	for _, value := range myMap {
		fmt.Println(value)
	}
```
✅ `for _, value := range myMap` → Uses `_` to ignore keys, iterating only over values.  
✅ **Expected Output (unordered):**
```
Donut
Operator
```
or  
```
Operator
Donut
```

---

### **🔹 Key Takeaways**
✅ **`for range` loops** are used to iterate over slices, arrays, and maps.  
✅ If iterating over **a slice**, the first value is **index**, and the second is **element**.  
✅ If iterating over **a map**, the first value is **key**, and the second is **value**.  
✅ Use `_` to ignore unwanted values.

Would you like a more advanced example, like iterating over a nested map? 😊