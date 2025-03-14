The code you've provided demonstrates the behavior of **`iota`** in Go, which is used to generate a sequence of related constant values. Here's a step-by-step explanation of the rules behind the code:

### Code Breakdown:

```go
const (
	a = iota // 0
	b        // 1
	c        // 2
)
```

- **`iota`** is a special Go identifier that is used inside constant declarations. It is a way to automatically generate successive integer values. The first time `iota` is used in a `const` block, it is set to `0`. For every subsequent constant in the same block, `iota` increments by 1.

### Rules and Behavior:

1. **First Constant (`a`):**
   - **`a = iota`**: 
     - When `iota` is used for the first time in a `const` block, it is set to `0`. 
     - So, `a` is assigned the value `0`.

2. **Second Constant (`b`):**
   - **`b`**:
     - When `iota` is used again in the next line, it **automatically increments** to `1`. 
     - So, `b` is assigned the value `1` without needing to explicitly write `iota = 1`.

3. **Third Constant (`c`):**
   - **`c`**:
     - Similarly, `iota` is automatically incremented again to `2`, so `c` is assigned the value `2`.

### Key Points:
- **Automatic Increment:** Once `iota` is introduced in a `const` block, each subsequent constant in that block automatically gets the next incremented value of `iota`, starting from `0`.
- **No Need to Explicitly Define `iota`:** After the first occurrence of `iota`, you don't need to repeat `iota` in each line. Go automatically increments it for you.
  
### Output:
- **`a = 0`**
- **`b = 1`**
- **`c = 2`**

This is how the constants are assigned automatically using `iota`, and the printed output will be:

```
0
1
2
```

### Summary:
- `iota` is a special identifier in Go used to generate a sequence of constant values, starting at `0` and automatically incrementing by 1 for each subsequent constant.
- By using `iota` once at the beginning of a `const` block, you can let Go automatically assign increasing values to all the constants in that block without explicitly writing the value for each one.