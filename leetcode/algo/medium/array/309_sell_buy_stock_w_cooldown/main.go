package sellbuystockwcooldown

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	hold := -prices[0] // bought at first day
	sold := 0          // not sold
	rest := 0          // do nothing

	for i := 1; i < len(prices); i++ {
		prevSold := sold
		sold = hold + prices[i]          // sell today
		hold = max(hold, rest-prices[i]) // hold ot buy (only from rest, not from sold) rest always on 1 step behind sold
		rest = max(rest, prevSold)       // cooldown or nothing
	}

	return max(sold, rest) // choose the best
}
