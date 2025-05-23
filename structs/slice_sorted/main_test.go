package lesson3

import "testing"

func assertEquals(t *testing.T, expected []Token, actual []Token) bool {
	if len(expected) != len(actual) {
		return false
	}

	for i := range expected {
		if expected[i] != actual[i] {
			t.Logf("Compare error: expected = %v, actual = %v\n", expected[i], actual[i])
			return false
		}
	}

	return true
}

func testWithCase(t *testing.T, expectedTokens []Token, testString string) {
	t.Logf("Testing string: %q", testString)
	tokens, err := PrintTenMajorTokens(testString)

	if err != nil {
		t.Errorf("Test string: %s. Error: %s.", testString, err)
	}

	if !assertEquals(t, expectedTokens, tokens) {
		t.Errorf("Expected: %+v;\n Actual: %+v", expectedTokens, tokens)
	}
}

func TestPrintTenMajorTokens(t *testing.T) {
	expectedTokens := []Token{
		Token{"three", 3},
		Token{"two", 2},
		Token{"one", 1},
	}

	testWithCase(t, expectedTokens, "one two three two three three")
	testWithCase(t, make([]Token, 0), "")

	expectedTokens = []Token{
		Token{"11", 11},
		Token{"10", 10},
		Token{"9", 9},
		Token{"8", 8},
		Token{"7", 7},
		Token{"6", 6},
		Token{"5", 5},
		Token{"4", 4},
		Token{"3", 3},
		Token{"2", 2},
	}

	testString :=
		"11 10 9 8 7 6 5 4 3 2 1 " +
		"11 10 9 8 7 6 5 4 3 2 " +
		"11 10 9 8 7 6 5 4 3 " +
		"11 10 9 8 7 6 5 4 " +
		"11 10 9 8 7 6 5 " +
		"11 10 9 8 7 6 " +
		"11 10 9 8 7 " +
		"11 10 9 8 " +
		"11 10 9 " +
		"11 10 " +
		"11"

	testWithCase(t, expectedTokens, testString)
}
