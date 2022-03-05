package zuichangbuhanzhongfuzifudezizifuchuanlcof

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	start, end := 0, 0
	visitedMap := make(map[byte]bool)
	for {
		if end >= len(s) {
			if end-start > maxLength {
				maxLength = end - start
			}
			break
		}
		if visitedMap[s[end]] {
			if end-start > maxLength {
				maxLength = end - start
			}
			for start < end && s[start] != s[end] {
				visitedMap[s[start]] = false
				start++
			}
			start++
		}
		visitedMap[s[end]] = true
		end++
	}
	return maxLength
}
