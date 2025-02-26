package closure

import (
	"testing"
)

func Test_closure(t *testing.T) {
	result := 0
	expected := 55
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	err := closure(nums, func(nums []int) {
		for _, num := range nums {
			result += num
		}
	})
	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Error("expected:", expected)
		t.Log("found", result)
	}
}
