Let's generate some examples for the function `columnId(columnNumber int)`, which converts a column number into an Excel-style column name.

### **Understanding the Algorithm**
This function converts a column number (1-based index) into an Excel-like column name, similar to how Excel labels columns (A, B, ..., Z, AA, AB, ..., ZZ, AAA, etc.).

### **Examples**
| Column Number | Column Name |
|--------------|------------|
| 1            | A          |
| 2            | B          |
| 3            | C          |
| 26           | Z          |
| 27           | AA         |
| 28           | AB         |
| 52           | AZ         |
| 53           | BA         |
| 702          | ZZ         |
| 703          | AAA        |
| 704          | AAB        |
| 18278        | ZZZ        |

### **Example Run in Go**
If we run the function with different inputs, we get:

```go
package main

import (
	"fmt"
)

func columnId(columnNumber int) string {
	if columnNumber <= 0 {
		return ""
	}

	result := ""

	for columnNumber > 0 {
		remainder := columnNumber % 26

		if remainder == 0 {
			remainder = 26
			columnNumber -= 1
		}

		result = string(remainder+'A'-1) + result
		columnNumber /= 26
	}

	return result
}

func main() {
	testCases := []int{1, 2, 3, 26, 27, 28, 52, 53, 702, 703, 704, 18278}

	for _, num := range testCases {
		fmt.Printf("Column Number: %d => Column Name: %s\n", num, columnId(num))
	}
}
```

### **Output**
```
Column Number: 1 => Column Name: A
Column Number: 2 => Column Name: B
Column Number: 3 => Column Name: C
Column Number: 26 => Column Name: Z
Column Number: 27 => Column Name: AA
Column Number: 28 => Column Name: AB
Column Number: 52 => Column Name: AZ
Column Number: 53 => Column Name: BA
Column Number: 702 => Column Name: ZZ
Column Number: 703 => Column Name: AAA
Column Number: 704 => Column Name: AAB
Column Number: 18278 => Column Name: ZZZ
```

Would you like any modifications or further explanations? 😊