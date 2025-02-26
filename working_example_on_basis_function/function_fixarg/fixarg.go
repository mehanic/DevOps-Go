package fixarg

import (
	"bufio"
	"io"
)

type MultiSet map[string]int

func NewMultiSet() MultiSet {
	return MultiSet{}
}

func Readfrom(r io.Reader, f func(string)) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
}

func Insert(m MultiSet, val string) {
	m[val] = 1
}

func InsertFunc(m MultiSet) func(val string) {
	return func(val string) {
		Insert(m, val)
	}
}
