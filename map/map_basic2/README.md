This Go program demonstrates **maps (dictionaries), adding and deleting key-value pairs, checking if a key exists, and initializing maps with values**. Let's go through it step by step.

---

## **1. Creating and Populating a Map**
```go
sozluk := make(map[string]string)
sozluk["chair"] = "sandalye"
sozluk["fury"] = "ofke"
sozluk["car"] = "araba"
```
- **`make(map[string]string)`** creates an empty map where:
  - **Keys are `string`** (English words),
  - **Values are `string`** (Turkish translations).
- Adds three key-value pairs:
  - `"chair" → "sandalye"`
  - `"fury" → "ofke"`
  - `"car" → "araba"`

---

## **2. Accessing a Value by Key**
```go
fmt.Println(sozluk["chair"])
```
- Retrieves the value for `"chair"`, which is `"sandalye"`.

**Output:**
```
sandalye
```

---

## **3. Getting the Length of the Map**
```go
fmt.Println(len(sozluk))
```
- **`len(sozluk)`** returns the number of key-value pairs in the map.
- Since we added 3 words, it prints `3`.

**Output:**
```
3
```

---

## **4. Printing the Entire Map**
```go
fmt.Println(sozluk)
```
- **Prints the entire map** (order may vary because maps in Go are unordered).
- Example output:
  ```
  map[car:araba chair:sandalye fury:ofke]
  ```

---

## **5. Deleting a Key from the Map**
```go
delete(sozluk, "car")
fmt.Println(sozluk)
```
- **`delete(sozluk, "car")`** removes `"car"` from the map.
- Now, `"car"` no longer exists.

**Output:**
```
map[chair:sandalye fury:ofke]
```

---

## **6. Checking if a Key Exists in the Map**
```go
deger, varmi := sozluk["xxx"]
fmt.Println(deger)
fmt.Println(varmi)
```
- `"xxx"` does **not** exist in the map.
- **When a key is missing:**
  - `deger` gets the **default value** (empty string `""` for `string` type).
  - `varmi` (a boolean) becomes `false`.

**Output:**
```
(empty line)
false
```

---

## **7. Declaring and Initializing a Map in One Line**
```go
sozluk2 := map[string]string{"glass": "bardak", "window": "cam"}
fmt.Println(sozluk2)
```
- Creates `sozluk2` with predefined values:
  - `"glass" → "bardak"`
  - `"window" → "cam"`

**Output:**
```
map[glass:bardak window:cam]
```

---

## **Final Output of the Program**
```
sandalye
3
map[car:araba chair:sandalye fury:ofke]
map[chair:sandalye fury:ofke]

false
map[glass:bardak window:cam]
```

---

## **Key Takeaways**
| Concept | Explanation |
|---------|------------|
| **Creating a Map** | `sozluk := make(map[string]string)` |
| **Adding Values** | `sozluk["chair"] = "sandalye"` |
| **Accessing Values** | `fmt.Println(sozluk["chair"]) // sandalye` |
| **Getting Map Length** | `len(sozluk) // 3` |
| **Deleting a Key** | `delete(sozluk, "car")` |
| **Checking Key Existence** | `value, exists := sozluk["xxx"]` (exists is `false`) |
| **Initializing a Map with Values** | `sozluk2 := map[string]string{"glass": "bardak", "window": "cam"}` |

This is a great example of working with **maps** (dictionaries) in Go! 🚀