package utils

import (
	"reflect"
	"unsafe"
)

func SlicesPtrEqual[T any](ts1 []T, ts2 []T) bool {
	if ts1 == nil && ts2 == nil {
		return false
	}
	if len(ts1) == 0 && len(ts2) == 0 {
		return false
	}
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&ts1))
	h2 := (*reflect.SliceHeader)(unsafe.Pointer(&ts2))
	return h1.Data == h2.Data
}
