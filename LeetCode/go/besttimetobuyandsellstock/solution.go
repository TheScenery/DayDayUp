package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	n := len(prices)
	mp := 0
	if n > 1 {
		mc := prices[0]
		for i := 1; i < n; i++ {
			p := prices[i]
			profit := p - mc
			if profit > mp {
				mp = profit
			}
			if p < mc {
				mc = p
			}
		}
	}
	return mp
}
