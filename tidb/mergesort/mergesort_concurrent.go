package main

import (
	"runtime"
	"sync"
)

type partSrc struct {
	start int
	end   int
}

func MergeSortConcurrent1(src []int64) {
	if len(src) <= 1 {
		return
	}
	mid := len(src)/2
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		MergeSortConcurrent1(src[:mid])
	}()
	go func() {
		defer wg.Done()
		MergeSortConcurrent1(src[mid:])
	}()
	wg.Wait()
	merge2(src, 0, mid, len(src))
}

func MergeSortConcurrent2(src []int64) {
	if len(src) <= 1 {
		return
	}
	nums := runtime.NumCPU()
	parts := make([]partSrc, nums)
	interSrc = make([]int64, len(src))
	var wg sync.WaitGroup
	wg.Add(nums)
	size := len(src)/nums
	for i := 0; i < nums; i++ {
		start := i * size
		end := start + size
		if i == nums-1 {
			end = len(src)
		}
		parts[i] = partSrc{start, end}
		go func(start, end int) {
			defer wg.Done()
			coreSort(src, start, end)
		}(start, end)
	}
	wg.Wait()
	//b2UpMerge(src, parts)
	concurrencyMerge(src, nums)
	return
}


func concurrencyMerge(src []int64, nums int) {
	size := len(src)/nums
	for i := 1; i < nums; i *=2 {
		var wg sync.WaitGroup
		for j := 0; j < nums; j += i*2 {
			start := j * size
			mid := start + i * size
			end := start + 2*i*size
			if j+2*i == nums {
				end = len(src)
			}
			wg.Add(1)
			go func(start, mid, end int) {
				defer wg.Done()
				merge2(src, start, mid, end)
			}(start, mid, end)
		}
		wg.Wait()
	}
	return
}

func b2UpMerge(src []int64, parts []partSrc) {
	n := len(parts)
	for size := 1; size < n; size *= 2 {
		var wg sync.WaitGroup

		for low := 0; low < n-size; low += size * 2 {
			start := parts[low].start
			mid := parts[low+size-1].end
			endIdx := low + size*2 - 1
			if endIdx > n-1 {
				endIdx = n - 1
			}
			end := parts[endIdx].end
			wg.Add(1)
			go func(start, mid, end int) {
				defer wg.Done()
				merge2(src, start, mid, end)
			}(start, mid, end)
		}
		wg.Wait()
	}
}