This Go program simulates a random mood generator. It takes a name as a command-line argument and then randomly assigns a mood to that name from a predefined list. Let's break down the code and explain the rules behind each part:

### 1. **Command-Line Arguments Parsing:**
```go
args := os.Args[1:]
if len(args) != 1 {
    fmt.Println("[your name]")
    return
}
```
- **`os.Args[1:]`**: This retrieves the command-line arguments passed to the program, excluding the program's name (the first element at index `0`). It stores the arguments in the `args` slice.
- **`if len(args) != 1`**: The program checks if exactly one argument is passed (the user's name). If no argument is passed or if more than one argument is provided, it prints a usage message `[your name]` and exits (`return`).
  
  This ensures that only one name is passed when running the program.

### 2. **Defining the Moods Array:**
```go
moods := [...]string{
    "happy 😀", "good 👍", "awesome 😎",
    "sad 😞", "bad 👎", "terrible 😩",
}
```
- **`moods := [...]string{...}`**: This creates an array of strings (`moods`), each representing a different mood with an associated emoji.
  - The ellipsis (`...`) means that Go will automatically determine the length of the array based on the number of elements provided. The array `moods` has 6 different moods.

### 3. **Randomly Selecting a Mood:**
```go
n := rand.Intn(len(moods))
```
- **`rand.Intn(len(moods))`**: This generates a random integer between `0` and `len(moods)-1`. The `rand.Intn(n)` function returns a random integer from the range `[0, n)`.
  - Here, `len(moods)` returns `6`, so `rand.Intn(6)` will generate a random number from `0` to `5`, corresponding to the index of a mood in the `moods` array.
  
  For example:
  - If `n = 2`, the mood chosen will be `"awesome 😎"`.
  - If `n = 4`, the mood chosen will be `"bad 👎"`.

### 4. **Displaying the Name and Mood:**
```go
fmt.Printf("%s feels %s\n", name, moods[n])
```
- **`fmt.Printf("%s feels %s\n", name, moods[n])`**: This prints the output in the format `"<name> feels <mood>"`.
  - `%s` is used to format the strings:
    - The first `%s` will be replaced by the `name` (the command-line argument passed).
    - The second `%s` will be replaced by the randomly selected mood (using the index `n` to access `moods[n]`).

### Example Runs:

1. **Valid Input (with the name "max")**:
   ```bash
   go run main.go max
   ```
   **Output**:
   ```
   max feels awesome 😎
   ```
   - In this case, `rand.Intn(6)` generates a random number, say `2`, and the output would be `"max feels awesome 😎"`.

2. **Invalid Input (No Arguments)**:
   ```bash
   go run main.go
   ```
   **Output**:
   ```
   [your name]
   ```
   - Since no name is provided, the program prints the message `[your name]` and exits.

3. **Invalid Input (More than One Argument)**:
   ```bash
   go run main.go max john
   ```
   **Output**:
   ```
   [your name]
   ```
   - Since more than one argument is passed, the program prints the message `[your name]` and exits.

### **Summary**:
- **Command-Line Argument Handling**: The program expects one command-line argument (the user's name). If no argument or more than one is passed, it prints a message and stops.
- **Random Mood Selection**: The program uses `rand.Intn` to randomly select a mood from the `moods` array.
- **Output**: It prints the name followed by the randomly selected mood in a friendly format.

This program demonstrates basic command-line argument handling, array usage, and random number generation in Go.