In this Go code, you are exploring different ways to use the `for` loop, which is the only loop construct in Go (there is no `while` loop). Let's break down each part of the code and explain how it works.

### Code Breakdown:

```go
package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	// if i == 3 {
	// 	// 	break
	// 	// }
	// 	// if i == 3 {
	// 	// 	continue
	// 	// }
	// 	// fmt.Println(i)
	// }

	// while文的な書き方
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// while文的な書き方その２(あわや無限ループ)
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 3 {
			break
		}
	}
}
```

### 1. **Commented-out code (the first loop)**
```go
// for i := 0; i < 10; i++ {
// 	// if i == 3 {
// 	// 	break
// 	// }
// 	// if i == 3 {
// 	// 	continue
// 	// }
// 	// fmt.Println(i)
// }
```
This block of code is commented out, but let's explain it as if it were active:

- This is a typical `for` loop where the initialization (`i := 0`), condition (`i < 10`), and post-expression (`i++`) are defined explicitly.
- Inside the loop, there are two conditions:
  - **`break`**: If `i == 3`, the loop is terminated immediately (no further iterations will occur).
  - **`continue`**: If `i == 3`, the loop skips the current iteration and moves to the next iteration (without printing `3`).
- Finally, `fmt.Println(i)` will print the value of `i` in each iteration of the loop.

### 2. **First "while-like" loop (imitating a `while` loop)**
```go
// i := 0
// for i < 10 {
// 	fmt.Println(i)
// 	i++
// }
```
- This is another loop using the `for` construct, but it's written in a way that resembles a `while` loop (which Go doesn't have).
- **Condition-based loop**: The condition `i < 10` is evaluated before each iteration, and the loop continues running as long as `i` is less than 10.
- Inside the loop:
  - `fmt.Println(i)` prints the current value of `i`.
  - `i++` increments the value of `i` at the end of each loop iteration.
- When `i` reaches 10, the condition `i < 10` becomes `false`, and the loop exits.

### 3. **Second "while-like" loop (with an infinite loop)**
```go
i := 0
for {
	fmt.Println(i)
	i++
	if i == 3 {
		break
	}
}
```
This loop is a little different. Here’s how it works:
- **Infinite loop**: `for {}` creates an infinite loop because there is no condition, and it will continue forever unless explicitly broken out of.
- Inside the loop:
  - `fmt.Println(i)` prints the current value of `i`.
  - `i++` increments `i` after each iteration.
  - `if i == 3 { break }` checks if `i` equals `3`, and if so, the `break` statement stops the loop immediately.
- This loop prints the values `0`, `1`, and `2` before exiting the loop (because `i` reaches `3`, triggering the `break`).

### Summary of Loops in the Code:

1. **First Loop (Commented Out)**: This loop demonstrates the standard `for` loop, with a `break` and `continue` statement. The `break` would exit the loop when `i == 3`, and `continue` would skip printing `3` and go to the next iteration.
  
2. **Second Loop (Imitating `while`)**: This loop imitates a `while` loop in other languages. It uses the condition `i < 10` to keep looping, and after printing `i`, it increments `i` until `i == 10`.

3. **Third Loop (Infinite Loop)**: This loop demonstrates an infinite loop using `for {}`. It continues to run indefinitely until `i == 3`, at which point it uses the `break` statement to stop the loop. The loop prints `0`, `1`, `2`, and then stops.

### Expected Output:
For the third loop (the one that is not commented out):

```
0
1
2
```

The loop stops after printing `2` because `i` reaches `3` and the `break` statement is triggered.

### Key Takeaways:
- **For Loop in Go**: The `for` loop in Go is very versatile and can be used in several ways: as a traditional `for` loop, as a `while` loop, or even as an infinite loop with `for {}`.
- **`break` and `continue`**: These statements control the flow of the loop. `break` exits the loop, and `continue` skips the current iteration and moves to the next.
