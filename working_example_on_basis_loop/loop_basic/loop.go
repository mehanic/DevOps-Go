package main

import "fmt"

// Глобальные переменные, доступные во всех функциях
var sl = []int{1, 2, 3}                     // Срез
var profile = map[int]string{1: "Hoasker", 2: "RZ"} // Карта

func main() {
	// Бесконечный цикл for + break
	for {
		fmt.Println("Loop iteration")
		break
	}

	loopWithCondition("loopWithCondition")
	fmt.Println("------")
	loopWithContinue("loopWithContinue")
	fmt.Println("------")
	whileStyleLoop("whileStyleLoop")
	fmt.Println("------")
	cStyleLoop("cStyleLoop")
	fmt.Println("------")
	rangeOverSlice("rangeOverSlice")
	fmt.Println("------")
	rangeOverSliceWithValues("rangeOverSliceWithValues")
	fmt.Println("------")
	rangeOverMapKeys("rangeOverMapKeys")
	fmt.Println("------")
	rangeOverMapKeysAndValues("rangeOverMapKeysAndValues")
	fmt.Println("------")
	rangeOverMapValues("rangeOverMapValues")
	fmt.Println("------")
	rangeOverString("rangeOverString")
}

// Цикл for с логическим условием
func loopWithCondition(name string) {
	isRun := true
	for isRun {
		fmt.Println(name, "Loop iteration with condition")
		isRun = false
	}
}

// for с continue
func loopWithContinue(name string) {
	for i := 0; i < 2; i++ {
		fmt.Println(name, "Loop iteration", i)
		if i == 1 {
			continue
		}
	}
}

// while-стиль цикл (for с логическим условием)
func whileStyleLoop(name string) {
	idx := 0
	for idx < len(sl) {
		fmt.Println(name, "while-style loop idx:", idx, "value:", sl[idx])
		idx++
	}
}

// for в C-стиле
func cStyleLoop(name string) {
	for i := 0; i < len(sl); i++ {
		fmt.Println(name, "c-style loop", i, sl[i])
	}
}

// range по срезу (индексы)
func rangeOverSlice(name string) {
	for idx := range sl {
		fmt.Println(name, "range slice by index", idx)
	}
}

// range по срезу (индексы + значения)
func rangeOverSliceWithValues(name string) {
	for idx, val := range sl {
		fmt.Println(name, "range slice by idx-value", idx, val)
	}
}

// range по map (только ключи)
func rangeOverMapKeys(name string) {
	for key := range profile {
		fmt.Println(name, "range map by key", key)
	}
}

// range по map (ключи + значения)
func rangeOverMapKeysAndValues(name string) {
	for key, val := range profile {
		fmt.Println(name, "range map by key-val", key, val)
	}
}

// range по map (только значения)
func rangeOverMapValues(name string) {
	for _, val := range profile {
		fmt.Println(name, "range map by val", val)
	}
}

// range по строке (Unicode-символы)
func rangeOverString(name string) {
	str := "Hello, World!"
	for pos, char := range str {
		fmt.Printf("%s %#U at pos %d\n", name, char, pos)

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
