This Go program demonstrates **concurrent JSON unmarshalling** using **goroutines, channels, and a WaitGroup**. The program tries to **parse JSON strings into `Cat` structs in parallel** while handling potential errors.

---

## 📌 **Key Rules and Concepts**

### 🔹 **1. Defining the Struct (`Cat`)**
```go
type Cat struct {
	Name string `json:"name"`
}
```
- **`Cat` struct** represents a cat with a single field: `Name`.
- Uses a **JSON struct tag (`json:"name"`)** to map JSON fields.

📌 **Rule:**  
✅ `json.Unmarshal()` automatically assigns the `"name"` field from JSON to the `Name` field in `Cat`.

---

### 🔹 **2. Defining the JSON Data**
```go
catsJSONs := []string{`{"name": "Bobby"}`, `"name": "Billy"`, `{"name": "Васёк"}`}
```
- **A slice of JSON strings** represents different cats.
- The second entry (`"name": "Billy"`) is **invalid JSON** (missing `{}`).

📌 **Rule:**  
⚠️ **Malformed JSON will cause `json.Unmarshal()` to fail**.

---

### 🔹 **3. Using a Channel to Collect Results (`catsCh`)**
```go
catsCh := make(chan Cat)
```
- A **channel (`catsCh`) is used to pass parsed cats** between goroutines.

📌 **Rule:**  
🔄 **Channels allow safe concurrent data sharing** between goroutines.

---

### 🔹 **4. Processing JSON Concurrently Using Goroutines**
```go
var wg sync.WaitGroup
wg.Add(len(catsJSONs)) // Add the number of goroutines

for _, catData := range catsJSONs {
	go func(catData string) { // Goroutine for each JSON
		defer wg.Done()

		var cat Cat
		if err := json.Unmarshal([]byte(catData), &cat); err != nil {
			// Error occurs, but nothing is done (bad practice)
			return // Prevents sending an empty Cat struct
		}
		catsCh <- cat
	}(catData)
}
```
- **A goroutine is created for each JSON string**.
- Uses `json.Unmarshal()` to parse the JSON string into a `Cat` struct.
- If an error occurs (`err != nil`), the function **exits early** without sending a value to the channel.
- Each goroutine **calls `wg.Done()` when finished**.

📌 **Rule:**  
🚀 **Each JSON string is parsed in parallel** for efficiency.  

---

### 🔹 **5. Closing the Channel Once All Goroutines Finish**
```go
go func() {
	wg.Wait() // Wait for all goroutines
	close(catsCh) // Close the channel after all data is sent
}()
```
- A **separate goroutine waits for all JSON processing to complete**.
- Calls `close(catsCh)` to signal **no more data** will be sent.

📌 **Rule:**  
✅ **Closing a channel prevents deadlocks** when reading from it.

---

### 🔹 **6. Reading from the Channel**
```go
for cat := range catsCh {
	fmt.Println(cat)
}
```
- **Reads all parsed `Cat` structs from the channel**.
- Automatically stops when the channel **is closed**.

📌 **Rule:**  
🔄 **A `for` loop over a channel continues until the channel is closed**.

---

## ❌ **Potential Issues & Fixes**
### ⚠️ **Issue 1: No Error Handling**
❌ If `json.Unmarshal()` fails, the goroutine exits, but no message is logged.

✅ **Fix: Log the error**
```go
if err := json.Unmarshal([]byte(catData), &cat); err != nil {
	fmt.Println("Error unmarshalling:", err)
	return
}
```

---

### ⚠️ **Issue 2: Sending Empty Structs on Error**
❌ If `json.Unmarshal()` fails, an **empty `Cat{}`** may be sent to the channel.

✅ **Fix: Skip sending on error**
```go
if err := json.Unmarshal([]byte(catData), &cat); err != nil {
	return // Avoid sending empty struct
}
```

---

### ⚠️ **Issue 3: Goroutine Capturing Loop Variable (`catData`)**
❌ `go func(catData string) { ... }(catData)` correctly **passes `catData` as an argument**.  
✅ If written incorrectly as `go func() { ... }()`, **all goroutines would use the last value of `catData`**.

📌 **Rule:**  
🚀 **Always pass loop variables as function arguments in goroutines!**  

---

## ✅ **Final Takeaways**
| Concept | Explanation |
|---------|------------|
| **Concurrency with Goroutines** | Each JSON string is parsed in parallel. |
| **Synchronization with WaitGroup** | Ensures all goroutines finish before closing the channel. |
| **Channels for Safe Data Transfer** | Goroutines send parsed `Cat` structs through `catsCh`. |
| **Error Handling is Essential** | The program ignores invalid JSON instead of logging errors. |
| **Loop Variable Capturing** | Always pass variables as function arguments in goroutines. |

🚀 **A great example of concurrency, error handling, and Go best practices!**