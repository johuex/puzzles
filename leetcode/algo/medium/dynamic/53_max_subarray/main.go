package maxsubarray

func maxSubArray(nums []int) int {
	curr_max := nums[0]
	glob_max := nums[0]

	for i := 1; i < len(nums); i++ {
		curr_max = max(nums[i], nums[i]+curr_max)
		glob_max = max(glob_max, curr_max)
	}

	return glob_max
}
