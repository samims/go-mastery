package main

import (
	"testing"
)

func BenchmarkCreateDataEscape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewHeap()
	}
}

func BenchmarkCreateDataFixed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewStack()
	}
}
