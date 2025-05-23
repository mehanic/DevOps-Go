This Go code demonstrates how to manage concurrent HTTP requests using the `errgroup` package from `golang.org/x/sync/errgroup`. It performs concurrent HTTP GET requests for a list of URLs, logs the request times and errors, and waits for all requests to complete before handling errors.

### Key Concepts and Rules

1. **Error Group (`errgroup.Group`)**:
   - The `errgroup.Group` is used to manage a group of goroutines and handle their errors.
   - The `eg.Go()` method starts a new goroutine for each task (in this case, for each HTTP request).
   - The `eg.Wait()` method blocks until all the goroutines complete, collecting and returning the first encountered error. If any error occurs in the goroutines, `Wait()` will return that error.

2. **Concurrency with Goroutines**:
   - The `visitor` function returns a closure that performs an HTTP request. Each closure is executed in a separate goroutine using `eg.Go(visitor(url))`.
   - This allows the program to make HTTP requests concurrently (i.e., multiple requests are processed in parallel).

3. **Logging Request Duration**:
   - Each goroutine logs the URL being requested, the time it took to make the request, and any errors that occurred. The `defer` statement is used to ensure that the logging happens after the HTTP request completes, whether it succeeds or fails.
   - The `time.Since(s)` computes the elapsed time from when the request started.

4. **HTTP Requests**:
   - The `http.Get(url)` function sends a GET request to the specified URL.
   - If an error occurs during the request (for example, if the URL is invalid), the error is returned, and the goroutine will terminate early.
   - After a successful HTTP request, the response body is closed (`resp.Body.Close()`), which is crucial to free up resources (such as network connections) associated with the HTTP response.

5. **Error Handling**:
   - The error handling is managed by the `errgroup.Group`. If any error occurs in any of the goroutines (such as a failed HTTP request), it is captured by the `errgroup`.
   - If an error occurs in any goroutine, it will be passed to `eg.Wait()`, which will return the first error that was encountered. If there is an error, the program logs the error using `log.Fatalln()` and terminates. `log.Fatalln()` logs the error message and exits the program.

6. **Deferred Cleanup**:
   - The `defer` statement ensures that the elapsed time and error (if any) are logged after each HTTP request, regardless of whether the request was successful or failed.

### Detailed Explanation of Code Flow

1. **`visitor(url string)`**:
   - This function returns a closure (`func() error`) that performs the HTTP GET request to the given `url`.
   - The closure measures the time it takes to send the request using `time.Since(s)`.
   - The `defer` ensures that once the HTTP request completes (either successfully or with an error), the URL, duration, and any error are logged.
   - The actual HTTP GET request is made with `http.Get(url)`. If the request is successful, it closes the response body using `resp.Body.Close()`.

2. **`main()`**:
   - The `main` function creates an error group (`errgroup.Group{}`).
   - A list of URLs (`urlList`) is defined, which contains both valid and invalid URLs.
   - The loop iterates over the URLs and starts a new goroutine for each one by calling `eg.Go(visitor(url))`. This means that each URL will be processed concurrently.
   - The `eg.Wait()` function waits for all the goroutines to finish. If any of them encounter an error (e.g., a failed HTTP request), the first encountered error is returned by `eg.Wait()`.
   - If there’s an error, the program logs the error and terminates using `log.Fatalln("Error:", err)`.

### How It Works:
1. **URL List**:
   The `urlList` contains three URLs:
   - `http://www.golang.org/` (valid URL)
   - `http://invalidwebsite.hey/` (invalid URL, will likely cause an error)
   - `http://www.google.com/` (valid URL)

2. **Concurrent Requests**:
   The program loops over `urlList` and calls `eg.Go(visitor(url))` for each URL, which starts an HTTP request for each URL concurrently.
   - Each goroutine will execute the `visitor` function, sending an HTTP request and logging the result.

3. **Logging the Results**:
   After each request completes, the `defer` in `visitor` logs the URL, the duration of the request, and any error that occurred.

4. **Error Group**:
   The `eg.Wait()` ensures that the main function will wait until all requests are finished. If any goroutine encounters an error (such as the invalid URL), the error is captured and returned by `eg.Wait()`.

5. **Program Exit on Error**:
   If an error occurs in any of the goroutines, it will be passed to the main function, where it is logged and the program exits using `log.Fatalln()`.

### Example Output:

The output of the program would look like this (with actual times varying):

```
http://www.golang.org/ 300ms <nil>
http://invalidwebsite.hey/ 500ms Get http://invalidwebsite.hey/: dial tcp: lookup invalidwebsite.hey: no such host
http://www.google.com/ 400ms <nil>
Error: Get http://invalidwebsite.hey/: dial tcp: lookup invalidwebsite.hey: no such host
```

### Key Takeaways:

- **Concurrency**: The program uses `errgroup` to handle concurrent HTTP requests. Each request is processed in its own goroutine.
- **Error Handling**: Any errors that occur in the goroutines are captured by the `errgroup` and reported after all goroutines have finished.
- **Deferred Logging**: The `defer` ensures that logging occurs after the request finishes, regardless of success or failure.
- **Resource Management**: The program ensures that HTTP response bodies are properly closed, preventing resource leaks.