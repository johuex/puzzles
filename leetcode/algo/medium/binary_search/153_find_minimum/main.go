package findminimum

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		m := l + (r-l)/2
		// search minimum by comparing m with r
		// because we have to ASc subarray in array
		// and we need to find minimum of left subarray
		if nums[m] < nums[r] {
			r = m
		} else {
			l = m + 1
		}
	}

	return nums[l]
}
