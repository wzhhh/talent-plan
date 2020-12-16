package main

import (
	"sort"
	"testing"
)

func BenchmarkMergeSort(b *testing.B) {
	numElements := 16 << 10
	src := make([]int64, numElements)
	original := make([]int64, numElements)
	prepare(original)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(src, original)
		b.StartTimer()
		MergeSort(src)
	}
}

//func BenchmarkMergeSort0(b *testing.B) {
//	numElements := 16 << 2
//	src := make([]int64, numElements)
//	original := make([]int64, numElements)
//	prepare(original)
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		b.StopTimer()
//		copy(src, original)
//		b.StartTimer()
//		MergeSort0(src)
//	}
//}

func BenchmarkNormalSort(b *testing.B) {
	numElements := 16 << 10
	src := make([]int64, numElements)
	original := make([]int64, numElements)
	prepare(original)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(src, original)
		b.StartTimer()
		sort.Slice(src, func(i, j int) bool { return src[i] < src[j] })
	}
}
