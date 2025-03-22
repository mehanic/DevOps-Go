This Go code defines a custom error handling structure and a set of utility functions to create specific types of application errors, each associated with an HTTP status code. Let's break it down to understand the rules and behavior.

### Code Breakdown

```go
package custom_errors

import "net/http"
```
- The code is inside the `custom_errors` package, which suggests this is a package used to define custom errors for an application.
- The `net/http` package is imported to make use of HTTP status codes.

### `AppErrors` Struct
```go
type AppErrors struct {
    Code    int `json:",omitempty"`
    Message string
}
```
- **`AppErrors` struct**: This defines a custom error type, `AppErrors`, that has two fields:
  - `Code`: This is an integer field, typically used to store HTTP status codes (e.g., `404`, `500`). The `json:",omitempty"` tag ensures that if the `Code` is `0` (the zero value), it will be omitted from the JSON encoding. This can help avoid sending an unnecessary `Code` field if it's not needed.
  - `Message`: This is a string field that holds the error message associated with the error.

### `AsMessage` Method
```go
func (e AppErrors) AsMessage() *AppErrors {
    return &AppErrors{Message: e.Message}
}
```
- **`AsMessage` method**: This method takes an `AppErrors` object (`e`) and returns a new `AppErrors` object with only the `Message` field populated. This can be useful when you want to propagate only the error message without the status code. It returns a pointer to the new `AppErrors` instance with the `Message` copied from `e`.

### `NewNotFoundError` Function
```go
func NewNotFoundError(message string) *AppErrors {
    return &AppErrors{Code: http.StatusNotFound, Message: message}
}
```
- **`NewNotFoundError` function**: This function is a utility to create a `NotFound` error with HTTP status code `404`. It takes a message as input and returns a new `AppErrors` object with:
  - `Code`: set to `http.StatusNotFound` (which is `404`).
  - `Message`: the message passed as a parameter.
  
### `NewServerError` Function
```go
func NewServerError(message string) *AppErrors {
    return &AppErrors{Code: http.StatusInternalServerError, Message: message}
}
```
- **`NewServerError` function**: This function creates a `Server Error` with HTTP status code `500`. It works similarly to `NewNotFoundError`, but the status code is set to `http.StatusInternalServerError` (which is `500`).

### `NewValidationError` Function
```go
func NewValidationError(message string) *AppErrors {
    return &AppErrors{Code: http.StatusUnprocessableEntity, Message: message}
}
```
- **`NewValidationError` function**: This function creates a `Validation Error` with HTTP status code `422`. It is typically used to represent errors when the input data is invalid but the request could still be processed. It sets the status code to `http.StatusUnprocessableEntity` (which is `422`).

### Explanation of the Rules

1. **Custom Error Structure**:
   - The main structure `AppErrors` is a custom error type that includes both an HTTP status code and a message. This allows you to define errors in terms of HTTP status codes, which are typically used in web APIs.

2. **Utility Functions for Common Errors**:
   - Functions like `NewNotFoundError`, `NewServerError`, and `NewValidationError` are utility functions that simplify the creation of common HTTP errors. They provide predefined status codes for specific types of errors.
     - `NewNotFoundError` creates an error for "Not Found" (HTTP `404`).
     - `NewServerError` creates an error for internal server issues (HTTP `500`).
     - `NewValidationError` creates an error for validation problems (HTTP `422`).

3. **Method for Simplified Error Representation**:
   - The `AsMessage` method creates a new error instance with only the `Message` field, ignoring the `Code`. This could be useful when you only care about the error message and not the associated HTTP status code.
   
4. **Separation of Concerns**:
   - By using a custom error type like `AppErrors`, the application can handle errors in a consistent way, where the error message and status code are easily accessible. This separation helps you return more structured and meaningful error responses in APIs.

### Example Usage

1. **Creating and Using a `NotFound` Error**:
   ```go
   err := NewNotFoundError("Resource not found")
   fmt.Println(err.Code)    // 404
   fmt.Println(err.Message) // "Resource not found"
   ```

2. **Creating and Using a `ServerError`**:
   ```go
   err := NewServerError("Something went wrong")
   fmt.Println(err.Code)    // 500
   fmt.Println(err.Message) // "Something went wrong"
   ```

3. **Using `AsMessage` Method**:
   ```go
   err := NewServerError("Something went wrong")
   errMessage := err.AsMessage()
   fmt.Println(errMessage.Message) // "Something went wrong"
   ```

### Conclusion

This code provides a structured way to handle errors in a Go application, especially for web APIs. By using the `AppErrors` struct, you can encapsulate both an HTTP status code and an error message, making error handling more consistent and manageable. The utility functions simplify the creation of common errors, and the `AsMessage` method allows you to easily extract just the message from an error when needed.