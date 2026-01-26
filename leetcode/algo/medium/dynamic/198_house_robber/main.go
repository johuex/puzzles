package houserobber

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	nums = append([]int{0}, nums...)
	dp := make([]int, len(nums))
	dp[1] = nums[1] // init DP with first house

	for i := 2; i < len(nums); i++ {
		// choose prev value or accumulated prev-1 house ago + current money
		// because we can't rob two near house
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[len(dp)-1] // answer in last DP elem
}
