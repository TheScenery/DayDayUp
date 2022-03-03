package gupiaodezuidalirunlcof

func maxProfit(prices []int) int {
	max := 0
	n := len(prices)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			diff := prices[j] - prices[i]
			if diff > max {
				max = diff
			}
		}
	}
	return max
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func maxProfit1(prices []int) int {
	n := len(prices)
	dp := make([]int, n)
	dp[0] = 0
	minCost := prices[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1], prices[i]-minCost)
		if prices[i] < minCost {
			minCost = prices[i]
		}
	}
	return dp[n-1]
}

func maxProfit2(prices []int) int {
	n := len(prices)
	max := 0
	minCost := prices[0]
	for i := 1; i < n; i++ {
		profit := prices[i] - minCost
		if profit > max {
			max = profit
		}
		if prices[i] < minCost {
			minCost = prices[i]
		}
	}
	return max
}
