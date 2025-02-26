package namedresult

import (
	"os"
	"testing"
)

func TestWriteto(t *testing.T) {
	var expected int64
	expected = 12
	lines := []string{"hello", "js", "world"}
	n, err := Writeto(os.Stdout, lines)
	if err != nil {
		t.Error(err)
	}
	if n != expected {
		t.Error("expected:", expected)
		t.Log("found:", n)
	}

}
