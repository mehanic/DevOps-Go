package main

import (
	"fmt"
)

func main() {
	// Циклы
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("func num")
	fmt.Println("sum = 45 — потому что ты складываешь числа от 0 до 9.", sum)

	sum2 := 0
	for sum2 < 100 { // аналог while // Пока sum меньше 100 — выполняем
		sum2 += 10
	}
	fmt.Println("func 2")
	fmt.Println("sum2 =", sum2)

	sum3 := 0
	for sum3 < 10 { // аналог while
		sum3 += 10
	}
	fmt.Println("func 3")
	fmt.Println("sum3 =", sum3)

	fmt.Println("--------------------------------------------")
	// Ветвления
	a := 0
	for a < 10 {
		if a == 10 {
			fmt.Println("a is 10")
		} else {
			fmt.Println(fmt.Sprintf("a is not 10. a=%d", a))
		}
		a++
	}
	// Это цикл от 0 до 9. Проверяется: если a == 10, вывести одно сообщение, иначе — другое. Но условие a == 10 никогда не выполняется, так как цикл до 10 не включительно. Поэтому будет всегда a is not 10.

	fmt.Println("--------------------------------------------")

	a2 := 0
	for a2 < 10 {
		i := isTest(a2)
		fmt.Printf("isTest(%d) = %d -> ", a2, i)
		if i == 1 {
			fmt.Println("a2 = 2")
		} else if i == 2 {
			fmt.Println("a2 = 3")
		} else {
			fmt.Println(fmt.Sprintf("unknown a2. a2=%d", a2))
		}
		a2++
	}
	//	Классический if, else if, else
	// Использует отдельную переменную i := isTest(a2) перед if.
	// Логика: если a2 == 2 → i = 1, если a2 == 3 → i = 2, иначе i = 3.

	fmt.Println("--------------------------------------------")

	a3 := 0
	for a3 < 10 {
		if i := isTest(a3); i == 1 {
			fmt.Println("a3 = 2")
		} else if i == 2 {
			fmt.Println("a3 = 3")
		} else {
			fmt.Println(fmt.Sprintf("unknown a3. a3=%d", a3))
		}
		a3++
	}

	// Тот же результат, что и a2, но:
	// Внутри if, переменная i создаётся в той же строке (if i := ...; i == ...)
	// Это короче и локальнее (i существует только внутри if)
	fmt.Println("--------------------------------------------")

	a4 := 0
	for a4 < 10 {
		j := isTest(a4)
		switch j {
		case 1:
			fmt.Println("a4 = 2")
		case 2:
			fmt.Println("a4 = 3")
		default:
			fmt.Println(fmt.Sprintf("unknown a4. a4=%d", a4))
		}
		a4++
	}

	// Заменили if/else на switch.
	// Переменная j объявляется вне switch.
	// Хорошо читается, если много случаев (case 1, case 2 и т.д.).

	fmt.Println("--------------------------------------------")

	a5 := 0
	for a5 < 10 {
		switch j := isTest(a5); j {
		case 1:
			fmt.Println("a5 = 2")
		case 2:
			fmt.Println("a5 = 3")
		default:
			fmt.Println(fmt.Sprintf("unknown a5. a5=%d", a5))
		}
		a5++
	}

	// Тот же switch, но переменная j объявляется внутри switch (в одной строке).Аналог if i := ...; но для switch. Локальность переменной j.
}

// Циклы
func isTest(a int) int {
	if a == 2 {
		return 1
	} else if a == 3 {
		return 2
	}
	return 3
}
