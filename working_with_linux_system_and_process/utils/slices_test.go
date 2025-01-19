package utils

import "testing"

func TestSlicesPtrEqual(t *testing.T) {
	ns1 := []int{1, 2, 3}
	ns2 := []int{1, 2, 3}
	if SlicesPtrEqual(ns1, ns2) {
		t.Fatalf("want not equal ptr, got equal ptr")
	}
	if !SlicesPtrEqual(ns1, ns1) {
		t.Fatalf("want equal ptr, got not equal ptr")
	}
	ns3 := ns1[:2]
	if !SlicesPtrEqual(ns1, ns3) {
		t.Fatalf("want equal ptr, got not equal ptr")
	}

	e1 := []int{}
	e2 := []int{}
	if SlicesPtrEqual(e1, e2) {
		t.Fatalf("slices ptr should not be equal")
	}

}
