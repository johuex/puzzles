package maxincsubseq

func findNumberOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	// init DPs
	dp := make([]int, len(nums))    // max length of LIS
	dpCnt := make([]int, len(nums)) // how many LIS at idx
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
		dpCnt[i] = 1
	}
	maxLISLength := 1
	var maxCnt int

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					dpCnt[i] = dpCnt[j]
				} else if dp[i] == dp[j]+1 {
					dpCnt[i] += dpCnt[j]
				}
			}
			maxLISLength = max(maxLISLength, dp[i])
		}
	}

	for i := 0; i < len(dp); i++ {
		if dp[i] == maxLISLength {
			maxCnt += dpCnt[i]
		}
	}

	return maxCnt
}
