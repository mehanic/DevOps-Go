package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	i := 0
	n := len(s)

	// Step 1: Skip leading whitespaces
	for i < n && s[i] == ' ' {
		i++
	}

	if i == n {
		return 0
	}

	// Step 2: Check sign
	sign := 1
	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		sign = -1
		i++
	}

	// Step 3: Read digits
	num := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		// Step 4: Check for overflow
		if num > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		num = num*10 + digit
		i++
	}

	return sign * num
}



func myAtoi1(s string) int {
	s = strings.TrimLeft(s, " ") // Убираем пробелы слева
	if len(s) == 0 {
		return 0
	}

	sign := 1
	i := 0

	// Определяем знак
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	num := 0

	for ; i < len(s); i++ {
		if !unicode.IsDigit(rune(s[i])) {
			break
		}
		digit := int(s[i] - '0')

		// Проверка переполнения
		if num > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		num = num*10 + digit
	}

	return sign * num
}


func myAtoi2(s string) int {
	re := regexp.MustCompile(`^\s*([+-]?\d+)`)
	match := re.FindStringSubmatch(s)
	if len(match) < 2 {
		return 0
	}

	num, err := strconv.ParseInt(match[1], 10, 64)
	if err != nil {
		// переполнение — обрежем вручную
		if match[1][0] == '-' {
			return math.MinInt32
		}
		return math.MaxInt32
	}

	if num < int64(math.MinInt32) {
		return math.MinInt32
	} else if num > int64(math.MaxInt32) {
		return math.MaxInt32
	}

	return int(num)
}
func main() {
	fmt.Println(myAtoi("   -42"))          // Output: -42
	fmt.Println(myAtoi("4193 with words")) // Output: 4193
	fmt.Println(myAtoi("words and 987"))   // Output: 0
	fmt.Println(myAtoi("-91283472332"))    // Output: -2147483648
}
