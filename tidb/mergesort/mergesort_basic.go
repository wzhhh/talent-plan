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