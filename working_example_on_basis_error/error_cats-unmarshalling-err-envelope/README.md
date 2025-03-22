This Go program demonstrates **concurrent JSON unmarshalling with structured error handling**, using **goroutines, channels, `sync.WaitGroup`, and encapsulation via `CatEnvelope`**.

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

### 🔹 **2. Creating `CatEnvelope` for Safe Error Handling**
```go
type CatEnvelope struct {
	cat Cat
	err error // Stores error if JSON parsing fails.
}
```
- `CatEnvelope` **bundles**:
  - A `Cat` struct (if JSON parsing succeeds).
  - An `error` (if parsing fails).
- This **encapsulation pattern prevents direct modification of original values**.

📌 **Rule:**  
✅ **Encapsulation helps control access to data and avoid unintended modifications.**

---

### 🔹 **3. Constructor Functions for `CatEnvelope`**
```go
func NewCatEnvelope(c Cat) CatEnvelope {
	return CatEnvelope{cat: c}
}

func NewCatEnvelopeWithErr(err error) CatEnvelope {
	return CatEnvelope{err: err}
}
```
- `NewCatEnvelope(c Cat) → CatEnvelope`  
  - Creates a `CatEnvelope` **with a valid `Cat`**.
- `NewCatEnvelopeWithErr(err error) → CatEnvelope`  
  - Creates a `CatEnvelope` **with an error**.

📌 **Rule:**  
✅ **Factory functions provide controlled object creation.**

---

### 🔹 **4. Implementing `Unpack()` for Safe Access**
```go
func (c CatEnvelope) Unpack() (Cat, error) {
	return c.cat, c.err
}
```
- **Provides controlled access** to `CatEnvelope` contents.
- Returns both the `Cat` struct and the `error`.

📌 **Rule:**  
✅ **Encapsulation ensures immutability and controlled access.**

---

### 🔹 **5. Initializing a Channel for Parsed Cats**
```go
catsCh := make(chan CatEnvelope)
```
- This **channel carries `CatEnvelope` values**.
- It **transfers both successfully parsed cats and errors** from goroutines.

📌 **Rule:**  
🔄 **Channels provide safe concurrent communication between goroutines**.

---

### 🔹 **6. Spawning Goroutines to Process JSON Strings**
```go
var wg sync.WaitGroup
wg.Add(len(catsJSONs)) // Add the number of JSON strings

for _, catData := range catsJSONs {
	go func(catData string) {
		defer wg.Done() // Mark this goroutine as done when finished.

		var cat Cat
		if err := json.Unmarshal([]byte(catData), &cat); err != nil {
			catsCh <- NewCatEnvelopeWithErr(err) // Send an error if parsing fails.
			return
		}
		catsCh <- NewCatEnvelope(cat) // Send the parsed Cat struct.
	}(catData)
}
```
- **Each JSON string is processed in a separate goroutine**.
- If **parsing succeeds**, it sends `NewCatEnvelope(cat)` to `catsCh`.
- If **parsing fails**, it sends `NewCatEnvelopeWithErr(err)` to `catsCh`.
- Uses **`defer wg.Done()`** to ensure each goroutine signals completion.

📌 **Rule:**  
🚀 **Each goroutine runs independently, parsing JSON concurrently.**

---

### 🔹 **7. Closing the `catsCh` Channel When All Goroutines Finish**
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

### 🔹 **8. Reading and Processing Results from `catsCh`**
```go
for catEnvelope := range catsCh {
	// We can modify local copies after unpacking without affecting the original.
	c, err := catEnvelope.Unpack()
	c = Cat{Name: "Hacked"}
	err = errors.New("hacked")

	// Unpacking again to show the original remains unchanged.
	c, err = catEnvelope.Unpack()
	if err != nil {
		fmt.Println("ERROR:", err)
		continue
	}
	fmt.Println(c)
}
```
- Each `CatEnvelope` is **unpacked safely**.
- The **copies of `c` and `err` can be modified without affecting the original**.
- This ensures the **original `CatEnvelope` contents remain unchanged**.

📌 **Rule:**  
✅ **Encapsulation prevents unintended modifications to original data.**

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
| **Encapsulation with `CatEnvelope`** | Prevents unintended modifications and ensures safe data access. |
| **Unpacking Function `Unpack()`** | Provides controlled, immutable access to data. |
| **Closing Channels to Signal Completion** | Prevents deadlocks by closing `catsCh` when all goroutines finish. |
| **Looping Over a Channel (`range catsCh`)** | Automatically stops when the channel is closed. |

🚀 **This is an excellent example of concurrency, encapsulation, and Go best practices!**