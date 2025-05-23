package main

import (
	"fmt"
	"strconv"
)

func comaCode(l []interface{}) string {
	var result string

	for i := 0; i < len(l); i++ {
		var item string

		// Check the type of element, convert to string if not already
		switch v := l[i].(type) {
		case string:
			item = v
		case int:
			item = strconv.Itoa(v) // Convert int to string
		default:
			item = fmt.Sprintf("%v", v) // Convert any other type to string
		}

		// Add comma between items, add 'and' before the last item
		if i < len(l)-1 {
			result += item + ", "
		} else {
			result += "and " + item
		}
	}

	return result
}

func main() {
	spam := []interface{}{"apples", "banana", "tofu", "cats", 2, 6}
	fmt.Println("The given list:")
	fmt.Println(spam)
	fmt.Println("After passing the list to the comaCode function, the list becomes:")
	fmt.Println(comaCode(spam))
}

// The given list:
// [apples banana tofu cats 2 6]
// After passing the list to the comaCode function, the list becomes:
// apples, banana, tofu, cats, 2, and 6
