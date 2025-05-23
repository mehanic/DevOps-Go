This Go program demonstrates **context-based cancellation and timeout handling**. It uses `context.WithCancelCause` and `context.WithTimeoutCause` to gracefully handle **interrupt signals** and **orchestration timeouts**.  

---

## 📌 **Key Rules and Concepts**  
### 🔹 **1. Handling Interrupt Signals (`CTRL+C`)**  
```go
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, os.Interrupt) // Capture interrupt signal (CTRL+C)
		<-ch
		cancel(errProgramInterrupted) // Cancel context with a cause
	}()
```
- Listens for **OS interrupt signals (CTRL+C)**.
- Calls `cancel(errProgramInterrupted)`, propagating **cancellation cause**.

📌 **Rule:**  
🚀 Pressing `CTRL+C` cancels the program with `errProgramInterrupted`.

---

### 🔹 **2. Context Propagation with Timeout (`orchestrate`)**  
```go
func orchestrate(ctx context.Context) error {
	ctx, cancel := context.WithTimeoutCause(ctx, 3*time.Second, errOrchestrationTimeout)
	defer cancel()
	return doWork(ctx)
}
```
- Creates a **new context with a 3-second timeout**.
- Calls `doWork(ctx)`, passing the **orchestration context**.

📌 **Rule:**  
⏳ If the work exceeds **3 seconds**, `orchestrate` cancels with `errOrchestrationTimeout`.

---

### 🔹 **3. Context Propagation with Timeout (`doWork`)**  
```go
func doWork(ctx context.Context) error {
	ctx, cancel := context.WithTimeoutCause(ctx, 10*time.Second, errWorkTimeout)
	defer cancel()

	<-ctx.Done() // Wait until context is canceled
	return fmt.Errorf("aborting work: %v, because of %w", ctx.Err(), context.Cause(ctx))
}
```
- Creates a **new context with a 10-second timeout**.
- Uses `<-ctx.Done()` to **wait for cancellation**.
- **Returns error with cause** (e.g., timeout or external cancellation).

📌 **Rule:**  
⏳ If the work exceeds **10 seconds**, `doWork` cancels with `errWorkTimeout`.

---

## ✅ **Possible Outputs & What They Mean**
### **Case 1: The program runs normally for 3+ seconds → Timeout occurs**
```
aborting work: context deadline exceeded, because of orchestration timeout
```
**Explanation:**  
- `orchestrate` sets a **3-second timeout**.
- `doWork` expects **10 seconds**, but `orchestrate` cancels it after **3 seconds**.
- The **error message propagates** with `errOrchestrationTimeout`.

📌 **Rule:**  
⚠️ **If `orchestrate` cancels first, `doWork` also stops early.**

---

### **Case 2: User presses `CTRL+C` before timeout**
```
aborting work: context canceled, because of program interrupted
```
**Explanation:**  
- User **manually cancels the program (`CTRL+C`)**.
- The **cancel function is triggered** (`cancel(errProgramInterrupted)`).
- The **error message propagates** with `errProgramInterrupted`.

📌 **Rule:**  
⚠️ **User interruptions take priority over timeouts.**

---

## 🎯 **Final Takeaways**
| Concept | Explanation |
|---------|------------|
| **Graceful shutdown** | Uses `context.WithCancelCause` to stop tasks cleanly. |
| **Timeout handling** | Uses `context.WithTimeoutCause` to stop tasks if they run too long. |
| **Error propagation** | Uses `context.Cause(ctx)` to determine why a task was canceled. |
| **Signal handling** | Listens for `os.Interrupt` (`CTRL+C`) and cancels execution. |

### ✅ **Well-structured, production-ready timeout & cancellation handling in Go! 🚀**