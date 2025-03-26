### **Explanation of the Code**
This Go program demonstrates how to **create a map, pass it to a function, and iterate over it** using a `for-range` loop.

---

## **1. Creating a Map**
```go
colors := map[string]string{
	"red":   "#FF000",
	"green": "#4bf745",
	"white": "#ffffff",
}
```
- A **map** named `colors` is created with `string` keys (color names) and `string` values (hex color codes).
- The map contains:
  ```go
  {
      "red": "#FF000",
      "green": "#4bf745",
      "white": "#ffffff"
  }
  ```

---

## **2. Calling the `printMap` Function**
```go
printMap(colors)
```
- The `colors` map is passed to the `printMap` function.
- Maps in Go are **passed by reference**, meaning changes inside `printMap` would affect the original `colors` map.

---

## **3. Iterating Over the Map in `printMap`**
```go
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex, "whats range?", c)
	}
}
```
- The function receives the map `c`.
- A `for-range` loop is used to **iterate over the map**, extracting:
  - `color` → the key (e.g., `"red"`, `"green"`, `"white"`).
  - `hex` → the value (e.g., `"#FF000"`, `"#4bf745"`, `"#ffffff"`).
- **Each iteration prints a key-value pair**:
  ```
  Hex code for red is #FF000 whats range? map[red:#FF000 green:#4bf745 white:#ffffff]
  Hex code for green is #4bf745 whats range? map[red:#FF000 green:#4bf745 white:#ffffff]
  Hex code for white is #ffffff whats range? map[red:#FF000 green:#4bf745 white:#ffffff]
  ```
- The extra `c` in `fmt.Println()` prints the **entire map in each iteration**.

---

## **Key Takeaways**
✅ **Maps store key-value pairs for fast lookups.**  
✅ **Maps are passed by reference in Go, so modifications inside functions affect the original map.**  
✅ **Using `for-range`, you can iterate over all key-value pairs.**  
✅ **Maps in Go are unordered, so iteration order may change each time you run the program.**  

Would you like to modify the function to return values instead of printing them? 🚀