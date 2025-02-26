package fixarg

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Test_fixarg(t *testing.T) {
	expected := MultiSet{"hello": 1, "js": 1, "world": 1}
	m := NewMultiSet()
	r := strings.NewReader("hello\njs\nworld\n")
	Readfrom(r, InsertFunc(m))
	for key := range m {
		fmt.Println(key)
	}
	if reflect.DeepEqual(expected, m) == false {
		for key, val := range expected {
			t.Error("expected:", key, ",", val)
		}
		for key, val := range m {
			t.Log("found:", key, ",", val)
		}
	}

}
