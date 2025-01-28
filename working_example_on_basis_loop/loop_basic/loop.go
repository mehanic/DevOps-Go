package main

import "fmt"

func main() {
	for {
		fmt.Println("Loop iteration")
		break
	}

	isRun := true
	for isRun {
		fmt.Println("Loop iteration with condition")
		isRun = false
	}

	for i := 0; i < 2; i++ {
		fmt.Println("Loop iteration", i)
		if i == 1 {
			continue
		}
	}

	sl := []int{1, 2, 3}
	idx := 0

	for idx < len(sl) {
		fmt.Println("while-stype loop edx:", idx, "value:", sl[idx])
		idx++
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println("c-style loop", i, sl[i])
	}
	for idx := range sl {
		fmt.Println("range slice by index", idx)
	}
	for idx, val := range sl {
		fmt.Println("range slice by idx-value", idx, val)
	}

	profile := map[int]string{1: "Hoasker", 2: "RZ"}

	for key := range profile {
		fmt.Println("range map by key", key)
	}

	for key, val := range profile {
		fmt.Println("range map by key-val", key, val)
	}

	for _, val := range profile {
		fmt.Println("range map by val", val)
	}

	str := "Hello, World!"
	for pos, char := range str {
		fmt.Printf("%#U at pos %d\n", char, pos)
	}
}

// ╰─λ go run loop.go                                                  0 (0.001s) < 23:57:32
// Loop iteration
// Loop iteration with condition
// Loop iteration 0
// Loop iteration 1
// while-stype loop edx: 0 value: 1
// while-stype loop edx: 1 value: 2
// while-stype loop edx: 2 value: 3
// c-style loop 0 1
// c-style loop 1 2
// c-style loop 2 3
// range slice by index 0
// range slice by index 1
// range slice by index 2
// range slice by idx-value 0 1
// range slice by idx-value 1 2
// range slice by idx-value 2 3
// range map by key 1
// range map by key 2
// range map by key-val 2 RZ
// range map by key-val 1 Hoasker
// range map by val Hoasker
// range map by val RZ
// U+0048 'H' at pos 0
// U+0065 'e' at pos 1
// U+006C 'l' at pos 2
// U+006C 'l' at pos 3
// U+006F 'o' at pos 4
// U+002C ',' at pos 5
// U+0020 ' ' at pos 6
// U+0057 'W' at pos 7
// U+006F 'o' at pos 8
// U+0072 'r' at pos 9
// U+006C 'l' at pos 10
// U+0064 'd' at pos 11
// U+0021 '!' at pos 12
