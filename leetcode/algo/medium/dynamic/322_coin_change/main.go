package coinchange

func coinChange(coins []int, amount int) int {
	// init dp array
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1 // mark as infnity
	}
	dp[0] = 0

	for i := 1; i < amount+1; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				// choose
				// actual change
				// or i - coin combination
				dp[i] = min(dp[i], 1+dp[i-coin])
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1 // no way to coin change
	}
	return dp[amount]
}
