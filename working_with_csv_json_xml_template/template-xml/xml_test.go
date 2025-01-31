package main

import (
	"testing"
)

func BenchmarkCountStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountStruct()
	}
}

func BenchmarkCountDecoder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountDecoder()
	}
}


// go test -bench . -benchmem                                           0 (0.111s) < 16:32:43

// goos: linux
// goarch: amd64
// pkg: cal/csv_json_xml/template-xml
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// BenchmarkCountStruct-12     	   11186	    109339 ns/op	   26280 B/op	     664 allocs/op
// BenchmarkCountDecoder-12    	   13410	     87886 ns/op	   16632 B/op	     529 allocs/op
// PASS
// ok  	cal/csv_json_xml/template-xml	4.298s


