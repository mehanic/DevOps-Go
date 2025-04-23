// "Decrypt the quote from Julius Caesar: 'L fdph, L vdz, L frqtxhuhg.'

// Your program should shift the letters of both upper and 
// lower case by -3. Remember that 'a' becomes 'x', 'b' becomes 
// 'y', and 'c' becomes 'z'. The same applies to uppercase letters."

package main

import "fmt"

func main() {
	msg := "L fdph, L vdz, L frqtxhuhg"

	for _, k := range msg {
		if k >= 'a' && k <= 'z' {
			k -= 3
			if k < 'a' {
				k += 26
			}
		} else if k >= 'A' && k <= 'Z' {
			k -= 3
			if k < 'A' {
				k += 26
			}
		}

		fmt.Printf("%c", k)
	}
}
