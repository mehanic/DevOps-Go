This Go program demonstrates **concurrent JSON unmarshalling with error handling**, using **goroutines, channels, a `WaitGroup`, and `select` for synchronization**. The goal is to **parse multiple JSON strings into `Cat` structs concurrently while handling errors properly**.

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

### 🔹 **2. Initializing Channels**
```go
done := make(chan struct{})
catsCh := make(chan Cat)
errsCh := make(chan error)
```
- **`done` channel** signals when all goroutines finish.
- **`catsCh` channel** is used to send successfully parsed `Cat` structs.
- **`errsCh` channel** is used to send errors when JSON unmarshalling fails.

📌 **Rule:**  
🔄 **Channels are used to safely transfer data between goroutines.**

---

### 🔹 **3. Starting Goroutines to Process JSON Concurrently**
```go
for _, catData := range catsJSONs {
	go func(catData string) {
		var cat Cat
		if err := json.Unmarshal([]byte(catData), &cat); err != nil {
			errsCh <- err // Send error to errsCh
			return
		}
		catsCh <- cat // Send parsed Cat to catsCh
	}(catData)
}
```
- **Each JSON string is processed in a separate goroutine**.
- If **parsing succeeds**, the parsed `Cat` struct is sent to `catsCh`.
- If **parsing fails**, the error is sent to `errsCh`.

📌 **Rule:**  
🚀 **Each goroutine runs independently, parsing JSON concurrently.**

---

### 🔹 **4. Using `WaitGroup` to Track Goroutines**
```go
var wg sync.WaitGroup
wg.Add(len(catsJSONs))
```
- A `sync.WaitGroup` is created to track how many goroutines are running.
- `wg.Add(len(catsJSONs))` registers the total number of goroutines.

📌 **Rule:**  
🔄 `WaitGroup` ensures **we wait for all goroutines to finish**.

---

### 🔹 **5. Closing the `done` Channel When All Goroutines Finish**
```go
go func() {
	wg.Wait()
	close(done) // Signal that all goroutines are done
}()
```
- A separate goroutine **waits until all JSON processing finishes** (`wg.Wait()`).
- **Once all goroutines finish, it closes `done`**, which signals completion.

📌 **Rule:**  
✅ **Closing the `done` channel signals that all data has been processed.**

---

### 🔹 **6. Handling Results with `select`**
```go
for {
	select {
	case <-done:
		return

	case c := <-catsCh:
		wg.Done()
		fmt.Println(c)

	case err := <-errsCh:
		wg.Done()
		fmt.Println(err)
	}
}
```
- The **`select` statement waits for events from multiple channels**.
- **`case <-done:`** exits the loop when all goroutines are finished.
- **`case c := <-catsCh:`** prints successfully parsed cats.
- **`case err := <-errsCh:`** prints errors from failed JSON parsing.
- Calls `wg.Done()` to indicate that a goroutine's work is finished.

📌 **Rule:**  
🚀 `select` allows **handling multiple channel events concurrently**.

---

## ❌ **Problems in This Code**
### ⚠️ **Issue 1: `WaitGroup` is Added but Never Called in Goroutines**
❌ `wg.Add(len(catsJSONs))` is called, but `wg.Done()` is **not called inside the goroutines**.

✅ **Fix:** Move `wg.Done()` inside the goroutines:
```go
for _, catData := range catsJSONs {
	wg.Add(1) // Add one task per goroutine
	go func(catData string) {
		defer wg.Done() // Ensure wg.Done() is called when goroutine exits

		var cat Cat
		if err := json.Unmarshal([]byte(catData), &cat); err != nil {
			errsCh <- err
			return
		}
		catsCh <- cat
	}(catData)
}
```

---

### ⚠️ **Issue 2: Deadlock in `select` Statement**
❌ **`select` reads from `catsCh` and `errsCh`, but it doesn't close them!**  
✅ **Fix:** Close both channels in the goroutine that waits for `wg.Wait()`:
```go
go func() {
	wg.Wait()
	close(done)
	close(catsCh)
	close(errsCh)
}()
```

---

### ⚠️ **Issue 3: Goroutine Capturing Loop Variable (`catData`)**
❌ **Without explicitly passing `catData` as a function argument**, all goroutines may use the **last value of `catData`**.

✅ **Fix:** Pass `catData` explicitly:
```go
for _, catData := range catsJSONs {
	go func(catData string) {
		// Correctly captures the loop variable
	}(catData)
}
```

---

## ✅ **Fixed Version**
```go
package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Cat struct {
	Name string `json:"name"`
}

func main() {
	catsJSONs := []string{`{"name": "Bobby"}`, `"name": "Billy"`, `{"name": "Васёк"}`}

	done := make(chan struct{})
	catsCh := make(chan Cat)
	errsCh := make(chan error)

	var wg sync.WaitGroup

	for _, catData := range catsJSONs {
		wg.Add(1) // Add before starting goroutine
		go func(catData string) {
			defer wg.Done() // Ensure wg.Done() is called when goroutine exits

			var cat Cat
			if err := json.Unmarshal([]byte(catData), &cat); err != nil {
				errsCh <- err
				return
			}
			catsCh <- cat
		}(catData)
	}

	go func() {
		wg.Wait()
		close(done)
		close(catsCh)
		close(errsCh)
	}()

	for {
		select {
		case <-done:
			return

		case c, ok := <-catsCh:
			if !ok {
				continue
			}
			fmt.Println(c)

		case err, ok := <-errsCh:
			if !ok {
				continue
			}
			fmt.Println(err)
		}
	}
}
```

---

## ✅ **Final Takeaways**
| Concept | Explanation |
|---------|------------|
| **Concurrency with Goroutines** | Each JSON string is parsed in parallel. |
| **Synchronization with WaitGroup** | Ensures all goroutines finish before signaling `done`. |
| **Channels for Safe Data Transfer** | Used for passing successful `Cat` structs and errors. |
| **Error Handling with a Dedicated Channel** | Errors are captured separately. |
| **`select` Statement for Concurrent Event Handling** | Listens for `done`, `catsCh`, and `errsCh`. |
| **Loop Variable Capturing** | Always pass loop variables as function arguments in goroutines. |

🚀 **This is a great example of proper concurrency, error handling, and Go best practices!**