This Go program demonstrates how to handle strings and characters (runes) in UTF-8 encoding, showing several ways to inspect, manipulate, and display string data. Let's go step-by-step through the rules and operations being applied:

### 1. **String Declaration and Iteration**
```go
words := []string{
	"cool",
	"güzel",
	"jīntiān",
	"今天",
	"read 🤓",
}
```
- This creates a slice `words` that holds multiple strings. These strings contain a mix of ASCII and Unicode characters (UTF-8 encoded). For example:
  - `"cool"` is a simple ASCII string.
  - `"güzel"` contains a special character (`ü`).
  - `"jīntiān"` contains accents over characters (`ī`, `ā`).
  - `"今天"` contains Chinese characters.
  - `"read 🤓"` contains an emoji (`🤓`).

### 2. **Printing String Information**
Inside the loop, the program processes each string and prints various details:

```go
for _, s := range words {
    fmt.Printf("%q\n", s)
```
- **`%q`** prints the string in double quotes, including any special characters or spaces. It shows how each string appears.

### 3. **String Length: Bytes vs. Runes**
```go
fmt.Printf("\thas %d bytes and %d runes\n", len(s), utf8.RuneCountInString(s))
```
- **`len(s)`** returns the **number of bytes** used by the string. Each character in the string takes up a different number of bytes depending on its encoding. For example:
  - ASCII characters (like `'c'`, `'o'`, etc.) take **1 byte** each.
  - Unicode characters (like `ü` or Chinese characters) take **more than 1 byte**.
  
- **`utf8.RuneCountInString(s)`** counts the **number of runes** in the string. A **rune** represents a single Unicode character (it might take more than 1 byte).

### 4. **Hexadecimal Representation of Bytes**
```go
fmt.Printf("\tbytes   : % x\n", s)
```
- **`% x`** formats the string as a series of hexadecimal byte values. This shows how the string is encoded in memory. For example:
  - `"cool"` as bytes would appear as `63 6f 6f 6c` in hex (ASCII characters).
  - `"güzel"` would appear as `67 c3 bc 7a 65 6c` in hex (with `ü` being multi-byte).

### 5. **Hexadecimal Representation of Runes**
```go
fmt.Print("\trunes   :")
for _, r := range s {
    fmt.Printf("% x", r)
}
```
- This loop iterates over each **rune** in the string and prints its hexadecimal value. A rune is a Unicode code point, which represents a character. This allows us to see each character as a code point.

### 6. **Rune Literals (Character Representation)**
```go
fmt.Print("\trunes   :")
for _, r := range s {
    fmt.Printf("%q", r)
}
```
- **`%q`** prints the rune as a **quoted character** (similar to how it would appear in a Go string literal). This shows the character, not its underlying byte or Unicode value.
  - For example, `güzel` would print as `'g'`, `'ü'`, `'z'`, `'e'`, `'l'`.

### 7. **First Rune and Its Byte Size**
```go
r, size := utf8.DecodeRuneInString(s)
fmt.Printf("\tfirst   : %q (%d bytes)\n", r, size)
```
- **`utf8.DecodeRuneInString(s)`** extracts the first rune and its byte size from the string `s`. It returns:
  - `r`: the first rune (Unicode character),
  - `size`: the number of bytes that this rune occupies in the string.
  
For example, in the string `"güzel"`, the first rune is `'g'`, which takes 1 byte.

### 8. **Last Rune and Its Byte Size**
```go
r, size = utf8.DecodeLastRuneInString(s)
fmt.Printf("\tlast    : %q (%d bytes)\n", r, size)
```
- **`utf8.DecodeLastRuneInString(s)`** works similarly to `DecodeRuneInString`, but it extracts the **last rune** and its byte size.

For example, in `"güzel"`, the last rune is `'l'`, which takes 1 byte.

### 9. **First Two Runes (Slicing)**
```go
_, first := utf8.DecodeRuneInString(s)
_, second := utf8.DecodeRuneInString(s[first:])
fmt.Printf("\tfirst 2 : %q\n", s[:first+second])
```
- This uses **slicing** and `utf8.DecodeRuneInString` to extract the first two runes from the string `s`.
- **First rune**: We decode the first rune and get its byte size.
- **Second rune**: We slice the string starting at the byte position of the first rune and decode the second rune.
  
The program then prints the first two characters as a string.

### 10. **Last Two Runes (Slicing)**
```go
_, last1 := utf8.DecodeLastRuneInString(s)
_, last2 := utf8.DecodeLastRuneInString(s[:len(s)-last1])
fmt.Printf("\tlast 2  : %q\n", s[len(s)-last2-last1:])
```
- This similarly extracts the last two runes from the string using slicing and `utf8.DecodeLastRuneInString`.

### 11. **Converting String to Rune Slice**
```go
rs := []rune(s)
fmt.Printf("\tfirst 2 : %q\n", string(rs[:2]))
fmt.Printf("\tlast 2  : %q\n", string(rs[len(rs)-2:]))
```
- **`[]rune(s)`** converts the string `s` into a slice of runes (characters).
- The program then extracts and prints the first two and last two runes from the rune slice.

### Summary of String Operations:
1. **`len(s)`**: Gives the byte size of the string.
2. **`utf8.RuneCountInString(s)`**: Counts the number of characters (runes) in the string.
3. **`utf8.DecodeRuneInString(s)`**: Decodes the first rune and its byte size.
4. **`utf8.DecodeLastRuneInString(s)`**: Decodes the last rune and its byte size.
5. **Slicing and converting to rune slices**: Extracts parts of the string or converts it into a slice of runes.

### Output Breakdown:
For each string:
- The program prints the string and shows its byte size and rune count.
- It prints the string's bytes in hexadecimal.
- It prints each rune's value in hexadecimal and as quoted characters.
- It extracts the first and last runes, along with their byte sizes.
- It slices and prints the first two and last two runes.
- It converts the string to a rune slice and prints the first and last two runes.

### Example:
For the string `"güzel"`:
- It has 6 bytes and 5 runes (`g`, `ü`, `z`, `e`, `l`).
- The byte representation is `67 c3 bc 7a 65 6c`.
- The runes are `'g'`, `'ü'`, `'z'`, `'e'`, `'l'`.
- The first rune is `'g'` (1 byte), and the last rune is `'l'` (1 byte).
- The first two runes are `"gü"`, and the last two runes are `"el"`.

This program helps to understand the difference between bytes and runes (characters) in Go, especially when working with non-ASCII characters.