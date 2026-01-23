package searchinrotated

func search(nums []int, target int) int {
	// first: find the smallest number, it's our pivot
	if len(nums) == 0 {
		return -1
	}
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
	pivot := l // it's our minimum in array and splitter for array
	if target == nums[pivot] {
		// our target in pivot :)
		return pivot
	}

	// second step: decide in which half to search target
	// usual binary search in one of the half
	if target > nums[len(nums)-1] {
		l = 0
		r = pivot - 1
	} else {
		l = pivot
		r = len(nums) - 1
	}

	for l < r {
		m := l + (r-l)/2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return m
		}
	}

	if nums[l] != target {
		return -1
	}
	return l
}
