package main

import (
	"sort"
	"testing"
)

func BenchmarkMergeSortConcurrent2(b *testing.B) {
	numElements := 16 << 20
	src := make([]int64, numElements)
	original := make([]int64, numElements)
	prepare(original)
	interSrc = make([]int64, len(src))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(src, original)
		b.StartTimer()
		MergeSortConcurrent2(src)
	}
}

//func BenchmarkMergeSortBasic(b *testing.B) {
//	numElements := 16 << 10
//	src := make([]int64, numElements)
//	original := make([]int64, numElements)
//	prepare(original)
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		b.StopTimer()
//		copy(src, original)
//		b.StartTimer()
//		MergeSortBasic(src)
//	}
//}

func BenchmarkMergeSortBasic2(b *testing.B) {
	numElements := 16 << 20
	src := make([]int64, numElements)
	original := make([]int64, numElements)
	prepare(original)

	interSrc = make([]int64, len(src))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(src, original)
		b.StartTimer()
		MergeSortBasic2(src)
	}
}

func BenchmarkMergeSortConcurrent1(b *testing.B) {
	numElements := 16 << 20
	src := make([]int64, numElements)
	original := make([]int64, numElements)
	prepare(original)
	interSrc = make([]int64, len(src))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(src, original)
		b.StartTimer()
		MergeSortConcurrent1(src)
	}
}

//func TestMergeSort(t *testing.T) {
//	numElements := 16 << 10
//	src := make([]int64, numElements)
//	original := make([]int64, numElements)
//	prepare(original)
//	copy(src, original)
//
//	MergeSortConcurrent2(original)
//	sort.Slice(src, func(i, j int) bool { return src[i] < src[j] })
//	assert.Equal(t, original, src)
//}

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
	numElements := 16 << 20
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
