package longestincseqlength

func findLengthOfLCIS(nums []int) int {
	// init DPs
	dp := make([]int, len(nums)) // max length of LIS
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	maxLISLength := 1

	for i := 1; i < len(nums); i++ {
		j := i - 1
		if nums[j] < nums[i] {
			if dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		maxLISLength = max(maxLISLength, dp[i])
	}

	return maxLISLength
}
