package bashuzifanyichengzifuchuanlcof

import "strconv"

func translateNum(num int) int {
	s := strconv.Itoa(num)
	n := len(s)
	if n < 2 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	if s[0:2] < "26" {
		dp[1] = 2
	} else {
		dp[1] = 1
	}
	for i := 2; i < n; i++ {
		if s[i-1] != '0' && s[i-1:i+1] < "26" {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n-1]
}
