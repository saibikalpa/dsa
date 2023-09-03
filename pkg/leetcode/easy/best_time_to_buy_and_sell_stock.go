package easy

func maxProfit(prices []int) int {
	var (
		aux       = make([]int, len(prices))
		maxProfit int
		max       int
	)
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > max {
			max = prices[i]
		}
		aux[i] = max
	}

	for i := 0; i < len(prices); i++ {
		if aux[i]-prices[i] > maxProfit {
			maxProfit = aux[i] - prices[i]
		}
	}
	return maxProfit
}

func maxProfitOne(prices []int) int {
	maxProfit, minCP := 0, prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minCP {
			minCP = prices[i]
		} else if profit := prices[i] - minCP; profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
