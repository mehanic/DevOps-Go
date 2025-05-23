Let's break down the Go program step by step and explain how it works:

### Code Breakdown:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	responses := make(map[string]string) // 1. Initialize a map to store responses
	pollingActive := true                 // 2. Polling control flag

	scanner := bufio.NewScanner(os.Stdin) // 3. Create a scanner to read input from the user

	for pollingActive { // 4. Loop to conduct the poll until polling is active
		// 5. Get user's name
		fmt.Print("\nWhat is your name? ")
		scanner.Scan()
		name := scanner.Text()

		// 6. Get the user's response for the mountain they'd like to climb
		fmt.Print("Which mountain would you like to climb someday? ")
		scanner.Scan()
		response := scanner.Text()

		// 7. Store the response in the map
		responses[name] = response

		// 8. Ask if another person wants to respond
		fmt.Print("Would you like to let another person respond? (yes/no) ")
		scanner.Scan()
		repeat := scanner.Text()

		// 9. If "no", end the polling; if "yes", continue to the next iteration
		if strings.ToLower(repeat) == "no" {
			pollingActive = false
		}
	}

	// 10. Print the final poll results
	fmt.Println("\n--- Poll Results ---")
	for name, response := range responses {
		fmt.Printf("%s would like to climb %s.\n", name, response)
	}
}
```

### Explanation:

1. **Map Initialization**:
   - `responses := make(map[string]string)`: This creates an empty map named `responses` that will store user names as keys and their mountain climbing preferences as values. The map is of type `map[string]string`, meaning both the keys and values are strings.

2. **Polling Control Flag**:
   - `pollingActive := true`: This variable controls the loop that conducts the poll. It starts as `true`, meaning the polling will continue until the user opts to stop.

3. **Scanner Setup**:
   - `scanner := bufio.NewScanner(os.Stdin)`: This creates a scanner to read input from the standard input (keyboard). It will be used to read user responses during the poll.

4. **Main Loop**:
   - `for pollingActive { ... }`: This loop runs the polling process repeatedly, allowing multiple users to participate. The loop continues as long as `pollingActive` is `true`.

5. **Prompt for User's Name**:
   - `fmt.Print("\nWhat is your name? ")`: This prompts the user to enter their name.
   - `scanner.Scan()`: Reads the user input from the keyboard.
   - `name := scanner.Text()`: Stores the user's input (name) in the `name` variable.

6. **Prompt for User's Response**:
   - `fmt.Print("Which mountain would you like to climb someday? ")`: This prompts the user to input the mountain they would like to climb someday.
   - `scanner.Scan()`: Reads the response from the user.
   - `response := scanner.Text()`: Stores the response (mountain name) in the `response` variable.

7. **Storing the Response**:
   - `responses[name] = response`: This stores the user's response in the `responses` map using the user's name as the key.

8. **Prompt to Allow Another Person to Respond**:
   - `fmt.Print("Would you like to let another person respond? (yes/no) ")`: Asks the user if they want to allow another person to take the poll.
   - `scanner.Scan()`: Reads the user's input.
   - `repeat := scanner.Text()`: Stores the response ("yes" or "no") in the `repeat` variable.

9. **Check If Polling Should End**:
   - `if strings.ToLower(repeat) == "no" { pollingActive = false }`: This checks if the response to allow another person is `"no"`. If it is, the `pollingActive` flag is set to `false`, which will stop the loop, ending the polling process. The `strings.ToLower()` function ensures that the input is case-insensitive, so "NO" or "no" will both work.

10. **Print Final Results**:
   - After the loop ends (when `pollingActive` becomes `false`), the program prints the results of the poll. 
   - `for name, response := range responses { fmt.Printf("%s would like to climb %s.\n", name, response) }`: This loops through each entry in the `responses` map and prints the user's name along with the mountain they wish to climb.

### Sample Interaction:

- The program first asks for a name and a mountain to climb.
- It then asks if another person can respond. If the user types "yes," the process repeats. If they type "no," the polling stops.
- Finally, the program prints out all the responses it collected.

Example:

```
What is your name? Alice
Which mountain would you like to climb someday? Everest
Would you like to let another person respond? (yes/no) yes

What is your name? Bob
Which mountain would you like to climb someday? K2
Would you like to let another person respond? (yes/no) no

--- Poll Results ---
Alice would like to climb Everest.
Bob would like to climb K2.
```

### Key Concepts:

- **Maps**: This program uses a `map` to store the poll responses. A `map` is a key-value store, which is useful for associating a name with a mountain.
- **Loops**: The `for` loop is used for repeatedly asking questions until the user decides to stop.
- **String Manipulation**: `strings.ToLower()` is used to make the program case-insensitive when checking the user's response.

### Summary:

This program simulates a simple polling system, where users are asked to input their names and the mountains they'd like to climb. The results are stored in a map and displayed at the end. The loop allows multiple users to respond, and the polling stops when the user indicates no more responses are needed.