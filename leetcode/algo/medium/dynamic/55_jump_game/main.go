package jumpgame

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	// set possible jumps from left to right
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i-1]-1, dp[i-1]-1, nums[i])
	}

	var nullExist bool
	// find jump breks: point where we can't jump
	for i := 0; i < len(nums)-2; i++ {
		if dp[i] == 0 {
			nullExist = true
			break
		}
	}

	if !nullExist && dp[len(nums)-2] > 0 {
		// one jump exist to last elem
		// and no jump breaks before
		return true
	}
	return false
}
