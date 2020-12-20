package main

func MergeSortBasic(src []int64) {
	if len(src) <= 1 {
		return
	}
	mid := len(src)/2
	MergeSortBasic(src[:mid])
	MergeSortBasic(src[mid:])
	merge(src, 0, mid, len(src)-1)
}

func MergeSortBasic2(src []int64) {
	if len(src) <= 1 {
		return
	}
	mid := len(src)/2
	MergeSortBasic2(src[:mid])
	MergeSortBasic2(src[mid:])
	merge2(src, 0, mid, len(src)-1)
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

func merge2(src []int64, start, mid, end int) {
	left := start
	right := mid
	idx := start
	for left < mid && right < end {
		if src[left] > src[right] {
			interSrc[idx] = src[right]
			right++
		} else {
			interSrc[idx] = src[left]
			left++
		}
		idx++
	}

	for left < mid {
		interSrc[idx] = src[left]
		left++
		idx++
	}

	for right < end {
		interSrc[idx] = src[right]
		right++
		idx++
	}

	for i := start; i < end; i++ {
		src[i] = interSrc[i]
	}
}

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