Let's break down the Go code you've shared and explain how it works:

### Code Breakdown:

```go
package main

import (
	"fmt"
)

func main() {
	var usrStr string
	fmt.Print("Enter your string: ")
	fmt.Scanln(&usrStr)

	index := 0
	for _, eachChar := range usrStr {
		fmt.Printf("%c --> %d\n", eachChar, index)
		index = index + 1
	}
}
```

### Explanation:

#### 1. **Package and Imports**:
   - `package main`: This indicates that the program is part of the `main` package, which is necessary for executable Go programs. The `main` package contains the `main()` function, which is the entry point for the program.
   - `import "fmt"`: The `fmt` package is imported here, which provides input and output functionalities (like printing to the console or reading from the console).

#### 2. **Variable Declaration**:
   - `var usrStr string`: This declares a variable `usrStr` of type `string`. It will hold the string input from the user.

#### 3. **Taking User Input**:
   - `fmt.Print("Enter your string: ")`: This prompts the user to enter a string, displaying the message `"Enter your string: "` in the console.
   - `fmt.Scanln(&usrStr)`: This reads the user's input and stores it in the `usrStr` variable. The `&usrStr` syntax is a way to pass the address of `usrStr` so that `Scanln` can modify the variable with the entered string.

#### 4. **Iterating Over the String**:
   - `index := 0`: This initializes a variable `index` to 0. It will be used to keep track of the position of each character in the string.

   - `for _, eachChar := range usrStr { ... }`: 
     - This is a `for` loop that iterates over each character in the string `usrStr`. 
     - `range` is a Go keyword that allows you to loop over elements in an array, slice, string, map, or channel. In this case, it loops through the string `usrStr`.
     - `eachChar` is the character from the string, and the `_` (underscore) is used because we don't need the index of the character, but rather only the value of the character itself.
   
#### 5. **Printing the Character and Index**:
   - `fmt.Printf("%c --> %d\n", eachChar, index)`: This prints each character (`eachChar`) in the string, followed by its corresponding index (`index`).
     - `%c`: This is a format specifier used to print the character.
     - `%d`: This is used to print an integer (the index).
     - `\n`: This is a newline character that ensures each output is printed on a new line.

   - `index = index + 1`: After printing each character and its index, the index is incremented by 1 to track the position of the next character.

#### 6. **Output Example**:
Let’s say the user enters the string `"hello"`. The program will print the following output:

```
Enter your string: hello
h --> 0
e --> 1
l --> 2
l --> 3
o --> 4
```

### Summary of Steps:
1. Prompt the user to input a string.
2. Loop through the string one character at a time.
3. Print each character and its index in the string.
4. The index is incremented as the loop goes through each character.

### Key Points:
- **For-loop with `range`**: It’s used to iterate through each character of the string. The first value returned by `range` is the index, but we are ignoring it by using `_` because we don't need it explicitly.
- **Character Indexing**: The loop prints both the character itself and its position in the string.
