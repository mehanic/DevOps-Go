package __7

import (
	"errors"
	"fmt"
	"strings"
)

type Stringer interface {
	String() string
}

type Person struct {
	First, Last string
}

func (p Person) String() string {
	return fmt.Sprintf("%s,%s", p.Last, p.First)
}

type StrList []string

func (s StrList) String() string {
	return strings.Join(s, ",")
}

// PrintStringer prints the value of a Stringer to stdout.
func PrintStringer(s Stringer) {
	fmt.Println(s.String())
}

func main1() {
	john := Person{First: "John", Last: "Doak"}
	var nameList Stringer = StrList{"David", "Sarah"}
	PrintStringer(john)
	// Prints: Doak,John
	PrintStringer(nameList) // Prints: David,Sarah
}

//------------

type error interface {
	Error() string
}

//err := errors.New("this is an error")
//err := fmt.Errorf("user %s had an error: %s", user, msg)

//---

func Divide(num int, div int) (int, error) {
	if div == 0 {
		// We return the zero value of int (0) and an error.
		return 0, errors.New("cannot divide by 0")
	}
	return num / div, nil
}
func main() {
	divideBy := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, div := range divideBy {
		res, err := Divide(100, div)
		if err != nil {
			fmt.Printf("100 by %d error: %s\n", div, err)
			continue
		}
		fmt.Printf("100 divided by %d = %d\n", div, res)
	}
}

//-----
