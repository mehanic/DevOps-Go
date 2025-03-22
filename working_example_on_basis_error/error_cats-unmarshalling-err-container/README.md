This Go program demonstrates **concurrent JSON unmarshalling with structured error handling** using **goroutines, channels, `sync.WaitGroup`, and a custom struct (`CatContainer`) for error management**.

---

## 📌 **Key Rules and Concepts**

### 🔹 **1. Defining the `Cat` Struct**
```go
type Cat struct {
	Name string `json:"name"`
}
```
- The `Cat` struct represents a cat with a **single field `Name`**.
- Uses a **JSON struct tag (`json:"name"`)** to map the `"name"` field in JSON to `Name`.

📌 **Rule:**  
✅ `json.Unmarshal()` automatically assigns JSON values to struct fields.

---

### 🔹 **2. Creating `CatContainer` for Handling Errors**
```go
type CatContainer struct {
	Cat Cat
	Err error // Stores error if JSON parsing fails.
}
```
- `CatContainer` is a **wrapper struct** that stores:
  - A `Cat` struct (if JSON parsing succeeds).
  - An `error` field (if parsing fails).
- **Instead of using a separate error channel, this struct holds both values**.

📌 **Rule:**  
✅ **Structs can be used to bundle a value and an error together**.

---

### 🔹 **3. Initializing a Channel for Parsed Cats**
```go
catsCh := make(chan CatContainer)
```
- This **channel carries `CatContainer` values**.
- It **transfers both successfully parsed cats and errors** from goroutines.

📌 **Rule:**  
🔄 **Channels provide safe concurrent communication between goroutines**.

---

### 🔹 **4. Spawning Goroutines to Process JSON Strings**
```go
var wg sync.WaitGroup
wg.Add(len(catsJSONs)) // Add the number of JSON strings

for _, catData := range catsJSONs {
	go func(catData string) {
		defer wg.Done() // Mark this goroutine as done when finished.

		var cat Cat
		if err := json.Unmarshal([]byte(catData), &cat); err != nil {
			catsCh <- CatContainer{Err: err} // Send an error if parsing fails.
			return
		}
		catsCh <- CatContainer{Cat: cat} // Send the parsed Cat struct.
	}(catData)
}
```
- **Each JSON string is processed in a separate goroutine**.
- If **parsing succeeds**, it sends a `CatContainer{Cat: cat}` to `catsCh`.
- If **parsing fails**, it sends a `CatContainer{Err: err}` to `catsCh`.
- Uses **`defer wg.Done()`** to ensure each goroutine signals completion.

📌 **Rule:**  
🚀 **Each goroutine runs independently, parsing JSON concurrently.**

---

### 🔹 **5. Closing the `catsCh` Channel When All Goroutines Finish**
```go
go func() {
	wg.Wait()
	close(catsCh) // Close channel after all goroutines finish.
}()
```
- A separate goroutine **waits for all JSON processing goroutines to finish**.
- **Once all goroutines finish, it closes `catsCh`**, signaling completion.

📌 **Rule:**  
✅ **Closing the `catsCh` channel ensures no more data is sent, preventing deadlocks.**

---

### 🔹 **6. Reading Results from `catsCh`**
```go
for catContainer := range catsCh {
	if catContainer.Err != nil {
		fmt.Println("ERROR:", catContainer.Err)
		continue
	}
	fmt.Println(catContainer.Cat)
}
```
- The **loop reads values from `catsCh`** until it's closed.
- If **`catContainer.Err` is non-nil**, an error occurred, and it's printed.
- Otherwise, the successfully parsed `Cat` struct is printed.

📌 **Rule:**  
🔄 **Reading from a closed channel returns no more values, ending the loop.**

---

## 🚨 **Output & Explanation**
```
ERROR: invalid character ':' after top-level value
{Васёк}
{Bobby}
```
### 🔹 **What Happened?**
1. **Parsing Success** ✅  
   - `"{"name": "Bobby"}"` → Successfully parsed → `{Bobby}`
   - `"{"name": "Васёк"}"` → Successfully parsed → `{Васёк}`  

2. **Parsing Failure** ❌  
   - `"name": "Billy"` → **Invalid JSON format** (missing curly braces `{}`)  
   - `json.Unmarshal()` returns an error:  
     `invalid character ':' after top-level value`

---

## ✅ **Key Takeaways**
| Concept | Explanation |
|---------|------------|
| **Goroutines for Parallel Processing** | Each JSON string is processed in a separate goroutine. |
| **WaitGroup for Synchronization** | Ensures all goroutines finish before closing `catsCh`. |
| **Using a Struct for Errors (`CatContainer`)** | Bundles a `Cat` struct with an error for structured handling. |
| **Closing Channels to Signal Completion** | Prevents deadlocks by closing `catsCh` when all goroutines finish. |
| **Looping Over a Channel (`range catsCh`)** | Automatically stops when the channel is closed. |

🚀 **This is an excellent example of concurrency, error handling, and Go best practices!**