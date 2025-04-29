package otus_lesson2

import (
	"fmt"
	"testing"
)

/*
* "a4bc2d5e" => "aaaabccddddde"
* "abcd" => "abcd"
* "45"` => "" (некорректная строка)
* "qwe\4\5" => "qwe45" (*)
* "qwe\45" => "qwe44444" (*)
* "qwe\\5" => "qwe\\\\\" (*)
 */

func TestUnpack(t *testing.T) {
	testItems := prepareTestItems()

	for i := range testItems {
		testItem := testItems[i]
		fmt.Printf("Testing string: \"%s\".\n", testItem.in)
		result, err := Unpack(testItem.in)

		if err != nil && !testItem.expectError {
			t.Fatalf("Error unpacking for input string \"%s\": expected \"%s\", actual error: %s", testItem.in, testItem.out, err)
		}

		if result != testItem.out {
			t.Fatalf("Error unpacking string \"%s\": expected \"%s\", actual \"%s\"", testItem.in, testItem.out, result)
		}
	}
}

func prepareTestItems() []testItem {
	return []testItem{
		{in: "a4bc2d5e", out: "aaaabccddddde"},
		{in: "abcd", out: "abcd"},
		{in: "45", out: "", expectError: true},
		{in: "qwe\\4\\5", out: "qwe45"},
		{in: "qwe\\45", out: "qwe44444"},
		{in: "qwe\\\\5", out: "qwe\\\\\\\\\\"},
	}
}

type testItem struct {
	in string
	out string
	expectError bool
}