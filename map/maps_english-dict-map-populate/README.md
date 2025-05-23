This Go program implements an **English-to-Turkish dictionary** using **maps** (`map[string]string`) for fast lookups. It retrieves translations from **command-line arguments** and prints the result.

---

## **🔹 Code Breakdown**

### **1️⃣ Retrieving Command-Line Arguments**
```go
args := os.Args[1:]
if len(args) != 1 {
    fmt.Println("[english word] -> [turkish word]")
    return
}
query := args[0]
```
- **`os.Args[1:]`** gets input from the command line (ignoring the program name).
- If no word (or more than one) is provided, it **prints a usage message** and exits.
- The input word is stored in `query`.

---

### **2️⃣ Defining a Dictionary (Map)**
```go
dict := map[string]string{
    "good":    "kötü",      // (initially incorrect, but gets updated)
    "great":   "harika",
    "perfect": "mükemmel",
}
```
- Defines a **map** `dict` that stores **English words as keys** and their **Turkish translations as values**.
- `"good"` is initially set to `"kötü"` (incorrect), but this is later **overwritten**.

---

### **3️⃣ Adding & Overwriting Entries**
```go
dict["up"] = "yukarı"   // adds a new entry
dict["down"] = "aşağı"  // adds a new entry
dict["good"] = "iyi"    // overwrites the previous value
dict["mistake"] = ""    // adds a key with an empty value
```
- `"good"` is updated from `"kötü"` to `"iyi"`.
- `"mistake"` has an **empty string value**, which is considered a **zero-value** in Go.

---

### **4️⃣ Searching for the Input Word (O(1) Lookup)**
```go
if value, ok := dict[query]; ok {
    fmt.Printf("%q means %#v\n", query, value)
    return
}
fmt.Printf("%q not found.\n", query)
```
- Uses **map lookup (`dict[query]`)**, which has **O(1) time complexity**.
- The **comma-ok idiom** (`value, ok := dict[query]`) checks:
  - If `query` exists in the map (`ok == true`), prints its value.
  - If it **doesn't exist** (`ok == false`), prints `"not found"`.

---

### **5️⃣ Printing Total Number of Keys**
```go
fmt.Printf("# of Keys : %d\n", len(dict))
```
- `len(dict)` returns the **number of key-value pairs** in the map.

---

### **🔹 Additional Features (Commented Code)**
1. **Comparing Two Maps**
   ```go
   copied := map[string]string{ ... } // Same key-value pairs as dict
   if fmt.Sprintf("%s", dict) == fmt.Sprintf("%s", copied) {
       fmt.Println("Maps are equal")
   }
   ```
   - Converts maps to strings and compares them (since Go **does not allow direct map comparisons**).

2. **Iterating Over the Map**
   ```go
   for k, v := range dict {
       fmt.Printf("%q means %#v\n", k, v)
   }
   ```
   - Prints all key-value pairs **in insertion order** (Go 1.12+).

3. **Retrieving Values Directly**
   ```go
   fmt.Println("good      -> ", dict["good"])
   fmt.Println("great     -> ", dict["great"])
   fmt.Println("perfect   -> ", dict["perfect"])
   ```
   - Directly accesses values using keys.

---

## **🔹 Example Usage**
### ✅ **When Word Exists**
```bash
$ go run main.go good
"good" means "iyi"
```
### ❌ **When Word Doesn't Exist**
```bash
$ go run main.go bad
"bad" not found.
```
### ⚠️ **When No Input Is Given**
```bash
$ go run main.go
[english word] -> [turkish word]
```

---

## **🔹 Key Takeaways**
✅ **Uses a map for fast O(1) lookups.**  
✅ **Handles missing words properly using `ok` check.**  
✅ **Supports dynamic updates to the dictionary.**  
✅ **Prevents overwriting errors by explicitly updating keys.**  

Would you like me to improve this with **case-insensitive search** or **JSON input support**? 🚀