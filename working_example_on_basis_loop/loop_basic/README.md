This **Go program** demonstrates different types of loops using the `for` statement. Let's break it down step by step:

---

## **1. Infinite `for` Loop with `break`**
```go
for {
    fmt.Println("Loop iteration")
    break
}
```
- **`for {}`** → An infinite loop (keeps running indefinitely).
- **`break`** → Immediately exits the loop after one iteration.
- **Output:** `"Loop iteration"`

---

## **2. `for` Loop with a Boolean Condition**
```go
isRun := true
for isRun {
    fmt.Println("Loop iteration with condition")
    isRun = false
}
```
- The loop runs while `isRun` is **true**.
- After one iteration, `isRun = false`, so the loop exits.
- **Output:** `"Loop iteration with condition"`

---

## **3. `for` Loop with `continue`**
```go
for i := 0; i < 2; i++ {
    fmt.Println("Loop iteration", i)
    if i == 1 {
        continue
    }
}
```
- **`i := 0; i < 2; i++`** → A classic C-style `for` loop.
- **`continue`** → Skips the rest of the loop body when `i == 1`, but since it's the last iteration, it has no real effect here.
- **Output:**
  ```
  Loop iteration 0
  Loop iteration 1
  ```

---

## **4. While-style Loop (`for` with a condition)**
```go
sl := []int{1, 2, 3}
idx := 0

for idx < len(sl) {
    fmt.Println("while-stype loop idx:", idx, "value:", sl[idx])
    idx++
}
```
- This mimics a **`while` loop** (`for` in Go can act like `while` in other languages).
- Iterates while `idx < len(sl)`, printing values from a slice.
- **Output:**
  ```
  while-stype loop idx: 0 value: 1
  while-stype loop idx: 1 value: 2
  while-stype loop idx: 2 value: 3
  ```

---

## **5. Classic `for` Loop (C-style)**
```go
for i := 0; i < len(sl); i++ {
    fmt.Println("c-style loop", i, sl[i])
}
```
- A traditional `for` loop iterating through a slice (`sl`).
- **Output:**
  ```
  c-style loop 0 1
  c-style loop 1 2
  c-style loop 2 3
  ```

---

## **6. `for range` Loop with Slices**
```go
for idx := range sl {
    fmt.Println("range slice by index", idx)
}
```
- Iterates over the **indexes** of `sl`.
- **Output:**
  ```
  range slice by index 0
  range slice by index 1
  range slice by index 2
  ```

```go
for idx, val := range sl {
    fmt.Println("range slice by idx-value", idx, val)
}
```
- Iterates over **index-value pairs**.
- **Output:**
  ```
  range slice by idx-value 0 1
  range slice by idx-value 1 2
  range slice by idx-value 2 3
  ```

---

## **7. `for range` Loop with Maps**
```go
profile := map[int]string{1: "Hoasker", 2: "RZ"}

for key := range profile {
    fmt.Println("range map by key", key)
}
```
- Iterates over **keys** in a `map`.
- **Output (order may vary):**
  ```
  range map by key 1
  range map by key 2
  ```

```go
for key, val := range profile {
    fmt.Println("range map by key-val", key, val)
}
```
- Iterates over **key-value pairs**.
- **Output:**
  ```
  range map by key-val 1 Hoasker
  range map by key-val 2 RZ
  ```

```go
for _, val := range profile {
    fmt.Println("range map by val", val)
}
```
- Uses `_` to **ignore the key**, printing only values.
- **Output:**
  ```
  range map by val Hoasker
  range map by val RZ
  ```

---

## **8. Iterating Over Strings (`for range` on `string`)**
```go
str := "Hello, World!"
for pos, char := range str {
    fmt.Printf("%#U at pos %d\n", char, pos)
}
```
- Iterates over **each Unicode character** in the string.
- **`char` is a `rune` (Unicode code point)**.
- **`%#U`** prints the character in Unicode format.
- **Output:**
  ```
  U+0048 'H' at pos 0
  U+0065 'e' at pos 1
  U+006C 'l' at pos 2
  U+006C 'l' at pos 3
  U+006F 'o' at pos 4
  U+002C ',' at pos 5
  U+0020 ' ' at pos 6
  U+0057 'W' at pos 7
  U+006F 'o' at pos 8
  U+0072 'r' at pos 9
  U+006C 'l' at pos 10
  U+0064 'd' at pos 11
  U+0021 '!' at pos 12
  ```

---

## **Summary of Loop Types in Go**
| Loop Type | Example | Description |
|-----------|---------|-------------|
| Infinite Loop | `for {}` | Runs forever unless `break` is used |
| Condition-based Loop | `for isRun {}` | Mimics a `while` loop |
| C-style Loop | `for i := 0; i < n; i++ {}` | Classic loop with initialization, condition, and increment |
| `while`-style Loop | `for idx < len(sl) {}` | Runs while a condition holds |
| `for range` (slice) | `for idx, val := range sl {}` | Iterates over slice indexes and values |
| `for range` (map) | `for key, val := range profile {}` | Iterates over map keys and values |
| `for range` (string) | `for pos, char := range str {}` | Iterates over Unicode characters |

---

### **Key Takeaways**
✅ **Go has only one loop type: `for`**, but it can be used in various ways.  
✅ **`for range` is best for iterating over slices, maps, and strings**.  
✅ **`continue` skips the rest of the current loop iteration**.  
✅ **`break` exits a loop immediately**.  
✅ **Scoped variables in loops (`idx, val := range sl`) exist only inside the loop**.  

Let me know if you need further clarification! 🚀