package main

import (
	"fmt"
	"sync"
)

var interSrc []int64

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
		MergeSort(src[0:mid])
	}()
	go func() {
		defer wg.Done()
		MergeSort(src[mid:len(src)])
	}()
	wg.Wait()
	r := MergeByChan(src[0:mid], src[mid:len(src)])
	src = r
	fmt.Println(src)
	//if len(src) <= 1 {
	//	return
	//}
	//nums := 8
	//size := 16
	//srcChan := produce(src, size)
	//outChanSlice := fanOut(srcChan, nums, size)
	//resultChan := mergeChanSlice(outChanSlice, nums, size)
	//
	//res := make([]int64, 0, len(src))
	//for i := range resultChan[0] {
	//	res = append(res, i)
	//}
	//
	//fmt.Println(res)
}

func produce(src []int64, size int) <-chan int64 {
	outChan := make(chan int64, size)
	go func() {
		defer close(outChan)
		for _, i := range src {
			outChan <- i
		}
	}()
	return outChan
}

func fanOut(inChan <-chan int64, nums, size int) []chan int64 {
	outChanSlice := make([]chan int64, nums)
	var wg sync.WaitGroup
	wg.Add(nums)
	for i := 0; i < nums; i++ {
		go func(i int) {
			outChanSlice[i] = make(chan int64, size)
			defer func() {
				wg.Done()
				close(outChanSlice[i])
			}()
			for n := range inChan {
				outChanSlice[i] <- n
			}
		}(i)
	}
	go func() {
		wg.Wait()
	}()
	return outChanSlice
}

func mergeChanSlice(chanSlice []chan int64, nums, size int) []chan int64 {
	if nums == 1 {
		return chanSlice
	}
	outChanSlice := make([]chan int64, nums/2)
	var wg sync.WaitGroup
	wg.Add(nums / 2)
	for i := 0; i < nums; i += 2 {
		go func(i int) {
			defer func() {
				wg.Done()
			}()
			outChanSlice[i/2] = mergeChan(chanSlice[i], chanSlice[i+1])
		}(i)
	}
	go func() {
		wg.Wait()
	}()
	return mergeChanSlice(outChanSlice, nums/2, size)
}

func mergeChan(ch1, ch2 <-chan int64) chan int64 {

	out := make(chan int64, 1)
	go func() { // 等上游的数据 （这里有阻塞，和常规的阻塞队列并无不同）
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2 // 取数据
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) { // 取到最小值, 就推到 out 中
				out <- v1
				v1, ok1 = <-ch1
			} else {
				out <- v2
				v2, ok2 = <-ch2
			}
		} // 显式关闭
		close(out)
	}() // 开完goroutine后, 主线程继续执行, 不会阻塞
	return out
}

func MergeByChan2(src []int64, start, mid, end int) {
	rChan := mergeChan(produce(src[start:mid], 2), produce(src[mid:end+1], 2))
	res := make([]int64, end-start+1)
	i := 0
	for n := range rChan {
		res[i] = n
		i++
	}
	for i := range res {
		src[i+start] = res[i]
	}
}

func MergeByChan(a, b []int64) []int64 {
	aChan := produce(a, 1)
	bChan := produce(b, 1)
	rChan := mergeChan(aChan, bChan)

	res := make([]int64, 0, len(a)+len(b))
	for i := range rChan {
		res = append(res, i)
	}

	return res
}

func main() {
	c := []int64{3, 6, 1, 4, 734, 5567, 432, 2, 57, 43, 5}
	MergeSortConcurrent2(c)
	fmt.Println(c)
}