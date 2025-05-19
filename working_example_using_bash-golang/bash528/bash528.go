package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	SQUARES = 16
	FAIL    = 70
)

var Puzzle = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", " "}
var moves int

// swap exchanges the elements in the puzzle at the given indices
func swap(i, j int) {
	Puzzle[i], Puzzle[j] = Puzzle[j], Puzzle[i]
}

// Jumble shuffles the puzzle randomly
func Jumble() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		pos1 := rand.Intn(SQUARES)
		pos2 := rand.Intn(SQUARES)
		swap(pos1, pos2)
	}
}

// PrintPuzzle displays the current state of the puzzle
func PrintPuzzle() {
	puzpos := 0
	fmt.Println("Enter 'quit' to exit.")
	fmt.Println(",----.----.----.----.")
	for i1 := 0; i1 < 4; i1++ {
		for i2 := 0; i2 < 4; i2++ {
			fmt.Printf("| %2s ", Puzzle[puzpos])
			puzpos++
		}
		fmt.Println("|")
		if i1 < 3 {
			fmt.Println("+----+----+----+----+")
		}
	}
	fmt.Println("'----'----'----'----'")
}

// GetNum retrieves a valid number to move from user input
func GetNum() int {
	var puznum int
	for {
		fmt.Printf("Moves: %d\n", moves)
		fmt.Print("Number to move: ")
		var input string
		fmt.Scanln(&input)

		if input == "quit" {
			fmt.Println()
			os.Exit(0)
		}

		_, err := fmt.Sscanf(input, "%d", &puznum)
		if err != nil || puznum <= 0 || puznum >= SQUARES {
			continue
		}
		break
	}
	return puznum
}

// GetPosFromNum finds the position of the number in the puzzle
func GetPosFromNum(num int) int {
	for pos := 0; pos < SQUARES; pos++ {
		if Puzzle[pos] == fmt.Sprintf("%d", num) {
			return pos
		}
	}
	return -1 // Not found, should not happen
}

// Move attempts to move the puzzle piece and returns success
func Move(pos int) bool {
	if pos > 3 && Puzzle[pos-4] == " " { // Move up
		swap(pos, pos-4)
		return true
	}
	if pos%4 != 3 && Puzzle[pos+1] == " " { // Move right
		swap(pos, pos+1)
		return true
	}
	if pos < 12 && Puzzle[pos+4] == " " { // Move down
		swap(pos, pos+4)
		return true
	}
	if pos%4 != 0 && Puzzle[pos-1] == " " { // Move left
		swap(pos, pos-1)
		return true
	}
	return false
}

// Solved checks if the puzzle is solved
func Solved() bool {
	for pos := 0; pos < SQUARES-1; pos++ {
		if Puzzle[pos] != fmt.Sprintf("%d", pos+1) {
			return false
		}
	}
	return true
}

func main() {
	Jumble()

	for {
		fmt.Println()
		PrintPuzzle()
		fmt.Println()

		for {
			puznum := GetNum()
			puzpos := GetPosFromNum(puznum)
			moves++
			if Move(puzpos) {
				break
			}
		}

		if Solved() {
			break
		}
	}

	fmt.Println()
	PrintPuzzle()
	fmt.Println("BRAVO!")
}

