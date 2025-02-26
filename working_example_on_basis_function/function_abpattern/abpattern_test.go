package abpattern

import (
	"testing"
)

func Test_abpattern(t *testing.T) {
	expected := VertexID(1)
	f := NewVertexIDGenerator()
	result := f()
	if expected != result {
		t.Error("expected:", expected)
		t.Log("found:", result)
	}
}
