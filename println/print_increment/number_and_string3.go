package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	counter, factor := 45, 0.5

	counter++
	counter++
	counter++
	counter++
	counter++

	factor--
	factor--

	fmt.Println(float64(counter) * factor)
	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
	main7()
	main8()
	main9()
}

func main1() {
	var counter int

	counter++
	counter--

	counter += 5
	counter *= 10
	counter /= 5

	fmt.Println(counter)
}

func main2() {
	width, height := 10, 2

	width++
	width += height
	width--
	width -= height
	width *= 20
	width /= 25
	width %= 5

	fmt.Println(width)
}

func main3() {
	var (
		radius = 10.
		area   float64
	)

	area = math.Pi * radius * radius

	fmt.Printf("radius: %g -> area: %.2f\n",
		radius, area)

	// ALTERNATIVE:
	// math.Pow calculates the power of a float number
	// area = math.Pi * math.Pow(radius, 2)
}

func main4() {
	// The type of a string and a raw string literal
	// is the same. They both are strings.
	//
	// So, they both can be used as a string value.
	var s string
	s = "how are you?"
	s = `how are you?`
	fmt.Println(s)

	// string literal
	s = "<html>\n\t<body>\"Hello\"</body>\n</html>"
	fmt.Println(s)

	// raw string literal
	s = `
<html>
	<body>"Hello"</body>
</html>`

	fmt.Println(s)

	// windows path
	fmt.Println("c:\\my\\dir\\file") // string literal
	fmt.Println(`c:\my\dir\file`)    // raw string literal
}

func main5() {
	name, last := "carl", "sagan"

	fmt.Println(name + " " + last)
}

func main6() {
	name, last := "carl", "sagan"

	// assignment operation using string concat
	name += " edward"

	// equals to this:
	// name = name + " edward"

	fmt.Println(name + " " + last)
}

func main7() {
	fmt.Println(
		"hello" + ", " + "how" + " " + "are" + " " + "today?",
	)

	// you can combine raw string and string literals
	fmt.Println(
		`hello` + `, ` + `how` + ` ` + `are` + ` ` + "today?",
	)

	// ------------------------------------------
	// Converting non-string values into string
	// ------------------------------------------

	eq := "1 + 2 = "
	sum := 1 + 2

	// invalid op
	// string concat op can only be used with strings
	// fmt.Println(eq + sum)

	// you need to convert it using strconv.Itoa
	// Itoa = Integer to ASCII

	fmt.Println(eq + strconv.Itoa(sum))

	//

	// invalid op
	// eq = true + " " + false

	eq = strconv.FormatBool(true) +
		" " +
		strconv.FormatBool(false)

	fmt.Println(eq)
}

func main8() {
	msg := `
	
	The weather looks good.
I should go and play.

	`

	fmt.Println(strings.TrimSpace(msg))
}

func main9() {
	name := "inanç           "

	name = strings.TrimRight(name, " ")
	l := utf8.RuneCountInString(name)

	fmt.Println(l)
}


// -75
// 10
// 3
// radius: 10 -> area: 314.16
// how are you?
// <html>
// 	<body>"Hello"</body>
// </html>

// <html>
// 	<body>"Hello"</body>
// </html>
// c:\my\dir\file
// c:\my\dir\file
// carl sagan
// carl edward sagan
// hello, how are today?
// hello, how are today?
// 1 + 2 = 3
// true false
// The weather looks good.
// I should go and play.
// 5
