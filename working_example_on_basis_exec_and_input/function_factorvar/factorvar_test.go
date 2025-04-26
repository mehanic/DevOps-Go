package factorvar

import (
	"os"
	"testing"
)

func TestWriteto(t *testing.T) {
	expected := int64(12)
	n, err := Writeto(os.Stdout, "hello", "js", "world")
	if err != nil {
		t.Error(err)
	}
	if expected != n {
		t.Error("expected:", expected)
		t.Log("found:", n)
	}
}
