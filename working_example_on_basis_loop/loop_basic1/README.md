### **Go Loops and Control Statements Explained**

This Go program demonstrates different types of loops (`for` loops, `range` loops) and control flow mechanisms (`break`, `continue`, `goto`). Let's break it down step by step.

---

### **1. Declaring and Printing an Array (Slice)**
```go
days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
fmt.Println(days)
```
- **Creates a slice** `days` with 5 elements (not including Monday and Thursday).
- **Prints** `["Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"]`.

---

### **2. Different Ways to Loop Over a Slice**

#### **Loop Type 1: Traditional `for` Loop (C-style)**
```go
for d := 0; d < len(days); d++ {
    fmt.Println(days[d])
}
```
- **Loops from `0` to `len(days)-1`**.
- Uses `days[d]` to print each value.

#### **Loop Type 2: Using `range` with Index**
```go
for i := range days {
    fmt.Println(days[i])
}
```
- **Iterates over indexes only** (`range` returns only index if no second variable is used).
- Similar to `for d := 0; d < len(days); d++`.

#### **Loop Type 3: Using `range` with Index and Value**
```go
for index, day := range days {
    fmt.Printf("Index is %v and value is %v\n", index, day)
}
```
- **Gets both index and value** directly from `range`.
- More efficient than `days[i]`, since it avoids additional slice lookup.

---
### **3. Using `for` with a Condition (Like `while` Loop)**
```go
rogueValue := 1
for rogueValue < 10 {
```
- Acts like a **`while` loop**, runs while `rogueValue < 10`.

---

### **4. `goto` Statement**
```go
if rogueValue == 2 {
    goto lco
}
```
- If `rogueValue == 2`, execution **jumps** to the `lco:` label.
- **Avoid using `goto` unless absolutely necessary** (hard to debug and read).

```go
lco:
fmt.Println("Jumping at LearnCodeOnline")
```
- This line executes when `goto lco` is triggered.
- **Output:** `"Jumping at LearnCodeOnline"`

---

### **5. `break` Statement**
```go
if rogueValue == 5 {
    break
}
```
- **Stops the loop completely** when `rogueValue == 5`.
- The loop will **not execute for values `5 to 9`**.

---

### **6. `continue` Statement**
```go
if anotherValue == 5 {
    anotherValue++
    continue
}
```
- When `anotherValue == 5`, **it skips the rest of the loop body** and goes to the next iteration.
- `anotherValue++` is needed to **avoid an infinite loop** (otherwise, it would stay at `5` forever).

---

## **Final Output Breakdown**
```
Welcome to loops in GoLang
[Sunday Tuesday Wednesday Friday Saturday]
```
1. **Prints slice `days`**.

```
Loops example 1:
Sunday
Tuesday
Wednesday
Friday
Saturday
```
2. **Loops over `days` using `for d := 0; d < len(days); d++`**.

```
Loops example 2:
Sunday
Tuesday
Wednesday
Friday
Saturday
```
3. **Loops over `days` using `for i := range days`**.

```
Loops example 3:
Index is 0 and value is Sunday
Index is 1 and value is Tuesday
Index is 2 and value is Wednesday
Index is 3 and value is Friday
Index is 4 and value is Saturday
```
4. **Loops using `for index, day := range days`**.

```
Value is:  1
Jumping at LearnCodeOnline
```
5. **Loop with `goto` skips values and prints the jump statement**.

```
Value is:  1
Value is:  2
Value is:  3
Value is:  4
Value is:  6
Value is:  7
Value is:  8
Value is:  9
```
6. **Loop with `continue` skips printing `5` but continues execution**.

---

## **Key Takeaways**
✅ **Use `for range` to iterate over slices and maps efficiently**.  
✅ **Use `break` to exit a loop immediately**.  
✅ **Use `continue` to skip an iteration without breaking the loop**.  
✅ **Avoid `goto` unless absolutely necessary** (can make code harder to read).  
✅ **`for` with a condition can act like a `while` loop**.  

Would you like a modified version of this program without `goto`? 🚀