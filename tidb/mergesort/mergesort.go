package main

import (
	"fmt"
	"sync"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	if len(src) <= 1 {
		return
	}
	mid := len(src)/2
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		MergeSort(src[:mid])
	}()
	go func() {
		defer wg.Done()
		MergeSort(src[mid:])
	}()
	wg.Wait()
	merge(src, 0, mid, len(src)-1)
}

func MergeSort0(src []int64) {
	if len(src) <= 1 {
		return
	}
	mid := len(src)/2
	MergeSort(src[:mid])
	MergeSort(src[mid:])
	merge(src, 0, mid, len(src)-1)
}

func merge(a []int64, start, mid, end int)  {
	i := start
	j := mid
	k := end
	for i < j && j <= k {
		t := 0
		for i < j && a[i] <= a[j]  {
			i++
		}
		for j <= k && a[j] < a[i]  {
			j++
			t++
		}
		shift(a, i, j-t, j-1)
		i += t
	}
}

//func mergeChan(a, b <-chan int64) chan int64 {
//
//}

func shift(a []int64, start, mid, end int) {
	reverse(a, start, mid-1)
	reverse(a, mid, end)
	reverse(a, start, end)
}

func reverse(a []int64, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	a := []int64{3,6,1,4,734,5567,432,2,57,43,5}
	MergeSort(a)
	fmt.Println(a)
}