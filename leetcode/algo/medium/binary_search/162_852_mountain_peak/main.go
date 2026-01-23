package mountainpeak

func peakIndexInMountainArray(arr []int) int {
	// in task said that we have only one global maximum and sides increase and decrease linear

	l := 0
	r := len(arr) - 1

	for l < r {
		m := l + (r-l)/2
		if arr[m] < arr[m+1] {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}
