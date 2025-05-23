Let’s break down the rules and concepts used in this Go code step by step:

### **1. Strings are Immutable in Go**
```go
str := "Yūgen ☯ 💀"
```
- **Strings in Go**: Strings in Go are **immutable**, meaning you cannot change their contents directly. For example, the following code is commented out because it would result in an error:
  ```go
  // str[0] = 'N'  // error: strings are immutable
  // str[1] = 'o'  // error: strings are immutable
  ```

### **2. Converting a String to a Byte Slice**
```go
bytes := []byte(str)
```
- **String to Byte Slice**: You can convert a string to a byte slice (`[]byte`), which is a mutable data structure. This allows you to modify the individual bytes of the string (although this doesn’t work well for multibyte characters such as UTF-8 characters).

### **3. Modifying the Byte Slice**
```go
// bytes[0] = 'N'   // works, modifies the byte slice
// bytes[1] = 'o'   // works, modifies the byte slice
```
- **Byte Slice Modification**: Unlike strings, byte slices are **mutable**. You can modify individual bytes in the byte slice. However, because UTF-8 characters may consist of more than one byte, you need to be careful when modifying them directly.

### **4. Converting Byte Slice Back to String**
```go
str = string(bytes)
```
- **Byte Slice to String**: After modifying the byte slice, you can convert it back to a string using `string(bytes)`. In this case, the string `"Yūgen ☯ 💀"` is preserved, and after conversion, the modified string would be printed.

### **5. String Properties**
```go
fmt.Printf("%s\n", str)
fmt.Printf("\t%d bytes\n", len(str))
fmt.Printf("\t%d runes\n", utf8.RuneCountInString(str))
```
- **`len(str)`**: The `len()` function gives the **byte length** of the string, i.e., the number of bytes it takes to represent the string in memory.
  - For `"Yūgen ☯ 💀"`, the string takes **15 bytes** because each character may take more than 1 byte (due to UTF-8 encoding).
  
- **`utf8.RuneCountInString(str)`**: The `utf8.RuneCountInString()` function returns the **number of runes** in the string. A **rune** is a Unicode code point, and in UTF-8, a character can take more than one byte to represent a single rune.
  - `"Yūgen ☯ 💀"` contains **9 runes** (each character is counted as one rune, even if it consists of multiple bytes).

### **6. Displaying the Byte Representation**
```go
fmt.Printf("% x\n", bytes)
```
- **`% x` Format**: The `% x` format specifier prints the byte slice (`bytes`) in hexadecimal format. It displays the raw byte values in a human-readable way.
  - The byte slice for `"Yūgen ☯ 💀"` looks like this in hex:
    ```
    59 c5 ab 67 65 6e 20 e2 98 af 20 f0 9f 92 80
    ```
  - These hex values represent the characters in the string in UTF-8 encoding.

### **7. Working with Substrings and Runes**
```go
fmt.Printf("1st byte   : %c\n", str[0])           // ok
fmt.Printf("2nd byte   : %c\n", str[1])           // not ok
fmt.Printf("2nd rune   : %s\n", str[1:3])         // ok
fmt.Printf("last rune  : %s\n", str[len(str)-4:]) // ok
```
- **Accessing the String by Index**:
  - `str[0]`: This accesses the **first byte** of the string, which is `'Y'`. This works fine because `'Y'` is a single-byte character.
  - `str[1]`: This accesses the **second byte** of the string, but since `'ū'` (U+016B) is a multi-byte character, it will print the byte representation for the second byte, not the full character. This is **not correct** for multi-byte characters.
  - **Slicing by Runes**: 
    - `str[1:3]`: This gets the second rune (`'ū'`), as slicing by byte index might not correspond to a single rune if the string contains multi-byte characters.
    - `str[len(str)-4:]`: This slices the string to get the last rune (`'💀'`), starting from the 4th-to-last byte to the end.

### **8. Working with Runeliteral Data Type**
```go
runes := []rune(str)
```
- **Converting to Rune Slice**: Converting the string to a slice of runes (`[]rune(str)`) makes it easier to work with individual characters, especially for multi-byte characters.
  - Each rune is stored as a 32-bit integer (`int32`), allowing you to directly access and modify individual characters without worrying about byte lengths.

### **9. Rune-Specific Information**
```go
fmt.Printf("1st rune   : %c\n", runes[0])  // prints 'Y'
fmt.Printf("2nd rune   : %c\n", runes[1])  // prints 'ū'
fmt.Printf("first five : %c\n", runes[:5]) // prints [Y ū g e n]
```
- **Accessing Runes**:
  - `runes[0]` prints the first rune (`'Y'`).
  - `runes[1]` prints the second rune (`'ū'`).
  - `runes[:5]` prints the first five runes in the string, which are `[Y ū g e n]`.

### **10. Handling a Word Containing Multi-byte Characters**
```go
word := "öykü"
fmt.Printf("%q in runes: %c\n", word, []rune(word))
fmt.Printf("%q in bytes: % [1]x\n", word)
```
- **Multi-byte Characters**:
  - `"öykü"` is a string containing multi-byte characters. These characters may take more than one byte in UTF-8 encoding.
  - When converted to a rune slice, each character is treated as a single entity:
    - `"ö"` is a single rune (U+00F6), `"y"` is a single rune (U+0079), etc.
  - The byte representation of `"öykü"` in UTF-8 is:
    ```
    c3 b6 79 6b c3 bc
    ```
  - The `%q` format specifier prints the string with quotes around it, and `%[1]x` prints the byte representation in hex.

### **11. Slicing and Accessing Characters**
```go
fmt.Printf("%s %s\n", word[:2], []byte{word[0], word[1]})
fmt.Printf("%c\n", word[2]) // y
fmt.Printf("%c\n", word[3]) // k
fmt.Printf("%s %s\n", word[4:], []byte{word[4], word[5]}) // ü
```
- **Slicing and Accessing Bytes**:
  - `word[:2]` gives the first two characters (which are `"ö"` and `"y"`).
  - `[]byte{word[0], word[1]}` extracts the bytes corresponding to `"ö"`.
  - For individual characters like `word[2]` and `word[3]`, we can access and print the characters `y` and `k`.
  - Similarly, `word[4:]` slices the string to print the last character `"ü"`.

### **Summary of Concepts and Rules**
1. **Strings are immutable**, but byte slices (`[]byte`) are mutable.
2. **UTF-8 Encoding**: A single character may be represented by more than one byte, making it essential to handle strings properly when dealing with non-ASCII characters.
3. **Rune Representation**: A rune is a Unicode code point (32-bit), and it may be more efficient to work with runes for multi-byte characters.
4. **Indexing** a string by bytes may not work correctly for multi-byte characters, so you should convert to a rune slice to handle the characters properly.
5. **Slicing and Byte Conversion**: Be mindful of how string slicing and byte conversions work, especially when dealing with multi-byte characters like `ö` or `💀`.

