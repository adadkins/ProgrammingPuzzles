package programmingpuzzles

func StockMaxProfit(stockPrices []int) (maxProfit int) {
	for i := 0; i < len(stockPrices)-1; i++ {
		for j := i + 1; j < len(stockPrices); j++ {
			if stockPrices[j]-stockPrices[i] > maxProfit {
				maxProfit = stockPrices[j] - stockPrices[i]
			}
		}
	}
	return maxProfit
}
