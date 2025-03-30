package beststock

import "sort"

// [7,2,5,1,3,6] -> multiple, minimum righter
// [7,2,8,1,6,4] -> multiple, minimum lefter
// [7,6,4,3,1] -> impossible

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	var min int = prices[0]
	var max int = prices[0]
	maxProfit := make([]int, 0)
	for _, price := range prices {
		if price < min {
			min = price
			max = price
		}
		if price > max {
			max = price
			maxProfit = append(maxProfit, max-min)
		}
	}
	if len(maxProfit) == 0 {
		return 0
	}
	sort.Slice(maxProfit, func(i, j int) bool {
		// in descending ordre
		return maxProfit[i] > maxProfit[j]
	})
	return maxProfit[0]
}

func maxProfitOpt(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	var min int = prices[0]
	var profit int
	var currentProfit int

	for _, price := range prices {
		if price < min {
			min = price
			continue
		}
		currentProfit = price - min
		if currentProfit > profit {
			profit = currentProfit
		}
	}
	return profit
}
