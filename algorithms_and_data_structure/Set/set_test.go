package set

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func clone(ns []int) []int {
	cns := make([]int, len(ns))
	copy(cns, ns)
	return cns
}

func values(s Set) []int {
	var vs []int
	for k := range s {
		vs = append(vs, k)
	}
	sort.Ints(vs)
	return vs
}

func equalIgnoreOrder(ns1, ns2 []int) bool {
	if len(ns1) == 0 && len(ns2) == 0 {
		return true
	}
	cns1 := clone(ns1)
	cns2 := clone(ns2)
	sort.Ints(cns1)
	sort.Ints(cns2)
	return reflect.DeepEqual(cns1, cns2)
}

func is(set Set, vals []int) bool {
	cvals := clone(vals)
	svals := values(set)
	return equalIgnoreOrder(svals, cvals)
}

func TestNew(t *testing.T) {
	tests := []struct {
		in  []int
		exp []int
	}{
		{
			in:  []int{1, 2, 3},
			exp: []int{1, 2, 3},
		},
		{
			in:  nil,
			exp: nil,
		},
		{
			in:  []int{},
			exp: []int{},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := New(test.in...)
			if !is(res, test.exp) {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		in     []int
		insert []int
		exp    []int
	}{
		{
			in:     []int{1, 2, 3},
			insert: []int{2, 1, 4, 6},
			exp:    []int{1, 2, 3, 4, 6},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			set := New(test.in...)
			set.Insert(test.insert...)
			if !is(set, test.exp) {
				t.Fatalf("want %v, have %v", test.exp, set)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		in  []int
		del []int
		exp []int
	}{
		{
			in:  []int{1, 2, 3},
			del: []int{2, 1, 4, 6},
			exp: []int{3},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			set := New(test.in...)
			set.Delete(test.del...)
			if !is(set, test.exp) {
				t.Fatalf("want %v, have %v", test.exp, set)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		in       []int
		contains []int
		exp      bool
	}{
		{
			in:       []int{1, 2, 3},
			contains: []int{1, 2},
			exp:      true,
		},
		{
			in:       []int{1, 2, 3},
			contains: []int{4, 5, 8},
			exp:      false,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			set := New(test.in...)
			res := set.Contains(test.contains...)
			if test.exp != res {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		in       []int
		contains []int
		exp      bool
	}{
		{
			in:       []int{1, 2, 3},
			contains: []int{1, 6, 7},
			exp:      true,
		},
		{
			in:       []int{1, 2, 3},
			contains: []int{4, 5, 8},
			exp:      false,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			set := New(test.in...)
			res := set.ContainsAny(test.contains...)
			if test.exp != res {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}

func TestValues(t *testing.T) {
	tests := []struct {
		in  []int
		exp []int
	}{
		{
			in:  []int{1, 2, 3},
			exp: []int{1, 2, 3},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			set := New(test.in...)
			res := set.Values()
			if !equalIgnoreOrder(res, test.exp) {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		in1 []int
		in2 []int
		exp []int
	}{
		{
			in1: []int{1, 2, 3},
			in2: []int{3, 4, 5, 6},
			exp: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := Union(New(test.in1...), New(test.in2...))
			if !is(res, test.exp) {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		in1 []int
		in2 []int
		exp []int
	}{
		{
			in1: []int{1, 2, 3},
			in2: []int{3, 4, 5, 6},
			exp: []int{3},
		},
		{
			in1: []int{1, 2, 3},
			in2: []int{6, 4, 5, 6},
			exp: []int{},
		},
		{
			in1: []int{1, 2, 3},
			in2: []int{3, 4, 5, 6, 4, 6, 2},
			exp: []int{2, 3},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := Intersection(New(test.in1...), New(test.in2...))
			if !is(res, test.exp) {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}
