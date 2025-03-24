### Explanation of the Code:

This Go program calculates the volume of a sphere based on the radius provided as a command-line argument and prints the result.

### Code Breakdown:

```go
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var radius, vol float64

	// Parse the radius from the command-line argument
	radius, _ = strconv.ParseFloat(os.Args[1], 64)

	// Calculate the volume of the sphere
	vol = (4 * math.Pi * math.Pow(radius, 3)) / 3

	// Print the radius and calculated volume with formatted output
	fmt.Printf("radius: %g -> volume: %.2f\n", radius, vol)
}
```

### Key Concepts:

1. **Command-line Arguments (`os.Args`)**:
   - `os.Args` is a slice of strings that contains the arguments passed when running the program.
   - `os.Args[0]`: This is the program name or the path used to execute it.
   - `os.Args[1]`: This is the first argument passed to the program (in this case, the radius of the sphere).
   
   For example, when running:
   ```bash
   go run main.go 6
   ```
   - `os.Args[1]` will be `"6"`, representing the radius.

2. **Converting String to Float (`strconv.ParseFloat`)**:
   - `strconv.ParseFloat(os.Args[1], 64)` is used to convert the string value of `os.Args[1]` (the radius) to a `float64` value.
   - The `64` indicates that the float should have a precision of 64 bits.
   - The function returns two values: the parsed `float64` value (`radius`) and an error (`_` is used to ignore the error in this case).
   - If the argument can't be converted to a float, the error is ignored, and the program proceeds.

3. **Volume Formula**:
   - The formula to calculate the volume of a sphere is:
     \[
     V = \frac{4}{3} \pi r^3
     \]
     Where:
     - \( V \) is the volume.
     - \( r \) is the radius of the sphere.
     - \( \pi \) is the mathematical constant pi (available from the `math.Pi` constant in Go).
   
   - In the code:
     ```go
     vol = (4 * math.Pi * math.Pow(radius, 3)) / 3
     ```
     - `math.Pi` provides the value of pi.
     - `math.Pow(radius, 3)` calculates \( r^3 \) (radius cubed).
     - Then, the formula `(4 * math.Pi * math.Pow(radius, 3)) / 3` calculates the volume of the sphere.

4. **Formatted Output (`fmt.Printf`)**:
   - `fmt.Printf()` is used to print the result with formatted output.
   - The format string is:
     ```go
     fmt.Printf("radius: %g -> volume: %.2f\n", radius, vol)
     ```
     - `%g` is used for `radius` and it prints the value of the radius in the most compact form (either as an integer or floating-point number).
     - `%.2f` formats `vol` to print the floating-point number with exactly 2 decimal places.
   - Example output: `radius: 6 -> volume: 904.78`.

### Example:

#### Input:
When running the program with the following command:

```bash
go run main.go 6
```

- `os.Args[1]` is `"6"`, which is the radius of the sphere.
- The program will calculate the volume using the formula \( \frac{4}{3} \pi r^3 \), with \( r = 6 \).

#### Calculation:
Using the formula:
\[
V = \frac{4}{3} \pi (6)^3 = \frac{4}{3} \pi 216 = 904.7787...
\]
The volume is approximately `904.78` (rounded to two decimal places).

#### Output:
```text
radius: 6 -> volume: 904.78
```

### Behavior:

- **Command-line Arguments**:
  - The program expects a radius as a command-line argument when it is run.
  - If the radius is not provided or cannot be converted to a float, the program will panic or produce incorrect results. There is no explicit error handling for invalid inputs in the current code.

- **Volume Calculation**:
  - The formula is correctly implemented for calculating the volume of a sphere.

### Error Handling:

- If the user does not provide a valid number (for example, a non-numeric string or an empty argument), the program will fail because the `strconv.ParseFloat()` function will return an error. However, this error is not explicitly handled in the code, which could cause the program to panic if the conversion fails.
  
  For better error handling, you could add an explicit check like this:

  ```go
  radius, err := strconv.ParseFloat(os.Args[1], 64)
  if err != nil {
      fmt.Println("Error: invalid radius")
      return
  }
  ```

### Summary:

- The program calculates the volume of a sphere based on the radius passed as a command-line argument.
- It uses `strconv.ParseFloat` to convert the argument into a `float64` and then calculates the volume using the formula for the volume of a sphere.
- The result is printed in a formatted way, showing the radius and the calculated volume rounded to two decimal places.