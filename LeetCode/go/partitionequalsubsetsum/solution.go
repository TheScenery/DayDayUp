package partitionequalsubsetsum

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	sum := 0
	max := 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	if target < max {
		return false
	}
	dp := make([][]bool, n)
	for i, num := range nums {
		dp[i] = make([]bool, target+1)
		for j := 0; j <= target; j++ {
			if i == 0 {
				dp[i][j] = num == j
			} else {
				if num == j {
					dp[i][j] = true
				} else if j > num {
					dp[i][j] = dp[i-1][j-num] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}

	return dp[n-1][target]
}

func canPartition1(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	sum := 0
	max := 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	if target < max {
		return false
	}
	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}
