This Go program takes up to 5 numbers as command-line arguments, stores them in an array, and sorts the numbers in ascending order using a **bubble sort** algorithm. Let's break down how the program works:

### 1. **Command-Line Argument Parsing**:
```go
args := os.Args[1:]

switch l := len(args); {
case l == 0:
    fmt.Println("Please give me up to 5 numbers.")
    return
case l > 5:
    fmt.Println("Sorry. Go arrays are fixed.",
        "So, for now, I'm only supporting sorting 5 numbers...")
    return
}
```
- **`os.Args[1:]`**: This retrieves all the command-line arguments passed to the program, excluding the program's name (the first element at index `0`).
- **`switch l := len(args);`**: This checks the length of the arguments (`l`) and ensures the number of arguments is between 1 and 5 (inclusive). 
  - **If there are no arguments (`l == 0`)**: The program will print `"Please give me up to 5 numbers."` and exit using `return`.
  - **If there are more than 5 arguments (`l > 5`)**: It will print `"Sorry. Go arrays are fixed."`, and then inform the user that only 5 numbers are supported. After that, it will exit.

### 2. **Declaring and Populating the Array**:
```go
var nums [5]float64
```
- **`nums`** is a fixed-size array of `float64` values with a length of 5.
  
Next, the program fills the `nums` array with numbers from the command-line arguments:
```go
for i, v := range args {
    n, err := strconv.ParseFloat(v, 64)
    if err != nil {
        // skip if it's not a valid number
        continue
    }

    nums[i] = n
}
```
- **`for i, v := range args`**: This loop iterates over the command-line arguments, where `i` is the index and `v` is the argument (as a string).
- **`strconv.ParseFloat(v, 64)`**: This function tries to convert the string argument `v` to a `float64`. If the conversion fails (invalid input), an error (`err`) is returned.
- **`if err != nil`**: If the conversion failed (i.e., `err` is not `nil`), the program continues to the next argument without processing the invalid one.
- **`nums[i] = n`**: If the conversion succeeds, the number `n` is stored in the `nums` array at index `i`.

### 3. **Bubble Sort to Sort the Array**:
```go
for range nums {
    for i, v := range nums {
        if i < len(nums)-1 && v > nums[i+1] {
            nums[i], nums[i+1] = nums[i+1], nums[i]
        }
    }
}
```
- **Outer loop (`for range nums`)**: This loop runs over the entire `nums` array. It ensures that the sorting process repeats enough times to guarantee that all the elements are sorted.
- **Inner loop (`for i, v := range nums`)**: This loop iterates through each element in the `nums` array and compares it with the next element (`nums[i+1]`).
- **Condition `if i < len(nums)-1 && v > nums[i+1]`**:
  - `i < len(nums)-1` ensures we don’t attempt to compare the last element with the element after it, as that would lead to an index out-of-bounds error.
  - `v > nums[i+1]` checks if the current element `v` is greater than the next element `nums[i+1]`. If this is true, the elements are swapped.
- **Swapping the elements**: `nums[i], nums[i+1] = nums[i+1], nums[i]` swaps the values of `nums[i]` and `nums[i+1]` to sort them in ascending order.

### 4. **Output the Sorted Array**:
```go
fmt.Println(nums)
```
- After the sorting is complete, the program prints the sorted array `nums`.

### Example Run:

#### Valid Input:
```bash
go run main.go 56 12 90 34 5
```

**Output**:
```
[5 12 34 56 90]
```

#### Invalid Input (Too Many Arguments):
```bash
go run main.go 56 12 90 34 5 100
```

**Output**:
```
Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
```

#### Invalid Input (Non-Numeric):
```bash
go run main.go 56 abc 90 34 5
```

**Output**:
```
[56 0 90 34 5]
```
- In this case, `"abc"` is ignored as it’s not a valid number. The result is `[56 0 90 34 5]` because `0` is the default value in the array when an invalid number is skipped.

### **Summary**:
- **Command-Line Arguments**: The program uses `os.Args` to take input from the user.
- **Error Handling**: The program handles invalid input gracefully by skipping any invalid numbers.
- **Bubble Sort**: The program uses the **bubble sort** algorithm to sort the numbers in ascending order. It compares adjacent numbers and swaps them if they are out of order. This process repeats until the array is fully sorted.
- **Array Limit**: The program enforces a maximum of 5 numbers, as the array is fixed to hold 5 elements.

This program demonstrates basic error handling, sorting, and working with arrays in Go.