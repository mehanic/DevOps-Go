This Go code defines two main functions: `ReadDir` and `RunCmd`, which are used for reading environment variables from a directory and executing a command with those variables. Here's a detailed explanation of how each part works:

### 1. **`ReadDir` Function**
```go
func ReadDir(dir string) (map[string]string, error) {
	envs := make(map[string]string)
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, de := range files {
		if de.IsDir() {
			continue
		}
		if de.Name() == "" {
			continue
		}
		bites, _ := os.ReadFile(dir + "/" + de.Name())
		envs[de.Name()] = string(bites)
	}
	return envs, err
}
```

- **Purpose**: The `ReadDir` function reads all the files in a specified directory and loads their contents into a map of environment variables. The filenames become the keys in the map, and the contents of the files become the values.

- **Parameters**:
  - `dir`: The directory to be read.

- **Steps**:
  1. **Create a Map**: The function initializes a `map[string]string` to store the filenames (as keys) and their contents (as values).
  2. **Read the Directory**: It calls `os.ReadDir(dir)` to get a list of files and directories in the specified path. If an error occurs during reading, it returns the error.
  3. **Iterate Through Files**: It loops over the list of files:
     - Skips any subdirectories (using `de.IsDir()`).
     - Skips files with an empty name.
     - Reads the content of the file using `os.ReadFile` and stores the content in the `envs` map, with the filename as the key.
  4. **Return the Map**: After processing all files, it returns the `envs` map, which contains the filenames and their respective contents as key-value pairs, along with any error encountered.

- **Returns**:
  - `map[string]string`: A map of environment variables, where the keys are filenames, and the values are the contents of the files.
  - `error`: The error encountered, if any, during the directory reading process.

### 2. **`RunCmd` Function**
```go
func RunCmd(cmd []string, env map[string]string) int {
	for key, val := range env {
		err := os.Setenv(key, val)
		if err != nil {
			continue
		}
	}
	command := exec.Command(cmd[0], cmd[1:]...)
	err := command.Start()
	if err != nil {
		log.Fatal(err) // ошибка запуска
	}
	log.Printf("Waiting for command to finish...")
	err = command.Wait() // ошибка выполнения
	log.Printf("Command finished with error: %v", err)
	return 0
}
```

- **Purpose**: The `RunCmd` function sets environment variables and executes a command with those variables. It takes a command and a map of environment variables as input, sets those environment variables in the current process, and then runs the specified command.

- **Parameters**:
  - `cmd`: A slice of strings representing the command to execute. The first element (`cmd[0]`) is the command to run, and the remaining elements (`cmd[1:]`) are its arguments.
  - `env`: A map of environment variables (where the keys are the variable names, and the values are the variable values).

- **Steps**:
  1. **Set Environment Variables**: The function loops over the `env` map and sets each key-value pair as an environment variable using `os.Setenv`. If there is an error setting an environment variable, it skips that particular variable and continues to the next.
  2. **Prepare the Command**: It creates a new command using `exec.Command` by passing the command name (`cmd[0]`) and its arguments (`cmd[1:]...`).
  3. **Start the Command**: It calls `command.Start()` to start the command execution in a new process. If an error occurs during the start, it logs the error and exits the program.
  4. **Wait for Command Completion**: After starting the command, it logs the message "Waiting for command to finish..." and then calls `command.Wait()` to wait for the command to finish executing.
  5. **Log Command Completion**: After the command completes, it logs the result with the message "Command finished with error: %v", where `%v` is the error (if any) returned by `command.Wait()`.

- **Returns**:
  - `0`: The function returns 0 after the command completes, regardless of success or failure. This return value seems to be used for consistency or simplicity, as there isn't a more detailed error handling or return mechanism implemented.

### How These Functions Work Together
1. **Read Environment Variables**: First, the program calls `ReadDir` to load environment variables from a specified directory into a map. The directory contains files, and each file's name is treated as an environment variable name, with its contents as the variable's value.
   
2. **Set Environment Variables**: The `RunCmd` function is then called with a list of command arguments and the map of environment variables. It sets those environment variables in the current process using `os.Setenv`.

3. **Execute Command**: The `RunCmd` function then executes the specified command using `exec.Command` and waits for the command to finish. Any errors during the execution are logged.

### Example Scenario
Suppose you have a directory with files like this:
- `env1`: Contains `value1`
- `env2`: Contains `value2`

You would call the functions as follows:

```go
envs, err := runner.ReadDir("/path/to/envdir")
if err != nil {
    log.Fatal(err)
}

cmd := []string{"/path/to/script.sh"}  // Command to execute
runner.RunCmd(cmd, envs)               // Run command with environment variables
```

In this case:
- `ReadDir` reads all files in the `/path/to/envdir` directory and stores their contents in the `envs` map.
- `RunCmd` sets those contents as environment variables and then runs the script `/path/to/script.sh` with the environment variables set.

### Key Notes:
- The environment variables are set using `os.Setenv` before running the command, which affects the environment of the process executing the command.
- The command execution is non-blocking, but the function waits for the process to finish with `command.Wait()` and then logs the result.
