This Go program demonstrates the use of concurrency, error handling, and graceful shutdown using the `errgroup` package and context management. The program simulates workers performing tasks periodically, with error handling and the ability to handle system termination signals (like `SIGINT`, `SIGTERM`, and `SIGHUP`).

### Key Concepts and Rules:

#### 1. **Context and Signal Handling:**
   - **Signal Notification (`signal.NotifyContext`)**: The program uses `signal.NotifyContext` to listen for system signals (such as `SIGINT`, `SIGTERM`, and `SIGHUP`). These signals are commonly used to indicate that the program should terminate. When one of these signals is received, the `context` is canceled, signaling all goroutines that they should stop.
   - The `context` (`ctx`) is passed to all worker goroutines, allowing them to listen for cancellation and stop if needed.
   
   **Example Behavior:**
   - If you press `Ctrl+C` (which sends `SIGINT`), the program will receive that signal and gracefully stop the worker goroutines.

#### 2. **Error Group (`errgroup.Group`):**
   - **`errgroup.Group`**: This is used to manage multiple goroutines and propagate errors. If any goroutine returns an error, the `errgroup` will stop waiting for other goroutines and immediately return the first error encountered.
   - In this program, `errgroup.Group` is used to manage the worker goroutines, which are responsible for performing tasks. The `eg.Wait()` function is used to block the main goroutine until all the worker goroutines have finished, either successfully or with an error.

#### 3. **Worker Goroutines:**
   - **`eg.Go`**: The `eg.Go` function is used to launch worker goroutines. In this program, there are 10 worker goroutines (controlled by `workersCount`).
   - **Ticker (`time.NewTicker`)**: Each worker uses a `Ticker` to periodically perform work at an interval defined by `workerInterval` (3 seconds).
   - Inside each worker, the `select` statement is used to listen for two things:
     - **`ctx.Done()`**: This will be triggered when the context is canceled (for example, if a termination signal is received). If this happens, the worker stops and returns `nil` to indicate it should gracefully terminate.
     - **`t.C`**: This triggers when the ticker ticks (every 3 seconds in this case), and the worker performs its task by calling `task(ctx, i)`.

#### 4. **Task (`task` function):**
   - The `task` function simulates some work done by each worker. It logs the task ID and randomly returns an error in 30% of cases (simulating failure). 
   - **Random Error**: The line `if rand.Float64() <= 0.3` ensures that 30% of the time, an error is returned, simulating a failure that can happen during the work.
   - **Context Cancellation**: If the context is canceled (for example, if the program is terminated), the worker stops and returns the `ctx.Err()`, which signals that the task was interrupted due to cancellation.
   - If no error occurs, the worker waits for 5 seconds (`time.After(5 * time.Second)`) to simulate completing the task before moving on to the next one.

#### 5. **Graceful Shutdown:**
   - The program listens for termination signals (like `SIGINT`), and when one is received, the context is canceled (`cancel()`), which propagates to all the workers, signaling them to stop.
   - The workers check the `ctx.Done()` channel and exit gracefully when the program is being shut down.

#### 6. **`errgroup.Group` with Context (`WithContext`):**
   - The commented-out code `eg, ctx := errgroup.WithContext(ctx)` is an alternative approach to using `errgroup`. 
     - By calling `errgroup.WithContext(ctx)`, you can bind the `errgroup` to the provided context, which allows workers to listen for cancellation of the context more easily.
     - This is preferred when you need to handle cancellations in multiple goroutines (as in this case). You can experiment by swapping it in the code and see how the program behaves differently.
   - In the provided code, `eg` is created without context and the cancellation handling is done manually through the `signal.NotifyContext`.

### Program Flow:

1. **Signal Handling**: The program begins by setting up the signal handler (`signal.NotifyContext`), which listens for termination signals (`SIGINT`, `SIGTERM`, `SIGHUP`). If one of these signals is received, the context is canceled, which propagates to all workers, notifying them to stop.
   
2. **Worker Initialization**: The program launches 10 workers using `eg.Go`. Each worker runs the `task` function every 3 seconds.
   
3. **Task Execution**: Each worker performs a task:
   - Every 3 seconds, a worker either:
     - Performs a task successfully (waits for 5 seconds), or
     - Randomly encounters an error (30% chance).
   
4. **Graceful Shutdown**:
   - If an error occurs in any worker, it returns an error through the `errgroup.Group`, and the program stops waiting for further tasks.
   - If the program receives a termination signal (e.g., `Ctrl+C`), the context is canceled, and all workers stop gracefully.

5. **Error Handling**:
   - If any worker encounters an error, the `errgroup.Group` will propagate it, and the program will stop waiting for other workers, printing the error and terminating.

### Example Output:

```
0 : do useful work...
1 : do useful work...
2 : do useful work...
3 : do useful work...
4 : do useful work...
5 : do useful work...
6 : do useful work...
7 : do useful work...
8 : do useful work...
9 : do useful work...
```

Or, if an error happens (randomly, 30% of the time):

```
2 : do useful work...
3 : do useful work...
4 : do useful work...
5 : do useful work...
6 : do useful work...
7 : do useful work...
8 : do useful work...
9 : do useful work...
1 : do useful work...
task 3: unknown error
```

If the program receives a termination signal:

```
signal received, shutting down...
```

### Key Takeaways:

1. **Context Management**: The use of `context` ensures that the program can handle cancellation signals and stop the workers gracefully.
2. **Concurrency**: The program runs multiple worker goroutines concurrently, each performing a task periodically.
3. **Error Handling**: The `errgroup.Group` is used to manage multiple goroutines and handle errors, ensuring that the program stops if any worker encounters an error.
4. **Graceful Shutdown**: By listening for system signals, the program can shut down all workers gracefully when it needs to terminate.