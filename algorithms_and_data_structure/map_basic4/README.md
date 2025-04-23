This Go program finds the most frequently occurring number in a given slice of integers. Let's break it down step by step.

---

## **1. Initializing the Slice**
```go
slice := []int{1,1,4,5,3,3,3,3,1,2}
```
- This slice contains multiple integers, some of which repeat.

---

## **2. Counting Occurrences Using a Map**
```go
x_map := make(map[int]int)
```
- Creates a map `x_map` where:
  - The **keys** represent numbers from the slice.
  - The **values** represent how many times each number appears.

```go
for _, el := range slice {
    x_map[el]++
}
```
- Iterates over `slice`, and for each element `el`:
  - Increments its count in `x_map`.

### **Example: After the loop runs, `x_map` contains:**
```go
map[1:3 2:1 3:4 4:1 5:1]
```
This means:
- `1` appears **3** times.
- `2` appears **once**.
- `3` appears **4** times.
- `4` appears **once**.
- `5` appears **once**.

---

## **3. Finding the Most Frequent Number**
```go
max_key := 0
max := 0
```
- `max_key` stores the most frequently occurring number.
- `max` stores the highest frequency.

```go
for key, val := range x_map {
    if max < val {
        max = val
        max_key = key
    }
}
```
- Iterates through `x_map` to find the number with the highest count (`max`).

### **Example:**
- Starts with `max = 0`
- Checks each key-value pair:
  - `1 → 3 times` → Updates `max=3`, `max_key=1`
  - `2 → 1 time` (no change)
  - `3 → 4 times` → Updates `max=4`, `max_key=3`
  - `4 → 1 time` (no change)
  - `5 → 1 time` (no change)
  
At the end, the most frequent number is **3**, appearing **4 times**.

---

## **4. Printing the Result**
```go
fmt.Printf("Eng ko'p takrorlangan son %d %d marta takrorlangan\n", max_key, max)
```
- Prints:  
  ```
  Eng ko'p takrorlangan son 3 4 marta takrorlangan
  ```
  (Uzbek for: "The most repeated number is 3, which appears 4 times.")

---

## **Summary**
| Concept | Explanation |
|---------|------------|
| **Maps for Frequency Count** | `x_map[el]++` to count occurrences |
| **Finding Max** | Iterates `x_map` to track max value |
| **Printing Output** | Uses `fmt.Printf` for formatted output |

This approach runs in **O(n) time complexity**, making it efficient! 🚀