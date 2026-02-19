package sellbuystockwfee

func maxProfit(prices []int, fee int) int {
	cash := 0          // no profit yet
	hold := -prices[0] // bought first

	for i := 1; i < len(prices); i++ {
		cash = max(cash, hold+prices[i]-fee) // decide sell or not
		hold = max(hold, cash-prices[i])     // decide buy or nort
	}

	return cash // return our total profit
}
