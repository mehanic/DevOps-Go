package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Configuration
const (
	strength = 2 // Pre-read
	sleep    = 2 // Sleep time
)

var (
	namePlayerA   = "あなた"
	namePlayerB   = "AI"
	aiPlayerA     = "AI"
	aiPlayerB     = "AI"
	colorPlayerA  = 4
	colorPlayerB  = 1
	aikeyword     = "ai"
	colorHover    = 4
	hoverInit     = false
	timestamp     = time.Now().Unix()
	selectedX     = -1
	selectedY     = -1
	selectedNewX  = -1
	selectedNewY  = -1
	A             = -1
	B             = 1
	originY       = 4
	originX       = 7
	hoverX        = 0
	hoverY        = 0
	labelX        = -2
	labelY        = 9

	// Game field and configurations
	cacheLookup    = make(map[string]int)
	cacheFlag      = make(map[string]int)
	cacheDepth     = make(map[string]int)
	redraw         = make(map[string]string)
	field          = make(map[string]int)
	initLine       = []int{4, 2, 3, 6, 5, 3, 2, 4}
	figNames       = []string{"(空)", "ポーン", "ナイト", "ビショップ", "ルーク", "クイーン", "キング"}
	asciiNames     = []string{"k", "q", "r", "b", "n", "p", " ", "P", "N", "B", "R", "Q", "K"}
	figValues      = []int{0, 1, 5, 5, 6, 17, 42}
)

// Get terminal dimensions
func getTerminalDimensions() (int, int) {
	cmd := exec.Command("stty", "size")
	out, err := cmd.Output()
	if err != nil {
		return 24, 80 // Default dimensions
	}
	dimensions := strings.Fields(string(out))
	height, _ := strconv.Atoi(dimensions[0])
	width, _ := strconv.Atoi(dimensions[1])
	return height, width
}

// Initialize game field
func initializeField() {
	for x := 0; x < 8; x++ {
		field[fmt.Sprintf("0,%d", x)] = initLine[x]
		field[fmt.Sprintf("7,%d", x)] = -initLine[x]
	}

	for x := 0; x < 8; x++ {
		field[fmt.Sprintf("1,%d", x)] = 1
		field[fmt.Sprintf("6,%d", x)] = -1
	}

	for y := 2; y < 6; y++ {
		for x := 0; x < 8; x++ {
			field[fmt.Sprintf("%d,%d", y, x)] = 0
		}
	}
}

// Main function
func main() {
	height, width := getTerminalDimensions()
	fmt.Printf("Terminal Size: %d x %d\n", height, width)
	initializeField()

	// Example: Display initial state
	displayField()
}

// Display the chess field
func displayField() {
	fmt.Println("Current Chess Field:")
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			piece := field[fmt.Sprintf("%d,%d", y, x)]
			if piece != 0 {
				index := abs(piece) // Use the absolute value to get the correct index
				fmt.Printf("%s ", figNames[index])
			} else {
				fmt.Print("(空) ")
			}
		}
		fmt.Println()
	}
}

// Example of Negamax algorithm (simplified)
func negamax(depth, alpha, beta, player int, save bool) int {
	if save {
		// Save best move logic here...
	}
	// Check for endgame or depth condition
	if depth <= 0 {
		return evaluateBoard(player)
	}
	bestValue := -9999

	for fromY := 0; fromY < 8; fromY++ {
		for fromX := 0; fromX < 8; fromX++ {
			fig := field[fmt.Sprintf("%d,%d", fromY, fromX)] * player
			if fig <= 0 {
				continue // Skip if not a player's piece
			}
			// Possible moves logic would go here
			// Example:
			for toY := 0; toY < 8; toY++ {
				for toX := 0; toX < 8; toX++ {
					if canMove(fromY, fromX, toY, toX, player) {
						// Move piece
						oldFrom := field[fmt.Sprintf("%d,%d", fromY, fromX)]
						oldTo := field[fmt.Sprintf("%d,%d", toY, toX)]
						field[fmt.Sprintf("%d,%d", fromY, fromX)] = 0
						field[fmt.Sprintf("%d,%d", toY, toX)] = oldFrom

						val := -negamax(depth-1, -beta, -alpha, -player, false)

						// Undo move
						field[fmt.Sprintf("%d,%d", fromY, fromX)] = oldFrom
						field[fmt.Sprintf("%d,%d", toY, toX)] = oldTo

						if val > bestValue {
							bestValue = val
							if save {
								// Save move
							}
						}
						if val > alpha {
							alpha = val
						}
						if alpha >= beta {
							break
						}
					}
				}
			}
		}
	}
	return bestValue
}

// Check if a move is valid (simplified)
func canMove(fromY, fromX, toY, toX, player int) bool {
	// Logic to check if the move is valid
	// Simplified for this example
	return true
}

// Evaluate board (simple heuristic)
func evaluateBoard(player int) int {
	value := 0
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			fig := field[fmt.Sprintf("%d,%d", y, x)]
			if fig != 0 {
				if fig > 0 {
					value += figValues[abs(fig)]
				} else {
					value -= figValues[abs(fig)]
				}
			}
		}
	}
	return value
}

// Utility function to get absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

