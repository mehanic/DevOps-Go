### **Explanation of the Code**

This Go program demonstrates various ways to **declare, initialize, and use maps** in Go. Here's a detailed breakdown of each section.

---

### **1️⃣ Declaring a Map with the `var` Keyword (Without Initialization)**

```go
var map1 map[string]int
fmt.Println("Map declared with var keyword without initialization:", map1)
```

- **Map Declaration**: A map is declared using the `var` keyword without initializing it.
  - `map1` is a **nil map** of type `map[string]int` (key is a string, value is an integer).
- **Nil Map**: This map is **not initialized**, meaning it has a `nil` value. Trying to use this map without initializing it (like adding elements) will cause a runtime panic.
  
**Output**:
```
Map declared with var keyword without initialization: map[]
```

---

### **2️⃣ Declaring a Map with a Map Literal**

```go
map2 := map[string]int{"apple": 5, "banana": 7}
fmt.Println("Map declared with a map literal:", map2)
```

- **Map Literal**: A map can be initialized directly using a **map literal**.
  - Here, `map2` is initialized with two key-value pairs: `"apple": 5` and `"banana": 7`.
- **Instant Initialization**: This method is convenient when you know the values upfront.

**Output**:
```
Map declared with a map literal: map[apple:5 banana:7]
```

---

### **3️⃣ Declaring a Map Using the `make` Function**

```go
map3 := make(map[string]int)
map3["apple"] = 5
map3["banana"] = 7
fmt.Println("Map declared with make function:", map3)
```

- **`make()` Function**: You can use the `make()` function to create a map that is initialized but empty.
  - `map3` is created as an empty map of type `map[string]int`.
  - After creation, we add two key-value pairs to it (`"apple": 5` and `"banana": 7`).
  
**Output**:
```
Map declared with make function: map[apple:5 banana:7]
```

---

### **4️⃣ Declaring a Map Using the `make` Function with an Initial Capacity**

```go
map4 := make(map[string]int, 10) // 10 is the initial size
map4["apple"] = 5
map4["banana"] = 7
fmt.Println("Map declared with make function and specific size:", map4)
```

- **`make()` with Initial Size**: The `make()` function can also take a **capacity** parameter.
  - The capacity specifies the initial size of the map, which can optimize memory allocation for larger maps.
  - `map4` is initialized with a capacity of 10, meaning it can hold up to 10 elements before needing to reallocate memory.
- **Note**: The **capacity** is a performance optimization, not a hard limit on the number of elements that can be added.

**Output**:
```
Map declared with make function and specific size: map[apple:5 banana:7]
```

---

### **5️⃣ Using Structs as Map Values**

```go
type Fruit struct {
	Price int
	Color string
}

map5 := make(map[string]Fruit)
map5["apple"] = Fruit{Price: 5, Color: "green"}
map5["banana"] = Fruit{Price: 7, Color: "yellow"}
fmt.Println("Map with structs as values:", map5)
```

- **Using Structs as Map Values**: Instead of simple types like `int` or `string`, you can use structs as map **values**.
  - `Fruit` is a struct that has two fields: `Price` (integer) and `Color` (string).
  - `map5` is a map where the key is a string (fruit name), and the value is a `Fruit` struct.
  - We add two entries to the map (`"apple"` and `"banana"`), each associated with a `Fruit` struct.

**Output**:
```
Map with structs as values: map[apple:{5 green} banana:{7 yellow}]
```

---

### **6️⃣ Nested Maps**

```go
map6 := make(map[string]map[string]int)
map6["fruit"] = map[string]int{"apple": 5, "banana": 7}
map6["vegetable"] = map[string]int{"carrot": 3, "peas": 2}
fmt.Println("Nested map:", map6)
```

- **Nested Maps**: You can have a map where the **values are maps** themselves, creating a nested map structure.
  - `map6` is a map where each key is a string (like `"fruit"`, `"vegetable"`), and the value is another map.
  - The nested maps (`map[string]int`) hold key-value pairs for specific items (like `"apple": 5`, `"banana": 7`).

**Output**:
```
Nested map: map[fruit:map[apple:5 banana:7] vegetable:map[carrot:3 peas:2]]
```

---

### **Summary of Map Declaration and Initialization**

1. **`var` keyword**: Declares a **nil** map (uninitialized).
2. **Map Literal**: Initializes the map with predefined key-value pairs.
3. **`make()` function**: Initializes a map with an optional initial size for performance optimization.
4. **Using Structs**: You can use structs as values in the map, allowing more complex data structures.
5. **Nested Maps**: You can create maps where values are themselves maps, supporting hierarchical data.

---

### **Would you like more details or further examples on any of these map concepts?**