The Go program you've provided demonstrates the use of **nested for loops**. Let's break it down step-by-step to understand how it works:

### Code Breakdown

```go
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}
}
```

### Explanation:

1. **Package Declaration**:
   - `package main`: This line specifies that this file belongs to the `main` package, which is used to define an executable program in Go.

2. **Importing the `fmt` package**:
   - `import "fmt"`: The `fmt` package is imported so that the program can use functions like `fmt.Println` to output to the console.

3. **Main Function**:
   - `func main() { ... }`: The `main` function is the entry point for the program. It is where the execution begins.

4. **First (Outer) `for` Loop**:
   - `for i := 0; i < 5; i++ { ... }`: This is the outer loop that runs 5 times. The loop control variable `i` is initialized to 0. The loop will run as long as `i` is less than 5, and `i++` increments `i` by 1 after each iteration. This means `i` will take values from 0 to 4.
   
5. **Second (Inner) `for` Loop**:
   - `for j := 0; j < 5; j++ { ... }`: This is the inner loop that also runs 5 times. The loop control variable `j` is initialized to 0. The loop will run as long as `j` is less than 5, and `j++` increments `j` by 1 after each iteration. Just like the outer loop, `j` will take values from 0 to 4.
   
6. **Nested Loop Behavior**:
   - The inner loop is nested inside the outer loop, meaning for each iteration of the outer loop (`i`), the inner loop (`j`) will complete all of its iterations. The inner loop will run 5 times for each value of `i`.
   
7. **Printing the Values**:
   - `fmt.Println(i, " - ", j)`: Inside the inner loop, the program prints the current values of `i` and `j`. The `fmt.Println` function outputs the value of `i`, followed by the string `" - "`, and then the value of `j`.
   - For each combination of `i` and `j`, this will print a line showing the values of `i` and `j`.

### Loop Execution:

- The outer loop (`i`) runs 5 times (from 0 to 4).
- For each value of `i`, the inner loop (`j`) runs 5 times (from 0 to 4).
  
Thus, the total number of iterations is \(5 \times 5 = 25\) (5 iterations of the outer loop, and for each outer loop, the inner loop runs 5 times).

### Example Output:

Here is what the output will look like when you run the program:

```
0  -  0
0  -  1
0  -  2
0  -  3
0  -  4
1  -  0
1  -  1
1  -  2
1  -  3
1  -  4
2  -  0
2  -  1
2  -  2
2  -  3
2  -  4
3  -  0
3  -  1
3  -  2
3  -  3
3  -  4
4  -  0
4  -  1
4  -  2
4  -  3
4  -  4
```

### Summary:
- The outer loop (`i`) iterates 5 times, from 0 to 4.
- For each iteration of `i`, the inner loop (`j`) also iterates 5 times, from 0 to 4.
- The result is 25 printed lines, showing every combination of `i` and `j` from 0-4.
- The `fmt.Println(i, " - ", j)` prints the current value of `i` and `j` in the format `"i - j"`.