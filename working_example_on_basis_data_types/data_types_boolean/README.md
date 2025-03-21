This Go program demonstrates the use of boolean values, type conversion, and helper functions to convert integers and strings to booleans.

### Code Explanation:

1. **Defining boolean variables:**
   ```go
   a := true
   b := false
   ```
   Here, two boolean variables `a` and `b` are declared and initialized with `true` and `false`, respectively. These are simple boolean values.

2. **Printing boolean values:**
   ```go
   fmt.Println(a)
   fmt.Println(b)
   ```
   The values of `a` and `b` are printed to the console. The output will be:
   ```
   true
   false
   ```

3. **Boolean conversion from integers:**
   ```go
   fmt.Println(boolFromInt(0))
   fmt.Println(boolFromInt(1))
   fmt.Println(boolFromInt(2))
   ```
   - The function `boolFromInt()` is used to convert integers to booleans.
   - In Go, the conversion rule is: `0` is converted to `false`, and any non-zero number (like `1` and `2`) is converted to `true`.
   - `boolFromInt(0)` will return `false`, `boolFromInt(1)` will return `true`, and `boolFromInt(2)` will return `true`.

4. **Boolean conversion from strings:**
   ```go
   var c string
   fmt.Println(boolFromString(c))
   c = "Some Value"
   fmt.Println(boolFromString(c))
   ```
   - The function `boolFromString()` converts strings to booleans.
   - An empty string (`""`) is considered `false`, and any non-empty string is considered `true`.
   - Initially, the string `c` is empty, so `boolFromString(c)` will return `false`.
   - After assigning a non-empty string (`"Some Value"`) to `c`, `boolFromString(c)` will return `true`.

5. **Helper functions for conversion:**
   - `boolFromInt(n int) bool`: This function takes an integer `n` and returns `false` if `n` is `0`, and `true` for any non-zero integer.
   - `boolFromString(s string) bool`: This function takes a string `s` and returns `false` if the string is empty, and `true` for any non-empty string.

### Output:
The expected output from the program will be:
```
true
false
false
true
true
false
true
```

### Rules:
- **Integer to Boolean Conversion:** In Go, `0` is treated as `false` and any non-zero integer is treated as `true`.
- **String to Boolean Conversion:** An empty string (`""`) is treated as `false`, while any non-empty string is treated as `true`.
- **Helper Functions:** The program defines helper functions (`boolFromInt` and `boolFromString`) to handle the type conversions manually, as Go does not allow implicit conversions between these types.