package mulpilexceptself

func productExceptSelf(nums []int) []int {
	pref := make([]int, len(nums))
	suff := make([]int, len(nums))

	// init data
	pref[0] = 1
	suff[len(nums)-1] = 1

	// fill prefix
	for i := 1; i < len(nums); i++ {
		pref[i] = pref[i-1] * nums[i-1]
	}

	// fill suffix
	for i := len(nums) - 2; i > -1; i-- {
		suff[i] = suff[i+1] * nums[i+1]

	}

	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = pref[i] * suff[i]
	}

	return res
}

func productExceptSelfRamOptimized(nums []int) []int {
	n := len(nums)
	out := make([]int, n) // out -- both prefix and suffix array

	prefix := 1
	for i := range nums {
		out[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := range nums {
		j := n - i - 1
		out[j] *= postfix
		postfix *= nums[j]
	}

	return out
}
