Let's break down the code you've provided and explain the key rules and concepts involved:

### **1. Declaring and Initializing an Array**
```go
var arr [4]int64
```
- Here, `arr` is declared as an **array** of size 4 (`[4]int64`), meaning it can store 4 elements of type `int64`.
- Arrays in Go have a fixed size, and in this case, the size is 4. The array is initially filled with **zeroes** (`0`) because the default value for `int64` is `0`.

### **2. Reading User Input Using `bufio.Reader`**
```go
reader := bufio.NewReader(os.Stdin)
```
- A `bufio.Reader` is created using `os.Stdin` to read input from the standard input (keyboard).
- `bufio.Reader` allows reading input line-by-line and handling various input types.

### **3. Reading and Converting Strings to Integers**
```go
str, _ := reader.ReadString('\n')
str = strings.TrimSpace(str)
num, _ := strconv.ParseInt(str, 10, 32)
```
- `reader.ReadString('\n')` reads the input from the user until the newline character (`'\n'`), i.e., when the user presses "Enter".
- `strings.TrimSpace(str)` removes any leading or trailing whitespace characters from the input string.
- `strconv.ParseInt(str, 10, 32)` converts the string input (`str`) into an integer of type `int64`. It takes three arguments:
  - The string to be parsed (`str`).
  - The base to interpret the number (`10` for decimal).
  - The bit size (`32` here, but the result is cast to `int64`).

This process is repeated for three different numbers, which are then stored in the array `arr`.

### **4. Storing the User Input into the Array**
```go
arr[0] = num
arr[1] = num2
arr[2] = num3
```
- After converting the input string into an integer, the values are assigned to specific indices in the `arr` array.
- `arr[0]`, `arr[1]`, and `arr[2]` hold the first, second, and third numbers input by the user.

### **5. Printing Array Elements**
```go
fmt.Println("Array from index 0 to 2:", arr[0], arr[1], arr[2])
fmt.Println("Complete array :", arr)
```
- `arr[0]`, `arr[1]`, and `arr[2]` print the first three elements of the array.
- `arr` prints the complete array, including any uninitialized elements (in this case, the 4th element, which is `0` because Go initializes arrays with default values).

### **6. Printing the Length of the Array**
```go
fmt.Println("Length of array is :", len(arr))
```
- `len(arr)` returns the length of the array, which is `4` because `arr` is declared with a size of `4`.

### **7. Creating and Printing a New Array**
```go
var newarr = [5]string{"test1", "test2", "test3"}
fmt.Println("New array is : ", newarr)
```
- A new array `newarr` is created with a size of 5 (`[5]string`), and three elements are provided during initialization (`"test1"`, `"test2"`, and `"test3"`).
- The last two elements of `newarr` are automatically set to the default value for `string` (an empty string `""`) because arrays in Go are initialized with default values when not fully populated.
- When printed, the result is `["test1" "test2" "test3" "" ""]`.

### **Summary of Key Rules and Concepts:**

1. **Array Declaration and Initialization**:
   - Arrays are declared with a fixed size and type: `[size]type`. For example, `[4]int64` is an array of 4 `int64` elements.
   - If an array is partially initialized, the uninitialized elements will take the **default values** for their type (e.g., `0` for integers, `""` for strings).

2. **Reading User Input**:
   - `bufio.NewReader(os.Stdin)` is used to read input from the user.
   - The `ReadString('\n')` function reads the input line-by-line, and `strings.TrimSpace()` removes any unwanted whitespace from the input.
   - `strconv.ParseInt()` is used to convert the string input into an integer type.

3. **Array Assignment**:
   - Once user input is received and converted, it is stored in specific indices of the array.
   - Arrays in Go are indexed starting from 0, so `arr[0]` holds the first input value.

4. **Array Length**:
   - The `len()` function can be used to get the length of an array. This returns the number of elements in the array, not the size in memory.

5. **Default Values for Arrays**:
   - Arrays in Go are initialized with **default values** when not fully initialized. For example, an `int` array is initialized with `0`, and a `string` array is initialized with `""` (empty string).

---

### **Sample Output:**
```
Tutorial on array
Enter a number
6
Enter a number
4
Enter a number
7
Array from index 0 to 2: 6 4 7
Complete array : [6 4 7 0]
Length of array is : 4
New array is :  [test1 test2 test3  ]
```

This program demonstrates how to work with arrays, read user input, and handle basic array operations in Go.