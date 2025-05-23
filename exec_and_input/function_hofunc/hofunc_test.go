package hofunc

import (
	"strings"
	"testing"
)

func TestReadfrom(t *testing.T) {
	expected := "hello\njs\nworld\n"
	r := strings.NewReader(expected)
	data := ""
	Readfrom(r, func(line string) {
		data = data + line + "\n"
	})
	if expected != data {
		t.Error("expected:", expected)
		t.Log("found:", data)
	}
}
