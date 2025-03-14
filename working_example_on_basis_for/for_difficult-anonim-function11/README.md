The Go code you've provided is meant to reverse a slice of bytes (which is essentially a string, since strings are represented as slices of bytes in Go). Let's break it down step by step:

### Goal:
The function `reverseString` aims to reverse a string (represented as a slice of bytes `[]byte`) in place, and print the reversed string. The input is passed as a byte slice to the function.

### Code Breakdown:

#### 1. `st := make([]byte, len(s))`
```go
st := make([]byte, len(s))
```
- **Purpose**: This creates a new slice `st` of bytes with the same length as the input slice `s`.
- `make([]byte, len(s))` allocates a new byte slice with the same size as `s`. This slice will hold the reversed string temporarily.

#### 2. The `for` loop:
```go
for i, _ := range s {
    st[i] = s[len(s)-i-1]
}
```
- **Purpose**: This loop iterates over the original slice `s` and fills the `st` slice with the elements of `s` in reverse order.
- `for i, _ := range s`: The loop iterates over the slice `s`. `i` is the index of the slice, and `_` is used because we don't need the value of the elements in `s`.
- `st[i] = s[len(s)-i-1]`: 
  - `len(s)-i-1`: This calculates the index of the element in `s` that should be assigned to `st[i]`. For example, the last element of `s` should be assigned to `st[0]`, the second last element of `s` to `st[1]`, and so on.
  - So, this step is essentially reversing the elements of the slice.

#### 3. `s = append(s[:0], st...)`
```go
s = append(s[:0], st...)
```
- **Purpose**: This line is trying to reassign the reversed values from `st` back to `s`. However, this part is a bit redundant and doesn't actually modify the `s` slice in place because slices are reference types in Go.
  - `s[:0]`: This creates a new slice with zero length but the same capacity as `s`.
  - `st...`: This is a variadic argument, which means the elements of `st` are "spread out" and passed as individual arguments to `append`.
  - The idea is to append the reversed elements from `st` to the empty slice `s[:0]`.

However, the key point here is that `s` in the function is passed by value, meaning modifications to `s` inside the function will not persist in the caller unless you return the result. So the slice `s` will remain unchanged in the `main` function.

#### 4. `fmt.Println(string(s), st)`
```go
fmt.Println(string(s), st)
```
- **Purpose**: This line prints two values:
  1. `string(s)`: It tries to convert the (possibly modified) `s` byte slice back into a string and prints it.
  2. `st`: The reversed byte slice is printed as-is, which shows the reversed values.

However, the expected output might be confusing because the `s` slice doesn't actually get updated in the function due to Go's pass-by-value nature.

### 5. `main` function:
```go
var s = []byte{'h','e','l','l','o'}
reverseString(s)
```
- **Purpose**: The `main` function initializes a byte slice `s` containing the word "hello", and then calls the `reverseString` function to reverse it.

### Key Points:

1. **Slices in Go are reference types**: When you pass a slice to a function, the slice itself is passed by reference, but the slice variable inside the function is still a copy (i.e., the slice header is copied). Any changes to the slice (like appending elements or modifying its length) don't affect the original slice in the calling function unless you return the modified slice.
  
2. **Reversing logic**: The `for` loop fills the new slice `st` with the elements of `s` in reverse order. 

3. **Reassigning the slice**: The statement `s = append(s[:0], st...)` doesn't actually modify the original `s` slice, because `s` is a local copy of the original slice. In Go, if you want to modify the original slice, you need to either modify the slice in place (without reassigning `s`) or return the modified slice.

### Output:
- The expected output might look like this:
  ```
  hello [111 111 111 101 104]
  ```
  Here's why:
  - `string(s)`: It will still print `"hello"` because `s` was not actually modified (due to how slices work in Go).
  - `st`: This will print the reversed byte slice: `[111 111 111 101 104]`, which represents the ASCII values for `"o" "l" "l" "e" "h"`.

### To Fix the Issue:
To actually modify `s` and print the reversed string, you should return the reversed slice from the `reverseString` function. Here's the corrected version:

```go
func reverseString(s []byte) []byte {
    st := make([]byte, len(s))
    for i := range s {
        st[i] = s[len(s)-i-1]
    }
    return st
}

func main() {
    var s = []byte{'h', 'e', 'l', 'l', 'o'}
    reversed := reverseString(s)
    fmt.Println(string(reversed))
}
```

### Output:
```
olleh
```

Now, `reverseString` returns the reversed slice, and `main` correctly prints the reversed string.