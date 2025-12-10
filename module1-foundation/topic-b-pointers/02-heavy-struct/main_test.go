// Assignment 2: The Heavy Struct
//
// Goal: Benchmark passing a [10000]int struct by Value vs Pointer.
//       Measure the speed difference.
//
// Instructions:
// 1. Run: go test -bench=. -benchmem
// 2. Compare the ns/op between value and pointer
// 3. Observe the allocations difference
//
// Run benchmarks: go test -bench=. -benchmem

package main

import "testing"

// HeavyStruct is a large struct (80KB)
type HeavyStruct struct {
	Data [10000]int
}

// ProcessByValue receives the struct by value (copy)
func ProcessByValue(h HeavyStruct) int {
	return h.Data[0] + h.Data[9999]
}

// ProcessByPointer receives the struct by pointer (no copy)
func ProcessByPointer(h *HeavyStruct) int {
	return h.Data[0] + h.Data[9999]
}

func BenchmarkByValue(b *testing.B) {
	h := HeavyStruct{}
	h.Data[0] = 1
	h.Data[9999] = 2

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ProcessByValue(h)
	}
}

func BenchmarkByPointer(b *testing.B) {
	h := &HeavyStruct{}
	h.Data[0] = 1
	h.Data[9999] = 2

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ProcessByPointer(h)
	}
}
