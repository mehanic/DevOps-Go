package namedtype

import "testing"

func Test_namedtype(t *testing.T) {
	expected := EdgeID(1)
	egen := NewEdgeIDGenerator()
	vgen := NewVertexIDGenerator()
	result := NewEdge(egen())
	result2 := NewEdge(gen()) //error

	if expected != result {
		t.Error("expected:", expected)
		t.Log("found:", result)
	}

	if expected != result2 {
		t.Error("expected:", expected)
		t.Log("found:", result)
	}
}
