package climbingstairs

func climbStairs(n int) int {
	prev1 := 1
	prev2 := 2
	if n == 1 {
		return prev1
	}
	if n == 2 {
		return prev2
	}
	result := 0
	for i := 3; i <= n; i++ {
		result = prev1 + prev2
		prev1, prev2 = prev2, result
	}
	return result
}
