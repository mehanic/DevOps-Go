This Go program demonstrates **error wrapping**, **error inspection**, and the use of **custom error types** to add context and handle errors effectively. Here's a step-by-step explanation of how it works:

### Key Concepts
1. **Custom Error Types**: A custom error type `FileLoadError` is used to capture details about the URL and the underlying error.
2. **Error Wrapping**: The `fmt.Errorf` function is used to wrap errors, preserving the original error while adding more context.
3. **Error Inspection**: The `errors.As()` function is used to check if an error is of a specific type (`FileLoadError` in this case), allowing you to access the properties of the wrapped error.

### Code Breakdown

#### 1. **Custom Error Type: `FileLoadError`**

```go
type FileLoadError struct {
	URL string
	Err error // Stores the "parent" error.
}
```
- `FileLoadError` is a custom error type that contains:
  - `URL`: A string representing the URL where the file load failed.
  - `Err`: The underlying error (e.g., `context.Canceled`), representing the reason for the failure.

#### 2. **Error Method for `FileLoadError`**

```go
func (p *FileLoadError) Error() string {
	// The error message includes the URL and the "parent" error's message.
	return fmt.Sprintf("%q: %v", p.URL, p.Err)
}
```
- The `Error()` method implements the `error` interface for `FileLoadError`.
- It formats the error message as a string, including the URL and the underlying error's message (the parent error).
  - Example output might be `"www.golang-courses.ru": context canceled`.

#### 3. **`getFile` Function**

```go
func getFile(u string) (File, error) {
	return File{}, context.Canceled
}
```
- The `getFile` function simulates loading a file from a URL. It always returns `context.Canceled`, which is a predefined error indicating the context was canceled.
- The function returns a `File` and an error. In this case, it always returns the `context.Canceled` error.

#### 4. **`loadFiles` Function**

```go
func loadFiles(urls ...string) ([]File, error) {
	files := make([]File, len(urls))
	for i, url := range urls {
		f, err := getFile(url)
		if err != nil {
			return nil, &FileLoadError{url, err} // Wraps the error in FileLoadError.
		}
		files[i] = f
	}
	return files, nil
}
```
- This function attempts to load multiple files from a list of URLs.
- For each URL, it calls `getFile`. If an error occurs (in this case, `context.Canceled`), the error is **wrapped** in a `FileLoadError` and returned.
- If all files are successfully loaded, it returns the list of files. Otherwise, it returns the wrapped error.

#### 5. **`transfer` Function**

```go
func transfer() error {
	_, err := loadFiles("www.golang-courses.ru")
	if err != nil {
		return fmt.Errorf("cannot load files: %w", err)
	}
	return nil
}
```
- The `transfer` function calls `loadFiles`. If `loadFiles` returns an error, that error is wrapped using `fmt.Errorf` with additional context ("cannot load files").
- The `%w` verb is used to wrap the original error, allowing it to be unwrapped later.
- If no error occurs, the function proceeds normally.

#### 6. **`handle` Function**

```go
func handle() error {
	if err := transfer(); err != nil {
		return fmt.Errorf("cannot transfer files: %w", err)
	}
	return nil
}
```
- The `handle` function calls `transfer`. If `transfer` returns an error, it wraps it with additional context ("cannot transfer files").
- Like `transfer`, the `%w` verb is used to propagate the wrapped error.

#### 7. **`main` Function**

```go
func main() {
	var fileLoadErr *FileLoadError
	if err := handle(); errors.As(err, &fileLoadErr) {
		fmt.Println(fileLoadErr.URL) // www.golang-courses.ru
	}
}
```
- The `main` function calls `handle()`.
- If `handle` returns an error, the `errors.As()` function is used to check if the error is of type `*FileLoadError`. If it is, the `fileLoadErr` variable will hold the `FileLoadError`.
- The `errors.As()` function allows you to **extract** the underlying `FileLoadError` and access its fields (`URL` and `Err`).
  - In this case, `fileLoadErr.URL` will be printed, which is `"www.golang-courses.ru"`.
- The `errors.As()` function checks if the error is of the specified type (`*FileLoadError`). If it is, you can access the URL of the file that caused the error.

### Explanation of Error Wrapping and `errors.As()`

1. **Error Wrapping**: 
   - The `%w` verb in `fmt.Errorf` allows you to wrap an error in another error, preserving the original error while adding context.
   - In the `transfer` and `handle` functions, the error is wrapped with additional context (e.g., "cannot load files", "cannot transfer files").
   - This allows for a **stack of error messages** that give more context about where the error occurred and why.

2. **`errors.As()`**:
   - The `errors.As()` function is used to **extract** the wrapped error if it is of a certain type. 
   - In this case, `errors.As(err, &fileLoadErr)` checks if the error is of type `*FileLoadError`. If it is, the `fileLoadErr` variable is populated with the wrapped error, allowing you to access its fields (`URL` and `Err`).
   - This is particularly useful for working with custom error types where you want to access specific details of the error.

### Output

The program will print:
```
www.golang-courses.ru
```
This is the URL of the file that failed to load due to the `context.Canceled` error.

### Conclusion

This program demonstrates how to:
- Define and use custom error types.
- Wrap errors to add context at each level of the call stack.
- Use `errors.As()` to inspect wrapped errors and access specific details (like the URL of the file that caused the error).
This pattern is very useful for maintaining clear, context-rich error messages that help with debugging and troubleshooting in real-world applications.