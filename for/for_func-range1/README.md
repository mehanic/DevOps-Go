### 🧑‍🏫 **Explanation of the Go Code:**

This Go code demonstrates how to work with **runes** (Unicode characters) and how they behave in a string in Go, focusing on:
1. **Printing and examining runes**.
2. **UTF-8 encoding** and **rune manipulation** in strings.
3. The code provides insights into the handling of multi-byte characters (like Thai characters) in Go.

---

### 🛠 **Key Concepts:**

1. **Runes**:
    - A **rune** is an alias for `int32` in Go and represents a single Unicode character. 
    - In a string, characters can take up **multiple bytes**, especially for non-ASCII characters like Thai characters or emojis.
   
2. **UTF-8**:
    - Go strings are encoded in **UTF-8** by default, where each rune can occupy multiple bytes.

---

### 🧐 **Step-by-step Code Explanation**:

```go
package main

import (
    "fmt"
    "unicode/utf8"
)
```

- **Importing packages**:
  - `fmt` for formatted I/O (e.g., printing).
  - `unicode/utf8` provides functions to work with UTF-8 encoded strings (decoding and determining rune lengths).

---

#### 1. **`examineRune` function**:
```go
func examineRune(r rune) {
    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส'  {
        fmt.Println("found so sua")
    }
}
```

- **`examineRune`** accepts a **rune** (`r`) and checks whether it matches either the ASCII character `'t'` or the Thai character `'ส'` (the first character in the string `สวัสดี`).
- If a match is found, it prints the corresponding message.

---

#### 2. **`main` function**:
```go
const s = "สวัสดี"
```

- We define a constant string `s` which contains the Thai greeting **"สวัสดี"** (Hello).

---

#### 3. **Printing string length**:
```go
fmt.Println("Len:", len(s))
```

- **`len(s)`** gives the number of **bytes** in the string, not the number of runes.
- Since "สวัสดี" contains 7 bytes in UTF-8 encoding (each Thai character takes 3 bytes), `len(s)` would return **21**.

---

#### 4. **Printing byte values**:
```go
for i := 0; i < len(s); i++ {
    fmt.Printf("%x ", s[i])
}
fmt.Println()
```

- The **`for`** loop prints the **byte values** (hexadecimal format) of each byte in the string `s`.
- This will display the individual byte representation of each rune in the UTF-8 encoded string.

---

#### 5. **Rune count**:
```go
fmt.Println("Rune count:", utf8.RuneCountInString(s))
```

- **`utf8.RuneCountInString(s)`** counts the number of **runes** (Unicode characters) in the string `s`.
- For the string "สวัสดี", which contains 4 characters (each of which is a rune), the result would be **4**.

---

#### 6. **Iterating over the string using `range`**:
```go
for index, runeValue := range s {
    fmt.Printf("%#U starts at %d\n", runeValue, index)
}
fmt.Println()
```

- The **`range`** loop iterates over the string `s` and returns the **index** and **runeValue** for each rune.
- `runeValue` is the actual Unicode character, and **`%#U`** prints the **Unicode format** (e.g., `U+0E2A` for "ส").
- The index represents the **byte index**, where each rune starts in the string.

---

#### 7. **Using `utf8.DecodeRuneInString`**:
```go
for i, w := 0, 0; i < len(s); i += w {
    runeValue, width := utf8.DecodeRuneInString(s[i:])
    fmt.Printf("%#U starts at %d\n", runeValue, i)
    w = width

    examineRune(runeValue)
}
```

- This loop uses **`utf8.DecodeRuneInString`**, which decodes a **single rune** from a UTF-8 encoded string starting at a given byte position (`s[i:]`).
  - **`runeValue`** is the decoded rune.
  - **`width`** is the number of bytes that the rune occupies (e.g., Thai characters take 3 bytes each).
- The loop updates `i` by `w` (the width of the current rune), and continues decoding the string until all runes have been processed.
- It also calls **`examineRune(runeValue)`** to check if the rune matches 't' or 'ส', printing a message accordingly.

---

### 📈 **Expected Output**:
Assuming the string `s = "สวัสดี"`:

```bash
Len: 21
e0 a0 b8 e0 b8 a7 e0 b8 b5 e0 b8 a3 
Rune count: 4
U+0E2A starts at 0
U+0E27 starts at 3
U+0E2A starts at 6
U+0E35 starts at 9

Using DecodeRuneIsString
U+0E2A starts at 0
found so sua
U+0E27 starts at 3
U+0E2A starts at 6
found so sua
U+0E35 starts at 9
```

- **Byte values** are printed for each byte in the string.
- The **rune count** will show 4 runes.
- Each rune will be printed with its **Unicode representation** and **byte position** in the string.
- The loop using **`utf8.DecodeRuneInString`** decodes the string character by character, printing the Unicode values and calling `examineRune` to check for specific characters.

---

### 📝 **Conclusion:**
- **UTF-8 encoding** is important in Go for handling multi-byte characters like Thai characters.
- **Runes** are used to represent individual Unicode characters in Go, and they can be iterated and examined in strings.
- **`utf8.DecodeRuneInString`** helps to work with multi-byte characters efficiently.
