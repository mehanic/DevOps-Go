package main

import (
	"fmt"
)

func main() {
	// for keeping things easy to read and type-safe
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	for line := range digits[0] {
		// Print a line for each placeholder in digits
		for digit := range digits {
			fmt.Print(digits[digit][line], "  ")
		}
		fmt.Println()
	}
}

// ███  ██   ███  ███  █ █  ███  ███  ███  ███  ███
// █ █   █     █    █  █ █  █    █      █  █ █  █ █
// █ █   █   ███  ███  ███  ███  ███    █  ███  ███
// █ █   █   █      █    █    █  █ █    █  █ █    █
// ███  ███  ███  ███    █  ███  ███    █  ███  ███
// ╭─mehanic at SkyNet in ⌁/algorithm_and_data_structure/example/01-printing-the-digits
// ╰─λ                                                                                                                                                                  0 (0.119s) < 21:02:31
