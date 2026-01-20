package binaryindex

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := l + (r - l/2) // because we can't devide by 0
		tmpN := nums[m]
		if tmpN < target {
			l = m + 1
		} else if tmpN > target {
			r = m - 1
		} else {
			return m
		}
	}
	// case when left and right borders are the same or crossed
	if l >= len(nums) || l <= 0 || nums[l] != target {
		return -1
	}
	return l
}
