package maximumsubarray

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	maxValue := nums[0]
	for i := range nums {
		if i == 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = max(dp[i-1]+nums[i], nums[i])
		}
		if dp[i] > maxValue {
			maxValue = dp[i]
		}
	}
	return maxValue
}
