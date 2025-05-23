This Go program implements a **bidirectional dictionary** that translates words between **English and Turkish**. It processes a list of **test queries** and looks up each query in both **English-to-Turkish** and **Turkish-to-English** mappings.

---

## **🔹 Code Breakdown**
### **1️⃣ Define Example Queries**
```go
testQueries := []string{"good", "iyi", "perfect", "harika", "notfound"}
```
- A **slice of test words** to check if the translation exists in either direction.

---

### **2️⃣ Create English-to-Turkish Dictionary**
```go
dict := map[string]string{
    "good":    "iyi",
    "great":   "harika",
    "perfect": "mükemmel",
    "awesome": "mükemmel",
}
```
- Defines a **map** called `dict` with English words as **keys** and their Turkish equivalents as **values**.

---

### **3️⃣ Delete a Specific Entry**
```go
delete(dict, "awesome")
```
- Removes `"awesome": "mükemmel"` from `dict` for demonstration purposes.

---

### **4️⃣ Create a Reverse Turkish-to-English Dictionary**
```go
turkish := make(map[string]string, len(dict))
for k, v := range dict {
    turkish[v] = k
}
```
- Creates an **empty map** `turkish` for Turkish-to-English translations.
- Iterates over `dict` and swaps **key-value pairs** to build a **reverse lookup**.

🔹 **Example result:**  
```go
turkish = map[string]string{
    "iyi":       "good",
    "harika":    "great",
    "mükemmel":  "perfect",
}
```

---

### **5️⃣ Process Each Query**
```go
for _, query := range testQueries {
    fmt.Printf("Query: %q\n", query)
```
- Loops through `testQueries` to check if each word exists in **either dictionary**.

---

### **6️⃣ Lookup in English-to-Turkish Dictionary**
```go
if value, ok := dict[query]; ok {
    fmt.Printf("English: %q -> Turkish: %q\n\n", query, value)
    continue
}
```
- Checks if `query` exists in `dict` (English-to-Turkish).
- If found, prints the translation and moves to the **next query** (`continue`).

---

### **7️⃣ Lookup in Turkish-to-English Dictionary**
```go
if value, ok := turkish[query]; ok {
    fmt.Printf("Turkish: %q -> English: %q\n\n", query, value)
    continue
}
```
- If not found in `dict`, checks if `query` exists in `turkish` (Turkish-to-English).
- If found, prints the translation and moves to the **next query**.

---

### **8️⃣ Handle Missing Words**
```go
fmt.Printf("Sorry, %q not found.\n\n", query)
```
- If the query **does not exist** in either dictionary, prints an error message.

---

## **🔹 Example Output**
```bash
Query: "good"
English: "good" -> Turkish: "iyi"

Query: "iyi"
Turkish: "iyi" -> English: "good"

Query: "perfect"
English: "perfect" -> Turkish: "mükemmel"

Query: "harika"
Turkish: "harika" -> English: "great"

Query: "notfound"
Sorry, "notfound" not found.
```
🔹 **Key Observations:**
- **"good" translates to "iyi"** using `dict`.
- **"iyi" translates back to "good"** using `turkish`.
- **"notfound" does not exist** in either dictionary, so an error is displayed.

---

## **🔹 Key Takeaways**
✅ **Maps provide fast O(1) lookups** for key-value pairs.  
✅ **Reverse dictionaries** can be created dynamically for bidirectional lookup.  
✅ **Checking for a key in a map** is done using `value, ok := map[key]`.  
✅ **Maps in Go do not guarantee order**, so output may vary.  

Would you like me to modify it to include **more translations** or **handle case-insensitive input**? 🚀