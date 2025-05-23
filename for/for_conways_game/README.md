The provided code implements the **Game of Life** (a cellular automaton) using Go. The Game of Life is a simple simulation of cellular growth and decay, where each cell has certain rules governing its survival or death, based on its neighbors. Here's an explanation of how this code works, focusing on the Game of Life's rules:

### Game of Life Rules:
- A **live cell** with fewer than two live neighbors **dies** (underpopulation).
- A **live cell** with two or three live neighbors **lives** (survival).
- A **live cell** with more than three live neighbors **dies** (overpopulation).
- A **dead cell** with exactly three live neighbors **becomes alive** (reproduction).

### Explanation of the Code:

1. **Constants**:
   ```go
   const (
       height    = 10
       width     = 10
       countShow = 300
   )
   ```
   - `height` and `width`: Define the dimensions of the universe (the grid of cells). This is a 10x10 grid.
   - `countShow`: The number of generations (steps) the game will run before stopping.

2. **`Universe` Type**:
   ```go
   type Universe [][]bool
   ```
   - `Universe` is a 2D slice of booleans (`[][]bool`), where each boolean represents the state of a cell (alive or dead).
   - `true` means the cell is **alive**, and `false` means the cell is **dead**.

3. **`NewUniverse` Function**:
   ```go
   func NewUniverse() Universe {
       newUniverse := make(Universe, width)
       for k, _ := range newUniverse {
           newUniverse[k] = make([]bool, width)
       }
       return newUniverse
   }
   ```
   - This function creates a new 10x10 grid (2D slice) with all cells initialized to `false` (dead).

4. **`Show` Function**:
   ```go
   func (u Universe) Show() {
       for y := 0; y < height; y++ {
           for x := 0; x < width; x++ {
               if u[y][x] == true {
                   fmt.Print("* ")
               } else {
                   fmt.Print("_ ")
               }
           }
           fmt.Print("\n")
       }
   }
   ```
   - This function prints the current state of the universe to the console.
   - Alive cells are displayed as `*` and dead cells as `_`.

5. **`Seed` Function**:
   ```go
   func (u Universe) Seed() {
       for y := 0; y < rand.Intn(height); y++ {
           for x := 0; x < rand.Intn(width); x++ {
               u[y][x] = true
           }
           fmt.Print("\n")
       }
   }
   ```
   - This function randomly seeds the universe with live cells.
   - `rand.Intn(height)` and `rand.Intn(width)` are used to generate random numbers for the positions of live cells.

6. **`Alive` Function**:
   ```go
   func (u Universe) Alive(x, y int) bool {
       x = (x + width) % width
       y = (y + height) % height
       return u[y][x]
   }
   ```
   - This function checks if a cell at position `(x, y)` is alive.
   - The coordinates `(x, y)` are wrapped around the grid using modulo operations to handle the edge cases where a cell is on the border of the universe.

7. **`Neighbors` Function**:
   ```go
   func (u Universe) Neighbors(x, y int) int {
       n := 0
       for v := -1; v <= 1; v++ {
           for h := -1; h <= 1; h++ {
               if !(v == 0 && h == 0) && u.Alive(x+h, y+v) {
                   n++
               }
           }
       }
       return n
   }
   ```
   - This function counts the number of live neighbors for a given cell at `(x, y)`.
   - It iterates over the neighboring cells (in a 3x3 area), and the `!(v == 0 && h == 0)` condition ensures that the cell itself is not counted.

8. **`Next` Function**:
   ```go
   func (u Universe) Next(x, y int) bool {
       n := u.Neighbors(x, y)
       return n == 3 || n == 2 && u.Alive(x, y)
   }
   ```
   - This function applies the rules of the Game of Life to determine if a cell should be alive in the next generation.
   - If a cell has exactly 3 live neighbors, it becomes or stays alive.
   - If a cell has 2 live neighbors, it stays alive if it was alive, otherwise it stays dead.

9. **`Step` Function**:
   ```go
   func Step(a, b Universe) {
       for y := 0; y < height; y++ {
           for x := 0; x < width; x++ {
               b[y][x] = a.Next(x, y)
           }
       }
   }
   ```
   - This function computes the next generation of the universe.
   - It iterates through every cell and applies the `Next` function to update the state of the universe.
   - It updates the universe `b` based on the state of `a` from the previous generation.

10. **`CheckAliveUniverse` Function**:
   ```go
   func (u Universe) CheckAliveUniverse() bool {
       for y := 0; y < height; y++ {
           for x := 0; x < width; x++ {
               if u[y][x] == true {
                   return true
               }
           }
       }
       return false
   }
   ```
   - This function checks if there are any live cells in the universe. If there are no live cells (`false`), the simulation will stop.

11. **`main` Function**:
   ```go
   func main() {
       a, b := NewUniverse(), NewUniverse()
       a.Seed()

       for i := 0; i < countShow; i++ {
           Step(a, b)
           if !a.CheckAliveUniverse() { // если все false - останавливаем цикл
               break
           }
           a.Show()
           fmt.Print("\x0c")
           time.Sleep(time.Second)
           a, b = b, a
       }
   }
   ```
   - The main function initializes two universes (`a` and `b`), seeds `a` with live cells, and then runs the simulation for `countShow` iterations.
   - It repeatedly computes the next generation using `Step`, checks if there are any live cells, and displays the current universe.
   - The `fmt.Print("\x0c")` clears the screen to simulate an animation effect, and `time.Sleep(time.Second)` pauses for one second between generations.
   - The universes `a` and `b` are swapped after each iteration.

### Summary:
The code simulates the **Game of Life** using a 10x10 grid, where each cell follows the rules of life and death based on its neighbors. The universe evolves over multiple generations, and the game stops when there are no live cells left. The grid is randomly seeded, and the program visualizes the evolving universe at each step.