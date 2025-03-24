### Explanation of the Go Program (Command-line Arguments and Menu System):

This Go program defines a command-line interface (CLI) application where users can issue commands like `greet` and `version`, along with flags and positional arguments. It uses the `flag` package for handling flags and subcommands in the menu system.

Let's break down the program:

### 1. **Imports**:
```go
import (
	"flag"
	"fmt"
	"os"
	"strings"
)
```
- `flag`: Used for parsing command-line flags and arguments.
- `fmt`: Used for formatted I/O operations, such as printing output.
- `os`: Provides access to system-level functionality, such as reading command-line arguments.
- `strings`: Used for string manipulation functions, such as case-insensitive comparison (`strings.ToLower`).

### 2. **Constants**:
```go
const version = "1.0.0"
const usage = `Usage:

%s [command]

Commands:
    Greet
    Version
`
```
- `version`: Defines the version of the application (in this case, `1.0.0`).
- `usage`: Provides a general usage message for the CLI, listing available commands.

### 3. **MenuConf Struct**:
```go
type MenuConf struct {
	Goodbye bool
}
```
- The `MenuConf` struct holds the configuration for the menu system. It has a boolean field `Goodbye` which will be set if the user requests a "goodbye" message instead of a "hello" when greeting.

### 4. **SetupMenu Method**:
```go
func (m *MenuConf) SetupMenu() *flag.FlagSet {
	menu := flag.NewFlagSet("menu", flag.ExitOnError)
	menu.Usage = func() {
		fmt.Printf(usage, os.Args[0])
		menu.PrintDefaults()
	}
	return menu
}
```
- **Purpose**: Initializes the base flag set for the menu.
- `menu := flag.NewFlagSet("menu", flag.ExitOnError)` creates a new set of flags (`menu`) with error handling that exits on errors.
- `menu.Usage` is customized to print the general usage information (`usage` constant) for the program.
- The method returns the `menu` flag set.

### 5. **GetSubMenu Method**:
```go
func (m *MenuConf) GetSubMenu() *flag.FlagSet {
	submenu := flag.NewFlagSet("submenu", flag.ExitOnError)
	submenu.BoolVar(&m.Goodbye, "goodbye", false, "Say goodbye instead of hello")
	submenu.Usage = func() {
		fmt.Printf(greetUsage, os.Args[0])
		submenu.PrintDefaults()
	}
	return submenu
}
```
- **Purpose**: This method sets up the flag set for a specific subcommand, in this case, the `greet` subcommand.
- `submenu := flag.NewFlagSet("submenu", flag.ExitOnError)` creates a new flag set for the `greet` command.
- `submenu.BoolVar(&m.Goodbye, "goodbye", false, "Say goodbye instead of hello")` defines a boolean flag `-goodbye`, which, if set, will change the greeting message from "Hello" to "Goodbye".
- The method returns the `submenu` flag set.

### 6. **Greet Method**:
```go
func (m *MenuConf) Greet(name string) {
	if m.Goodbye {
		fmt.Println("Goodbye " + name + "!")
	} else {
		fmt.Println("Hello " + name + "!")
	}
}
```
- **Purpose**: This method is invoked when the user runs the `greet` command. It prints a greeting based on the value of the `Goodbye` flag.
- If the `Goodbye` flag is set to `true`, it prints "Goodbye [name]!".
- Otherwise, it prints "Hello [name]!".

### 7. **Version Method**:
```go
func (m *MenuConf) Version() {
	fmt.Println("Version: " + version)
}
```
- **Purpose**: This method is invoked when the user runs the `version` command. It simply prints the version of the application (in this case, `1.0.0`).

### 8. **Main Function**:
```go
func main() {
	c := MenuConf{}
	menu := c.SetupMenu()

	if err := menu.Parse(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing params %s, error: %v", os.Args[1:], err)
		return
	}

	// we use arguments to switch between commands
	// flags are also an argument
	if len(os.Args) > 1 {
		// we don't care about case
		switch strings.ToLower(os.Args[1]) {
		case "version":
			c.Version()
		case "greet":
			f := c.GetSubMenu()
			if len(os.Args) < 3 {
				f.Usage()
				return
			}
			if len(os.Args) > 3 {
				if err := f.Parse(os.Args[3:]); err != nil {
					fmt.Fprintf(os.Stderr, "Error parsing params %s, error: %v", os.Args[3:], err)
					return
				}
			}
			c.Greet(os.Args[2])

		default:
			fmt.Println("Invalid command")
			menu.Usage()
			return
		}
	} else {
		menu.Usage()
		return
	}
}
```
- **Menu Setup**: 
  - `menu := c.SetupMenu()` initializes the main menu with the base flags.
  - `menu.Parse(os.Args[1:])` parses the arguments passed to the program (ignoring the program name).
  - If there's an error in parsing, the program prints an error message and exits.

- **Switch Based on Command**: 
  - The `switch` statement checks the first argument (`os.Args[1]`) to determine which command the user has entered (`version`, `greet`, or an invalid command).
  - The `strings.ToLower(os.Args[1])` ensures that the command is case-insensitive.

- **Handle `greet` Command**:
  - If the user runs `greet`, it checks if a name is provided (i.e., `os.Args[2]`).
  - It also handles additional flags by calling `GetSubMenu` to parse further arguments.
  - If the `-goodbye` flag is passed, the `Greet` method will print a goodbye message instead of a hello message.

- **Handle `version` Command**:
  - If the user runs `version`, it simply prints the version using the `Version()` method.

- **Usage**: If the user runs the program without any valid command or provides invalid commands, it will print the general usage and exit.

### 9. **Examples of Program Execution**:

- **Command to check version**:
  ```sh
  $ ./cmdargs version
  Version: 1.0.0
  ```

- **Command to greet without goodbye**:
  ```sh
  $ ./cmdargs greet Alice
  Hello Alice!
  ```

- **Command to greet with goodbye**:
  ```sh
  $ ./cmdargs greet Alice -goodbye
  Goodbye Alice!
  ```

- **Invalid Command**:
  ```sh
  $ ./cmdargs unknown
  Invalid command
  Usage:
  [program name] [command]
  Commands:
    Greet
    Version
  ```

### Summary:
This program defines a simple CLI application that supports subcommands (`version` and `greet`) and flags (`-goodbye`). The use of `flag` package allows for parsing of commands and flags, while the `strings` package handles case-insensitive commands. The program structure is designed to be extensible, allowing for easy addition of more commands and flags.