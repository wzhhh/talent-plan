package main

import (
	"runtime"
	"sync"
)

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
	merge(src, 0, mid, len(src)-1)
}

func MergeSortConcurrent2(src []int64) {
	if len(src) <= 1 {
		return
	}
	nums := runtime.NumCPU()
	//fmt.Println(nums)
	var wg sync.WaitGroup
	wg.Add(nums)
	size := len(src)/nums
	for i := 0; i < nums; i++ {
		start := i * size
		end := start + size
		if i == nums-1 {
			end = len(src)
		}
		go func(start, end int) {
			defer wg.Done()
			MergeSortBasic(src[start:end])
		}(start, end)
	}
	wg.Wait()
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
			end := start + 2*i*size -1
			if j + 2*i >= nums {
				end = len(src)-1
			}
			wg.Add(1)
			go func(start, mid, end int) {
				defer wg.Done()
				merge(src, start, mid, end)
			}(start, mid, end)
		}
		wg.Wait()
	}
	return
}