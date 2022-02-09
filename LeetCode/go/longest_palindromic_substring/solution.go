package longestpalindromicsubstring

func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	start := 0
	maxLength := 1
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
		dp[i][i] = true
	}

	for i := length - 1; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			if s[i] == s[j] {
				if j-i < 2 || dp[i+1][j-1] {
					dp[i][j] = true
					subLen := j - i + 1
					if subLen > maxLength {
						maxLength = subLen
						start = i
					}
				}
			}
		}
	}

	return s[start : start+maxLength]
}
