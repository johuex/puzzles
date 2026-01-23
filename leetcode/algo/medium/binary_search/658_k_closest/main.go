package kclosest

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)

	// lower_bound: arr[i] >= x
	l, r := 0, n
	for l < r {
		m := l + (r-l)/2
		if arr[m] < x {
			l = m + 1
		} else {
			r = m
		}
	}

	// pointers for building result array
	left := l - 1
	right := l
	for right-left-1 < k {
		if left < 0 {
			// when target less than first elem
			right++
		} else if right >= n {
			// when target is greater than last elem
			left--
		} else if Abs(arr[left]-x) <= Abs(arr[right]-x) {
			left--
		} else {
			right++
		}
	}
	return arr[left+1 : right]
}
