package maxinwindow

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 || k > n {
		return []int{}
	}

	deque := make([]int, 0, k) // храним индексы
	res := make([]int, n-k+1)
	resIdx := 0

	for i := 0; i < n; i++ {
		// force pop left, out of window
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		// before adding -- pop right idx that greater actual by value (keep monotonously decreasing function)
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}

		// push right
		deque = append(deque, i)

		// store min as deq[0], fill only if window is full
		if i >= k-1 {
			res[resIdx] = nums[deque[0]]
			resIdx++
		}
	}

	return res
}
