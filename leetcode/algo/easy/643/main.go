package findmaxaverage

import "math"

func findMaxAverage(nums []int, k int) float64 {
	var queueSum int

	// init sliding window
	for i := range k {
		queueSum += nums[i]
	}
	maxAvg := float64(queueSum) / float64(k)
	for i := k; i < len(nums); i++ {
		queueSum -= nums[i-k]
		queueSum += nums[i]
		avg := float64(queueSum) / float64(k)
		maxAvg = math.Max(maxAvg, avg)
	}

	return maxAvg
}
