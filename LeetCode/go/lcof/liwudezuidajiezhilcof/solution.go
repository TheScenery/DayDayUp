package liwudezuidajiezhilcof

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func getValue(dp [][]int, i, j int) int {
	m := len(dp)
	n := len(dp[0])
	if i >= 0 && i < m && j >= 0 && j < n {
		return dp[i][j]
	}
	return 0
}

func maxValue(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else {
				dp[i][j] = grid[i][j] + max(getValue(dp, i-1, j), getValue(dp, i, j-1))
			}
		}
	}
	return dp[m-1][n-1]
}
