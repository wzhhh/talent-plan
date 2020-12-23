package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

func TestMergeSort(t *testing.T) {
	numElements := 16 << 20
	src := make([]int64, numElements)
	original := make([]int64, numElements)
	prepare2(original)
	copy(src, original)

	msg := fmt.Sprintf("%v\n",src)
	MergeSortConcurrent2(original)
	sort.Slice(src, func(i, j int) bool { return src[i] < src[j] })
	assert.Equal(t, original, src, msg)

	//src := []int64{20, 16, 19, 15, 14, 20, 17, 10, 8, 20, 1, 3, 1, 2, 16, 13, 21, 9, 5, 2, 17, 11}
	////original := []int64{20, 16, 19, 15, 14, 20, 17, 10, 8, 20, 1, 3, 1, 2, 16, 13, 21, 9, 5, 2, 17, 11}
	//interSrc = make([]int64, len(src))
	//original := make([]int64, len(src))
	//copy(src, original)
	//fmt.Println(len(src))
	////assert.Equal(t, src, original)
	////MergeSortBasic2(original)
	//MergeSortConcurrent2(original)
	//sort.Slice(src, func(i, j int) bool { return src[i] < src[j] })
	//assert.Equal(t, original, src)
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
