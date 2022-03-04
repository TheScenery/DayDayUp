package lianxuzishuzudezuidahelcof

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	maxValue := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > maxValue {
			maxValue = dp[i]
		}
	}
	return maxValue
}
