In the code you provided, we see the use of **`iota`** along with the **blank identifier** (`_`). Let's break it down to explain how this works:

### Code Breakdown:

```go
const (
	_ = iota      // 0
	b = iota * 10 // 1 * 10
	c = iota * 10 // 2 * 10
)
```

### Explanation:

1. **`_ = iota` (Blank Identifier for the first constant):**
   - **`iota`** starts at 0 when it is first used in a constant block.
   - However, the **blank identifier** (`_`) is used to **discard the value of `iota`** for the first constant. 
     - The blank identifier (`_`) is a special Go identifier that allows us to **ignore** values we don't need.
   - **Result:** The first constant will not hold a value at all, and `iota`'s value (which is `0`) is discarded. It does not get assigned to any constant.

2. **`b = iota * 10` (Second Constant):**
   - Now, `iota` is incremented to 1.
   - The value of `iota` (which is `1`) is multiplied by `10`.
   - **Result:** `b = 1 * 10 = 10`.

3. **`c = iota * 10` (Third Constant):**
   - Again, `iota` is incremented to 2.
   - The value of `iota` (which is `2`) is multiplied by `10`.
   - **Result:** `c = 2 * 10 = 20`.

### Key Points:

- **Blank Identifier (`_`):** The first constant uses the blank identifier (`_`), which effectively ignores the value of `iota` for that first constant (`0`). This allows us to avoid assigning `0` to that constant.
- **Using `iota` in Expressions:** For subsequent constants, we can use `iota` as part of an expression. In this case, `iota * 10` is used to assign values to `b` and `c`. The value of `iota` increments automatically with each constant declaration.
  
### Output:
- **`b = 10`**
- **`c = 20`**

So, when the program runs, it will print:

```
10
20
```

### Summary:
- **`iota`** is used to automatically increment values for constants.
- The **blank identifier (`_`)** discards the first value of `iota`, effectively skipping the first constant.
- The **second and third constants** (`b` and `c`) use the values of `iota` multiplied by `10`, giving `b = 10` and `c = 20`.