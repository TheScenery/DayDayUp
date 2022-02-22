package triangle

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		n := i + 1
		dp[i] = make([]int, n)
		if i == 0 {
			dp[0][0] = triangle[0][0]
			continue
		}
		for j := 0; j < n; j++ {
			currentValue := triangle[i][j]
			if j == 0 {
				dp[i][j] = dp[i-1][j] + currentValue
			} else if j == n-1 {
				dp[i][j] = dp[i-1][j-1] + currentValue
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + currentValue
			}
		}
	}
	leaf := dp[m-1]
	minValue := leaf[0]
	for i := 1; i < len(leaf); i++ {
		if leaf[i] < minValue {
			minValue = leaf[i]
		}
	}
	return minValue
}
