package namedfunc

import (
	"testing"
)

func Test_Example(t *testing.T) {
	expected := 3
	sum := func(a int, b int) int {
		return a + b
	}
	result := Example(sum)
	if expected != result {
		t.Error("expected:", expected)
		t.Log("found:", result)
	}
}

func Test_Example2(t *testing.T) {
	expected := 2
	result := Example(Example2(func(a int, b int) int {
		return a + b
	}))
	if expected != result {
		t.Error("expected:", 2)
		t.Log("found:", result)
	}
}
